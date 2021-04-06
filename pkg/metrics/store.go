package metrics

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type tableInfo struct {
	Name   string
	RollUp time.Duration
	Next   string
}

type tableKey struct {
	IngressController string
	Service           string
}

// WaterMarks contain low water marks for a table.
type WaterMarks map[tableKey]int

// Store is a metrics store.
type Store struct {
	tables []tableInfo

	mu    sync.RWMutex
	data  map[string]map[tableKey]DataPoints
	marks map[string]WaterMarks

	// NowFunc is the function used to test time.
	nowFunc func() time.Time
}

// NewStore returns metrics store.
func NewStore() *Store {
	tables := []tableInfo{
		{Name: "1m", RollUp: 10 * time.Minute, Next: "10m"},
		{Name: "10m", RollUp: time.Hour, Next: "1h"},
		{Name: "1h", RollUp: 24 * time.Hour, Next: "1d"},
		{Name: "1d", RollUp: 30 * 24 * time.Hour},
	}

	tbls := make(map[string]map[tableKey]DataPoints, len(tables))
	marks := make(map[string]WaterMarks, len(tables))
	for _, info := range tables {
		tbls[info.Name] = map[tableKey]DataPoints{}
		marks[info.Name] = map[tableKey]int{}
	}

	return &Store{
		tables:  tables,
		data:    tbls,
		marks:   marks,
		nowFunc: time.Now,
	}
}

// Populate populates the store with initial data points.
func (s *Store) Populate(tbl string, grps []DataPointGroup) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	table, ok := s.data[tbl]
	if !ok {
		return fmt.Errorf("table %q does not exist", tbl)
	}

	for _, v := range grps {
		key := tableKey{IngressController: v.IngressController, Service: v.Service}
		if len(v.DataPoints) == 0 {
			continue
		}

		dataPoints := v.DataPoints
		sort.Slice(dataPoints, func(i, j int) bool {
			return dataPoints[i].Timestamp < dataPoints[j].Timestamp
		})
		table[key] = dataPoints
		s.marks[tbl][key] = len(dataPoints)
	}

	return nil
}

// Insert inserts a value for an ingress and service.
func (s *Store) Insert(ic string, svcs map[string]DataPoint) {
	s.mu.Lock()
	defer s.mu.Unlock()

	table := s.data["1m"]

	for svc, pnt := range svcs {
		key := tableKey{IngressController: ic, Service: svc}
		pnts := table[key]
		pnts = append(pnts, pnt)
		table[key] = pnts
	}
}

// ForEachFunc represents a function that will be called while iterating over a table.
// Each time this function is called, a unique ingress controller and service will
// be given with their set of points.
type ForEachFunc func(ic, svc string, pnts DataPoints)

// ForEachUnmarked iterates over a table, executing fn for each row that
// has not been marked.
func (s *Store) ForEachUnmarked(tbl string, fn ForEachFunc) WaterMarks {
	s.mu.RLock()
	defer s.mu.RUnlock()

	table, ok := s.data[tbl]
	if !ok {
		return nil
	}

	newMarks := make(WaterMarks)
	for k, v := range table {
		newMarks[k] = len(v)

		mark := s.marks[tbl][k]
		if mark == len(v) {
			continue
		}

		fn(k.IngressController, k.Service, v[mark:])
	}

	return newMarks
}

// CommitMarks sets the new low water marks for a table.
func (s *Store) CommitMarks(tbl string, marks WaterMarks) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.marks[tbl]
	if !ok {
		return
	}

	s.marks[tbl] = marks
}

// RollUp creates combines data points.
//
// Rollup goes through each table and aggregates the points into complete
// points of the next granularity, if that point does not already exist in the
// next table.
func (s *Store) RollUp() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, tblInfo := range s.tables {
		tbl, gran, dest := tblInfo.Name, tblInfo.RollUp, tblInfo.Next
		if dest == "" {
			continue
		}

		rollUpEnd := s.nowFunc().UTC().Truncate(gran).Unix()

		res := map[tableKey]map[int64]DataPoints{}
		for key, data := range s.data[tbl] {
			destPnts := s.data[dest][key]

			for _, pnt := range data {
				// As data points are in asc order, when the current roll up period is reached, we can stop.
				if pnt.Timestamp >= rollUpEnd {
					continue
				}

				destTS := pnt.Timestamp - (pnt.Timestamp % int64(gran/time.Second))
				// Check if the timestamp exists in the destination table.
				if i, _ := destPnts.Get(destTS); i >= 0 {
					continue
				}

				tsPnts, ok := res[key]
				if !ok {
					tsPnts = map[int64]DataPoints{}
				}
				pnts := tsPnts[destTS]
				pnts = append(pnts, pnt)
				tsPnts[destTS] = pnts
				res[key] = tsPnts
			}
		}

		// Insert new computed points into dest.
		table := s.data[dest]
		for key, tsPnts := range res {
			for ts, pnts := range tsPnts {
				pnt := pnts.Aggregate()
				pnt.Timestamp = ts

				destPnts := table[key]
				destPnts = append(destPnts, pnt)
				table[key] = destPnts
			}
		}
	}
}

// Cleanup removes old data points no longer needed for roll up.
func (s *Store) Cleanup() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, tblInfo := range s.tables {
		tbl, gran := tblInfo.Name, tblInfo.RollUp

		keep := s.nowFunc().UTC().Truncate(gran).Unix()
		for k, data := range s.data[tbl] {
			pnts := data
			idx := sort.Search(len(pnts), func(i int) bool {
				return pnts[i].Timestamp >= keep
			})

			mark := s.marks[tbl][k]
			if idx > mark {
				idx = mark
			}

			copy(pnts[0:], pnts[idx:])
			pnts = pnts[0 : len(pnts)-idx]
			s.data[tbl][k] = pnts
			s.marks[tbl][k] = mark - idx
		}
	}
}
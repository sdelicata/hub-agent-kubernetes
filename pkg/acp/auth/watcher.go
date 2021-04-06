package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/traefik/neo-agent/pkg/acp"
	"github.com/traefik/neo-agent/pkg/acp/basicauth"
	"github.com/traefik/neo-agent/pkg/acp/digestauth"
	"github.com/traefik/neo-agent/pkg/acp/jwt"
	neov1alpha1 "github.com/traefik/neo-agent/pkg/crd/api/neo/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: if we use the same watcher for all resources, then we need to restart it when new CRDs are
// created/removed like for example when Traefik is installed and IngressRoutes are added.
// Always listening to non-existing resources would cause errors.
// Also, if multiple clients of this watcher are not interested in the same resources
// add a parameter to NewWatcher to subscribe only to a subset of events.

// Watcher watches access control policy resources and builds configurations out of them.
type Watcher struct {
	configsMu sync.RWMutex
	configs   map[string]*acp.Config
	previous  map[string]*acp.Config

	refresh chan struct{}

	switcher *HTTPHandlerSwitcher
}

// NewWatcher returns a new watcher to track ACP resources. It calls the given Updater when an ACP is modified at most
// once every throttle.
func NewWatcher(switcher *HTTPHandlerSwitcher) *Watcher {
	return &Watcher{
		configs:  make(map[string]*acp.Config),
		refresh:  make(chan struct{}, 1),
		switcher: switcher,
	}
}

// Run launches listener if the watcher is dirty.
func (w *Watcher) Run(ctx context.Context) {
	for {
		select {
		case <-w.refresh:
			w.configsMu.RLock()

			if reflect.DeepEqual(w.previous, w.configs) {
				w.configsMu.RUnlock()
				continue
			}

			cfgs := make(map[string]*acp.Config, len(w.configs))
			for k, v := range w.configs {
				cfgs[k] = v
			}

			w.previous = cfgs

			w.configsMu.RUnlock()

			log.Debug().Msg("Refreshing ACP handlers")

			routes, err := buildRoutes(cfgs)
			if err != nil {
				log.Error().Err(err).Msg("Unable to switch ACP handlers")
				continue
			}

			w.switcher.UpdateHandler(routes)

		case <-ctx.Done():
			return
		}
	}
}

// OnAdd implements Kubernetes cache.ResourceEventHandler so it can be used as an informer event handler.
func (w *Watcher) OnAdd(obj interface{}) {
	v, ok := obj.(*neov1alpha1.AccessControlPolicy)
	if !ok {
		log.Error().
			Str("component", "acpWatcher").
			Str("type", fmt.Sprintf("%T", obj)).
			Msg("Received add event of unknown type")
		return
	}

	w.configsMu.Lock()
	w.configs[canonicalName(v.ObjectMeta.Name, v.ObjectMeta.Namespace)] = acp.ConfigFromPolicy(v)
	w.configsMu.Unlock()

	select {
	case w.refresh <- struct{}{}:
	default:
	}
}

// OnUpdate implements Kubernetes cache.ResourceEventHandler so it can be used as an informer event handler.
func (w *Watcher) OnUpdate(_, newObj interface{}) {
	v, ok := newObj.(*neov1alpha1.AccessControlPolicy)
	if !ok {
		log.Error().
			Str("component", "acpWatcher").
			Str("type", fmt.Sprintf("%T", newObj)).
			Msg("Received update event of unknown type")
		return
	}

	polName := canonicalName(v.ObjectMeta.Name, v.ObjectMeta.Namespace)
	cfg := acp.ConfigFromPolicy(v)

	w.configsMu.Lock()
	w.configs[polName] = cfg
	w.configsMu.Unlock()

	select {
	case w.refresh <- struct{}{}:
	default:
	}
}

// OnDelete implements Kubernetes cache.ResourceEventHandler so it can be used as an informer event handler.
func (w *Watcher) OnDelete(obj interface{}) {
	v, ok := obj.(*neov1alpha1.AccessControlPolicy)
	if !ok {
		log.Error().
			Str("component", "acpWatcher").
			Str("type", fmt.Sprintf("%T", obj)).
			Msg("Received delete event of unknown type")
		return
	}

	w.configsMu.Lock()
	delete(w.configs, canonicalName(v.ObjectMeta.Name, v.ObjectMeta.Namespace))
	w.configsMu.Unlock()

	select {
	case w.refresh <- struct{}{}:
	default:
	}
}

func canonicalName(name, ns string) string {
	if ns == "" {
		ns = metav1.NamespaceDefault
	}

	return name + "@" + ns
}

func buildRoutes(cfgs map[string]*acp.Config) (http.Handler, error) {
	mux := http.NewServeMux()

	for name, cfg := range cfgs {
		switch {
		case cfg.JWT != nil:
			jwtHandler, err := jwt.NewHandler(cfg.JWT, name)
			if err != nil {
				return nil, fmt.Errorf("create %q JWT ACP handler: %w", name, err)
			}

			path := "/" + name

			log.Debug().Str("acp_name", name).Str("path", path).Msg("Registering JWT ACP handler")

			mux.Handle(path, jwtHandler)

		case cfg.BasicAuth != nil:
			h, err := basicauth.NewHandler(cfg.BasicAuth, name)
			if err != nil {
				return nil, fmt.Errorf("create %q basic auth ACP handler: %w", name, err)
			}
			path := "/" + name
			log.Debug().Str("acp_name", name).Str("path", path).Msg("Registering basic auth ACP handler")
			mux.Handle(path, h)

		case cfg.DigestAuth != nil:
			h, err := digestauth.NewHandler(cfg.DigestAuth, name)
			if err != nil {
				return nil, fmt.Errorf("create %q digest auth ACP handler: %w", name, err)
			}
			path := "/" + name
			log.Debug().Str("acp_name", name).Str("path", path).Msg("Registering digest auth ACP handler")
			mux.Handle(path, h)

		default:
			return nil, errors.New("unknown ACP handler type")
		}
	}

	return mux, nil
}
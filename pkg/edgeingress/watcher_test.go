/*
Copyright (C) 2022-2023 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package edgeingress

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	hubv1alpha1 "github.com/traefik/hub-agent-kubernetes/pkg/crd/api/hub/v1alpha1"
	hubkubemock "github.com/traefik/hub-agent-kubernetes/pkg/crd/generated/client/hub/clientset/versioned/fake"
	hubinformer "github.com/traefik/hub-agent-kubernetes/pkg/crd/generated/client/hub/informers/externalversions"
	traefikkubemock "github.com/traefik/hub-agent-kubernetes/pkg/crd/generated/client/traefik/clientset/versioned/fake"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	kubemock "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"
	"k8s.io/utils/pointer"
)

var toUpdate = hubv1alpha1.EdgeIngress{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "toUpdate",
		Namespace: "default",
		UID:       "uid",
		Labels: map[string]string{
			"app.kubernetes.io/managed-by": "traefik-hub",
		},
	},
	Spec: hubv1alpha1.EdgeIngressSpec{
		Service: hubv1alpha1.EdgeIngressService{
			Name: "service-1",
			Port: 8081,
		},
	},
	Status: hubv1alpha1.EdgeIngressStatus{
		Version:    "version-1",
		SyncedAt:   metav1.NewTime(time.Now().Add(-time.Hour)),
		URLs:       "https://sad-bat-456.hub-traefik.io",
		SpecHash:   "yxjMx+3w4R4B4YPzoGkqi/g9rLs=",
		Connection: hubv1alpha1.EdgeIngressConnectionDown,
	},
}

var toDelete = hubv1alpha1.EdgeIngress{
	ObjectMeta: metav1.ObjectMeta{
		Name:      "toDelete",
		Namespace: "default",
		Labels: map[string]string{
			"app.kubernetes.io/managed-by": "traefik-hub",
		},
	},
	Spec: hubv1alpha1.EdgeIngressSpec{
		Service: hubv1alpha1.EdgeIngressService{
			Name: "service-3",
			Port: 8080,
		},
	},
	Status: hubv1alpha1.EdgeIngressStatus{
		Version:    "version-1",
		SyncedAt:   metav1.NewTime(time.Now().Add(-time.Hour)),
		URLs:       "https://broken-cat-123.hub-traefik.io",
		SpecHash:   "7kfh53DLsXumNEaO/nkeVYs/5CI=",
		Connection: hubv1alpha1.EdgeIngressConnectionDown,
	},
}

func Test_WatcherRun(t *testing.T) {
	clientSetHub := hubkubemock.NewSimpleClientset([]runtime.Object{&toUpdate, &toDelete}...)
	clientSet := kubemock.NewSimpleClientset()

	ctx, cancel := context.WithCancel(context.Background())
	hubInformer := hubinformer.NewSharedInformerFactory(clientSetHub, 0)

	edgeIngressInformer := hubInformer.Hub().V1alpha1().EdgeIngresses().Informer()

	hubInformer.Start(ctx.Done())
	cache.WaitForCacheSync(ctx.Done(), edgeIngressInformer.HasSynced)

	edgeIngresses := []EdgeIngress{
		{
			Name:      "toCreate",
			Namespace: "default",
			Domain:    "majestic-beaver-123.hub-traefik.io",
			Version:   "version-1",
			Service:   Service{Name: "service-1", Port: 8080},
			ACP:       &ACP{Name: "acp-name"},
		},
		{
			Name:      "toUpdate",
			Namespace: "default",
			Domain:    "sad-bat-123.hub-traefik.io",
			Version:   "version-2",
			Service:   Service{Name: "service-2", Port: 8082},
			ACP:       &ACP{Name: "acp-name"},
		},
	}

	hashes := map[string]string{
		"toCreate": "gFO9z9bw0btZf3+lNIiaXfA8z0g=",
		"toUpdate": "4vJBrpeDJLuGzikpIg0ZJTca9FQ=",
	}

	client := newPlatformClientMock(t)
	client.OnGetWildcardCertificate().TypedReturns(Certificate{
		Certificate: []byte("cert"),
		PrivateKey:  []byte("private"),
	}, nil)

	var callCount int
	client.OnGetEdgeIngresses().
		TypedReturns(edgeIngresses, nil).
		Run(func(_ mock.Arguments) {
			callCount++
			if callCount > 1 {
				cancel()
			}
		})

	traefikClientSet := traefikkubemock.NewSimpleClientset()

	w, err := NewWatcher(client, clientSetHub, clientSet, traefikClientSet.TraefikV1alpha1(), hubInformer, WatcherConfig{
		IngressClassName:        "traefik-hub",
		TraefikTunnelEntryPoint: "traefikhub-tunl",
		AgentNamespace:          "hub-agent",
		EdgeIngressSyncInterval: time.Millisecond,
		CertRetryInterval:       time.Millisecond,
		CertSyncInterval:        time.Millisecond,
	})

	require.NoError(t, err)

	stop := make(chan struct{})
	go func() {
		w.Run(ctx)
		close(stop)
	}()

	<-stop

	_, err = clientSetHub.HubV1alpha1().
		EdgeIngresses("default").
		Get(ctx, "toDelete", metav1.GetOptions{})
	require.Error(t, err)

	for _, edgeIngress := range edgeIngresses {
		edgeIng, errL := clientSetHub.HubV1alpha1().
			EdgeIngresses(edgeIngress.Namespace).
			Get(ctx, edgeIngress.Name, metav1.GetOptions{})
		require.NoError(t, errL)

		assert.Equal(t, hubv1alpha1.EdgeIngressSpec{
			Service: hubv1alpha1.EdgeIngressService{
				Name: edgeIngress.Service.Name,
				Port: edgeIngress.Service.Port,
			},
			ACP: &hubv1alpha1.EdgeIngressACP{
				Name: edgeIngress.ACP.Name,
			},
		}, edgeIng.Spec)

		assert.WithinDuration(t, time.Now(), edgeIng.Status.SyncedAt.Time, 100*time.Millisecond)
		edgeIng.Status.SyncedAt = metav1.Time{}

		assert.Equal(t, hubv1alpha1.EdgeIngressStatus{
			Version:    edgeIngress.Version,
			SyncedAt:   metav1.Time{},
			Domain:     edgeIngress.Domain,
			URLs:       "https://" + edgeIngress.Domain,
			SpecHash:   hashes[edgeIngress.Name],
			Connection: hubv1alpha1.EdgeIngressConnectionUp,
		}, edgeIng.Status)

		// Make sure the ingress related to the edgeIngress is created.
		ctx = context.Background()
		ing, errL := clientSet.NetworkingV1().Ingresses(edgeIngress.Namespace).Get(ctx, edgeIngress.Name, metav1.GetOptions{})
		require.NoError(t, errL)

		assert.Equal(t, map[string]string{
			"app.kubernetes.io/managed-by": "traefik-hub",
		}, ing.ObjectMeta.Labels)

		assert.Equal(t, map[string]string{
			"hub.traefik.io/access-control-policy":             "acp-name",
			"traefik.ingress.kubernetes.io/router.tls":         "true",
			"traefik.ingress.kubernetes.io/router.entrypoints": "traefikhub-tunl",
		}, ing.ObjectMeta.Annotations)

		assert.Equal(t, []metav1.OwnerReference{
			{
				APIVersion: "hub.traefik.io/v1alpha1",
				Kind:       "EdgeIngress",
				Name:       edgeIng.Name,
				UID:        edgeIng.UID,
			},
		}, ing.ObjectMeta.OwnerReferences)

		wantPathType := netv1.PathTypePrefix
		assert.Equal(t, netv1.IngressSpec{
			IngressClassName: pointer.StringPtr("traefik-hub"),
			TLS: []netv1.IngressTLS{
				{
					Hosts:      []string{edgeIngress.Domain},
					SecretName: secretName,
				},
			},
			Rules: []netv1.IngressRule{
				{
					Host: edgeIngress.Domain,
					IngressRuleValue: netv1.IngressRuleValue{
						HTTP: &netv1.HTTPIngressRuleValue{
							Paths: []netv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &wantPathType,
									Backend: netv1.IngressBackend{
										Service: &netv1.IngressServiceBackend{
											Name: edgeIngress.Service.Name,
											Port: netv1.ServiceBackendPort{
												Number: int32(edgeIngress.Service.Port),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		}, ing.Spec)
	}
}

func Test_WatcherRun_appends_owner_reference(t *testing.T) {
	clientSetHub := hubkubemock.NewSimpleClientset()
	clientSet := kubemock.NewSimpleClientset()

	ctx, cancel := context.WithCancel(context.Background())
	hubInformer := hubinformer.NewSharedInformerFactory(clientSetHub, 0)

	edgeIngressInformer := hubInformer.Hub().V1alpha1().EdgeIngresses().Informer()

	hubInformer.Start(ctx.Done())
	cache.WaitForCacheSync(ctx.Done(), edgeIngressInformer.HasSynced)

	edgeIngresses := []EdgeIngress{
		{
			Name:      "toCreate",
			Namespace: "default",
			Domain:    "majestic-beaver-123.hub-traefik.io",
			Version:   "version-1",
			Service:   Service{Name: "service-1", Port: 8080},
			ACP:       &ACP{Name: "acp-name"},
		},
		{
			Name:      "toCreateAlso",
			Namespace: "default",
			Domain:    "sad-bat-123.hub-traefik.io",
			Version:   "version-2",
			Service:   Service{Name: "service-2", Port: 8082},
			ACP:       &ACP{Name: "acp-name"},
		},
	}

	client := newPlatformClientMock(t)
	client.OnGetWildcardCertificate().TypedReturns(Certificate{
		Certificate: []byte("cert"),
		PrivateKey:  []byte("private"),
	}, nil)

	var callCount int
	client.OnGetEdgeIngresses().
		TypedReturns(edgeIngresses, nil).
		Run(func(_ mock.Arguments) {
			callCount++
			if callCount > 1 {
				cancel()
			}
		})

	traefikClientSet := traefikkubemock.NewSimpleClientset()

	w, err := NewWatcher(client, clientSetHub, clientSet, traefikClientSet.TraefikV1alpha1(), hubInformer, WatcherConfig{
		IngressClassName:        "traefik-hub",
		TraefikTunnelEntryPoint: "traefikhub-tunl",
		AgentNamespace:          "hub-agent",
		EdgeIngressSyncInterval: time.Millisecond,
		CertRetryInterval:       time.Millisecond,
		CertSyncInterval:        time.Millisecond,
	})

	require.NoError(t, err)

	stop := make(chan struct{})
	go func() {
		w.Run(ctx)
		close(stop)
	}()

	<-stop

	var wantOwner []metav1.OwnerReference
	for _, edgeIngress := range edgeIngresses {
		edgeIng, errL := clientSetHub.HubV1alpha1().
			EdgeIngresses(edgeIngress.Namespace).
			Get(ctx, edgeIngress.Name, metav1.GetOptions{})
		require.NoError(t, errL)

		wantOwner = append(wantOwner, metav1.OwnerReference{
			APIVersion: "hub.traefik.io/v1alpha1",
			Kind:       "EdgeIngress",
			Name:       edgeIng.Name,
			UID:        edgeIng.UID,
		})
	}

	secret, err := clientSet.CoreV1().Secrets("default").Get(ctx, secretName, metav1.GetOptions{})
	require.NoError(t, err)

	assert.Len(t, secret.OwnerReferences, 2)
	assert.Equal(t, wantOwner, secret.OwnerReferences)
}

func Test_WatcherRun_handle_custom_domains(t *testing.T) {
	clientSetHub := hubkubemock.NewSimpleClientset(&toUpdate)
	clientSet := kubemock.NewSimpleClientset()

	ctx, cancel := context.WithCancel(context.Background())
	hubInformer := hubinformer.NewSharedInformerFactory(clientSetHub, 0)

	edgeIngressInformer := hubInformer.Hub().V1alpha1().EdgeIngresses().Informer()

	hubInformer.Start(ctx.Done())
	cache.WaitForCacheSync(ctx.Done(), edgeIngressInformer.HasSynced)

	edgeIngresses := []EdgeIngress{
		{
			Name:      "toUpdate",
			Namespace: "default",
			Domain:    "sad-bat-123.hub-traefik.io",
			Version:   "version-2",
			Service:   Service{Name: "service-2", Port: 8082},
			ACP:       &ACP{Name: "acp-name"},
			CustomDomains: []CustomDomain{
				{
					Name:     "customDomain.com",
					Verified: true,
				},
				{
					Name: "unverified.com",
				},
			},
		},
	}

	client := newPlatformClientMock(t).
		OnGetWildcardCertificate().TypedReturns(
		Certificate{
			Certificate: []byte("cert"),
			PrivateKey:  []byte("private"),
		}, nil).
		OnGetCertificateByDomains([]string{"customDomain.com"}).TypedReturns(
		Certificate{
			Certificate: []byte("cert"),
			PrivateKey:  []byte("private"),
		}, nil).
		Parent

	var callCount int
	client.OnGetEdgeIngresses().
		TypedReturns(edgeIngresses, nil).
		Run(func(_ mock.Arguments) {
			callCount++
			if callCount > 1 {
				cancel()
			}
		})

	traefikClientSet := traefikkubemock.NewSimpleClientset()

	w, err := NewWatcher(client, clientSetHub, clientSet, traefikClientSet.TraefikV1alpha1(), hubInformer, WatcherConfig{
		IngressClassName:        "traefik-hub",
		TraefikTunnelEntryPoint: "traefikhub-tunl",
		AgentNamespace:          "hub-agent",
		EdgeIngressSyncInterval: time.Millisecond,
		CertRetryInterval:       time.Millisecond,
		CertSyncInterval:        time.Millisecond,
	})
	require.NoError(t, err)

	stop := make(chan struct{})
	go func() {
		w.Run(ctx)
		close(stop)
	}()

	<-stop

	wantEdgeIngress := edgeIngresses[0]

	edgeIng, errL := clientSetHub.HubV1alpha1().
		EdgeIngresses(wantEdgeIngress.Namespace).
		Get(ctx, wantEdgeIngress.Name, metav1.GetOptions{})
	require.NoError(t, errL)

	assert.Equal(t, hubv1alpha1.EdgeIngressSpec{
		Service: hubv1alpha1.EdgeIngressService{
			Name: wantEdgeIngress.Service.Name,
			Port: wantEdgeIngress.Service.Port,
		},
		ACP: &hubv1alpha1.EdgeIngressACP{
			Name: wantEdgeIngress.ACP.Name,
		},
		CustomDomains: []string{"customDomain.com", "unverified.com"},
	}, edgeIng.Spec)

	assert.WithinDuration(t, time.Now(), edgeIng.Status.SyncedAt.Time, 100*time.Millisecond)
	edgeIng.Status.SyncedAt = metav1.Time{}

	assert.Equal(t, hubv1alpha1.EdgeIngressStatus{
		Version:       wantEdgeIngress.Version,
		SyncedAt:      metav1.Time{},
		Domain:        wantEdgeIngress.Domain,
		CustomDomains: []string{"customDomain.com"},
		URLs:          "https://customDomain.com,https://" + wantEdgeIngress.Domain,
		SpecHash:      "OxYSOU0yEUcLM1RnjLL83wymkUU=",
		Connection:    hubv1alpha1.EdgeIngressConnectionUp,
	}, edgeIng.Status)

	// Make sure secret related to the edgeIngress is created.
	ctx = context.Background()
	secret, err := clientSet.CoreV1().Secrets(wantEdgeIngress.Namespace).Get(ctx, secretCustomDomainsName+"-"+wantEdgeIngress.Name, metav1.GetOptions{})
	require.NoError(t, err)

	wantOwner := metav1.OwnerReference{
		APIVersion: "hub.traefik.io/v1alpha1",
		Kind:       "EdgeIngress",
		Name:       edgeIng.Name,
		UID:        edgeIng.UID,
	}
	assert.Len(t, secret.OwnerReferences, 1)
	assert.Equal(t, wantOwner, secret.OwnerReferences[0])

	// Make sure the ingress related to the edgeIngress is created.
	ing, errL := clientSet.NetworkingV1().Ingresses(wantEdgeIngress.Namespace).Get(ctx, wantEdgeIngress.Name, metav1.GetOptions{})
	require.NoError(t, errL)

	assert.Equal(t, map[string]string{
		"app.kubernetes.io/managed-by": "traefik-hub",
	}, ing.ObjectMeta.Labels)

	assert.Equal(t, map[string]string{
		"hub.traefik.io/access-control-policy":             "acp-name",
		"traefik.ingress.kubernetes.io/router.tls":         "true",
		"traefik.ingress.kubernetes.io/router.entrypoints": "traefikhub-tunl",
	}, ing.ObjectMeta.Annotations)

	assert.Equal(t, []metav1.OwnerReference{
		{
			APIVersion: "hub.traefik.io/v1alpha1",
			Kind:       "EdgeIngress",
			Name:       edgeIng.Name,
			UID:        edgeIng.UID,
		},
	}, ing.ObjectMeta.OwnerReferences)

	wantPathType := netv1.PathTypePrefix
	assert.Equal(t, netv1.IngressSpec{
		IngressClassName: pointer.StringPtr("traefik-hub"),
		TLS: []netv1.IngressTLS{
			{
				Hosts:      []string{wantEdgeIngress.Domain},
				SecretName: secretName,
			},
			{
				SecretName: secretCustomDomainsName + "-" + wantEdgeIngress.Name,
				Hosts:      []string{"customDomain.com"},
			},
		},
		Rules: []netv1.IngressRule{
			{
				Host: wantEdgeIngress.Domain,
				IngressRuleValue: netv1.IngressRuleValue{
					HTTP: &netv1.HTTPIngressRuleValue{
						Paths: []netv1.HTTPIngressPath{
							{
								Path:     "/",
								PathType: &wantPathType,
								Backend: netv1.IngressBackend{
									Service: &netv1.IngressServiceBackend{
										Name: wantEdgeIngress.Service.Name,
										Port: netv1.ServiceBackendPort{
											Number: int32(wantEdgeIngress.Service.Port),
										},
									},
								},
							},
						},
					},
				},
			},
			{
				Host: wantEdgeIngress.CustomDomains[0].Name,
				IngressRuleValue: netv1.IngressRuleValue{
					HTTP: &netv1.HTTPIngressRuleValue{
						Paths: []netv1.HTTPIngressPath{
							{
								Path:     "/",
								PathType: &wantPathType,
								Backend: netv1.IngressBackend{
									Service: &netv1.IngressServiceBackend{
										Name: wantEdgeIngress.Service.Name,
										Port: netv1.ServiceBackendPort{
											Number: int32(wantEdgeIngress.Service.Port),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, ing.Spec)
}

func Test_WatcherRun_sync_certificates(t *testing.T) {
	clientSetHub := hubkubemock.NewSimpleClientset()
	clientSet := kubemock.NewSimpleClientset()

	ctx, cancel := context.WithCancel(context.Background())
	hubInformer := hubinformer.NewSharedInformerFactory(clientSetHub, 0)

	edgeIngressInformer := hubInformer.Hub().V1alpha1().EdgeIngresses().Informer()

	hubInformer.Start(ctx.Done())
	cache.WaitForCacheSync(ctx.Done(), edgeIngressInformer.HasSynced)

	edgeIngresses := []EdgeIngress{
		{
			Name:      "toCreate",
			Namespace: "service",
			Domain:    "majestic-beaver-123.hub-traefik.io",
			Version:   "version-1",
			Service:   Service{Name: "service-1", Port: 8080},
			ACP:       &ACP{Name: "acp-name"},
		},
		{
			Name:      "toCreateAlso",
			Namespace: "service",
			Domain:    "sad-bat-123.hub-traefik.io",
			Version:   "version-2",
			CustomDomains: []CustomDomain{
				{
					Name:     "customdomain.com",
					Verified: true,
				},
			},
			Service: Service{Name: "service-2", Port: 8082},
			ACP:     &ACP{Name: "acp-name"},
		},
	}

	var callCountWildVard int
	var callCountCustom int
	client := newPlatformClientMock(t).
		OnGetWildcardCertificate().
		Run(
			func(_ mock.Arguments) {
				callCountWildVard++
				if callCountWildVard > 2 {
					cancel()
				}
			}).
		ReturnsFn(func() (Certificate, error) {
			if callCountWildVard == 0 {
				return Certificate{
					Certificate: []byte("cert"),
					PrivateKey:  []byte("private"),
				}, nil
			}

			return Certificate{
				Certificate: []byte("certRefresh"),
				PrivateKey:  []byte("private"),
			}, nil
		}).
		OnGetCertificateByDomains([]string{"customdomain.com"}).
		Run(
			func(_ mock.Arguments) {
				callCountCustom++
			}).
		ReturnsFn(
			func(strings []string) (Certificate, error) {
				if callCountCustom == 0 {
					return Certificate{
						Certificate: []byte("custom"),
						PrivateKey:  []byte("private"),
					}, nil
				}

				return Certificate{
					Certificate: []byte("customRefresh"),
					PrivateKey:  []byte("private"),
				}, nil
			}).
		OnGetEdgeIngresses().
		TypedReturns(edgeIngresses, nil).Parent

	traefikClientSet := traefikkubemock.NewSimpleClientset()

	w, err := NewWatcher(client, clientSetHub, clientSet, traefikClientSet.TraefikV1alpha1(), hubInformer, WatcherConfig{
		IngressClassName:        "traefik-hub",
		TraefikTunnelEntryPoint: "traefikhub-tunl",
		AgentNamespace:          "default",
		EdgeIngressSyncInterval: time.Millisecond,
		CertRetryInterval:       time.Millisecond,
		CertSyncInterval:        time.Millisecond,
	})

	require.NoError(t, err)

	stop := make(chan struct{})
	go func() {
		w.Run(ctx)
		close(stop)
	}()

	<-stop

	secret, err := clientSet.CoreV1().Secrets("default").Get(ctx, secretName, metav1.GetOptions{})
	require.NoError(t, err)
	assert.Equal(t, []byte("certRefresh"), secret.Data["tls.crt"])

	assert.Len(t, secret.OwnerReferences, 0)

	secret, err = clientSet.CoreV1().Secrets("service").Get(ctx, secretName, metav1.GetOptions{})
	require.NoError(t, err)
	assert.Equal(t, []byte("certRefresh"), secret.Data["tls.crt"])
	assert.Len(t, secret.OwnerReferences, 2)

	var wantOwner []metav1.OwnerReference
	for _, edgeIngress := range edgeIngresses {
		edgeIng, errL := clientSetHub.HubV1alpha1().
			EdgeIngresses(edgeIngress.Namespace).
			Get(ctx, edgeIngress.Name, metav1.GetOptions{})
		require.NoError(t, errL)

		wantOwner = append(wantOwner, metav1.OwnerReference{
			APIVersion: "hub.traefik.io/v1alpha1",
			Kind:       "EdgeIngress",
			Name:       edgeIng.Name,
			UID:        edgeIng.UID,
		})
	}

	assert.Equal(t, wantOwner, secret.OwnerReferences)

	secret, err = clientSet.CoreV1().Secrets("service").Get(ctx, secretCustomDomainsName+"-toCreateAlso", metav1.GetOptions{})
	require.NoError(t, err)
	assert.Equal(t, []byte("customRefresh"), secret.Data["tls.crt"])
	assert.Len(t, secret.OwnerReferences, 1)
}

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/traefik/hub-agent-kubernetes/pkg/crd/generated/client/hub/clientset/versioned/typed/hub/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHubV1alpha1 struct {
	*testing.Fake
}

func (c *FakeHubV1alpha1) APIs(namespace string) v1alpha1.APIInterface {
	return &FakeAPIs{c, namespace}
}

func (c *FakeHubV1alpha1) APIGroups() v1alpha1.APIGroupInterface {
	return &FakeAPIGroups{c}
}

func (c *FakeHubV1alpha1) APIPortals() v1alpha1.APIPortalInterface {
	return &FakeAPIPortals{c}
}

func (c *FakeHubV1alpha1) AccessControlPolicies() v1alpha1.AccessControlPolicyInterface {
	return &FakeAccessControlPolicies{c}
}

func (c *FakeHubV1alpha1) EdgeIngresses(namespace string) v1alpha1.EdgeIngressInterface {
	return &FakeEdgeIngresses{c, namespace}
}

func (c *FakeHubV1alpha1) IngressClasses() v1alpha1.IngressClassInterface {
	return &FakeIngressClasses{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHubV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

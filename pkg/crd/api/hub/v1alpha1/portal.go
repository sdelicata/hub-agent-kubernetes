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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Portal defines a portal.
// +kubebuilder:printcolumn:name="URLs",type=string,JSONPath=`.status.urls`
// +kubebuilder:printcolumn:name="APIURLs",type=string,JSONPath=`.status.apiUrls`
// +kubebuilder:resource:scope=Cluster
type Portal struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// The desired behavior of this portal.
	Spec PortalSpec `json:"spec,omitempty"`

	// The current status of this portal.
	// +optional
	Status PortalStatus `json:"status,omitempty"`
}

// PortalSpec configures a Portal.
type PortalSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// CustomDomains are the custom domains under which the portal will be exposed.
	// +optional
	CustomDomains []string `json:"customDomains,omitempty"`
	// APICustomDomains are the custom domains under which the API will be exposed.
	// +optional
	APICustomDomains []string `json:"apiCustomDomains,omitempty"`
}

// PortalStatus is the status of a Portal.
type PortalStatus struct {
	Version  string      `json:"version,omitempty"`
	SyncedAt metav1.Time `json:"syncedAt,omitempty"`

	// URLs are the URLs for accessing the portal WebUI.
	URLs string `json:"urls"`

	// APIURLs are the URLs for accessing the portal API.
	APIURLs string `json:"apiUrls"`

	// HubDomain is the hub generated domain of the portal WebUI.
	// +optional
	HubDomain string `json:"hubDomain"`

	// APIHubDomain is the hub generated domain of the portal API.
	// +optional
	APIHubDomain string `json:"apiHubDomain"`

	// CustomDomains are the custom domains for accessing the exposed portal WebUI.
	// +optional
	CustomDomains []string `json:"customDomains,omitempty"`

	// APICustomDomains are the custom domains for accessing the exposed apis.
	// +optional
	APICustomDomains []string `json:"apiCustomDomains,omitempty"`

	// Hash is a hash representing the Portal.
	Hash string `json:"hash,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PortalList defines a list of portals.
type PortalList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Portal `json:"items"`
}

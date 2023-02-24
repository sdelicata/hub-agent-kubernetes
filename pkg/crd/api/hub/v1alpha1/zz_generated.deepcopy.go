//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *API) DeepCopyInto(out *API) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new API.
func (in *API) DeepCopy() *API {
	if in == nil {
		return nil
	}
	out := new(API)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *API) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIGroup) DeepCopyInto(out *APIGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIGroup.
func (in *APIGroup) DeepCopy() *APIGroup {
	if in == nil {
		return nil
	}
	out := new(APIGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIGroupList) DeepCopyInto(out *APIGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIGroupList.
func (in *APIGroupList) DeepCopy() *APIGroupList {
	if in == nil {
		return nil
	}
	out := new(APIGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIGroupSpec) DeepCopyInto(out *APIGroupSpec) {
	*out = *in
	in.APISelector.DeepCopyInto(&out.APISelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIGroupSpec.
func (in *APIGroupSpec) DeepCopy() *APIGroupSpec {
	if in == nil {
		return nil
	}
	out := new(APIGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIGroupStatus) DeepCopyInto(out *APIGroupStatus) {
	*out = *in
	in.SyncedAt.DeepCopyInto(&out.SyncedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIGroupStatus.
func (in *APIGroupStatus) DeepCopy() *APIGroupStatus {
	if in == nil {
		return nil
	}
	out := new(APIGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIList) DeepCopyInto(out *APIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]API, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIList.
func (in *APIList) DeepCopy() *APIList {
	if in == nil {
		return nil
	}
	out := new(APIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPortal) DeepCopyInto(out *APIPortal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPortal.
func (in *APIPortal) DeepCopy() *APIPortal {
	if in == nil {
		return nil
	}
	out := new(APIPortal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIPortal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPortalList) DeepCopyInto(out *APIPortalList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIPortal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPortalList.
func (in *APIPortalList) DeepCopy() *APIPortalList {
	if in == nil {
		return nil
	}
	out := new(APIPortalList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIPortalList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPortalSpec) DeepCopyInto(out *APIPortalSpec) {
	*out = *in
	if in.CustomDomains != nil {
		in, out := &in.CustomDomains, &out.CustomDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APICustomDomains != nil {
		in, out := &in.APICustomDomains, &out.APICustomDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPortalSpec.
func (in *APIPortalSpec) DeepCopy() *APIPortalSpec {
	if in == nil {
		return nil
	}
	out := new(APIPortalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPortalStatus) DeepCopyInto(out *APIPortalStatus) {
	*out = *in
	in.SyncedAt.DeepCopyInto(&out.SyncedAt)
	if in.CustomDomains != nil {
		in, out := &in.CustomDomains, &out.CustomDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APICustomDomains != nil {
		in, out := &in.APICustomDomains, &out.APICustomDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPortalStatus.
func (in *APIPortalStatus) DeepCopy() *APIPortalStatus {
	if in == nil {
		return nil
	}
	out := new(APIPortalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIService) DeepCopyInto(out *APIService) {
	*out = *in
	out.Port = in.Port
	out.OpenAPISpec = in.OpenAPISpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIService.
func (in *APIService) DeepCopy() *APIService {
	if in == nil {
		return nil
	}
	out := new(APIService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceBackendPort) DeepCopyInto(out *APIServiceBackendPort) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceBackendPort.
func (in *APIServiceBackendPort) DeepCopy() *APIServiceBackendPort {
	if in == nil {
		return nil
	}
	out := new(APIServiceBackendPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISpec) DeepCopyInto(out *APISpec) {
	*out = *in
	out.Service = in.Service
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISpec.
func (in *APISpec) DeepCopy() *APISpec {
	if in == nil {
		return nil
	}
	out := new(APISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIStatus) DeepCopyInto(out *APIStatus) {
	*out = *in
	in.SyncedAt.DeepCopyInto(&out.SyncedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIStatus.
func (in *APIStatus) DeepCopy() *APIStatus {
	if in == nil {
		return nil
	}
	out := new(APIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlOAuthIntro) DeepCopyInto(out *AccessControlOAuthIntro) {
	*out = *in
	in.ClientConfig.DeepCopyInto(&out.ClientConfig)
	out.TokenSource = in.TokenSource
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlOAuthIntro.
func (in *AccessControlOAuthIntro) DeepCopy() *AccessControlOAuthIntro {
	if in == nil {
		return nil
	}
	out := new(AccessControlOAuthIntro)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlOAuthIntroClientConfig) DeepCopyInto(out *AccessControlOAuthIntroClientConfig) {
	*out = *in
	in.HTTPClientConfig.DeepCopyInto(&out.HTTPClientConfig)
	out.Auth = in.Auth
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlOAuthIntroClientConfig.
func (in *AccessControlOAuthIntroClientConfig) DeepCopy() *AccessControlOAuthIntroClientConfig {
	if in == nil {
		return nil
	}
	out := new(AccessControlOAuthIntroClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlOAuthIntroClientConfigAuth) DeepCopyInto(out *AccessControlOAuthIntroClientConfigAuth) {
	*out = *in
	out.Secret = in.Secret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlOAuthIntroClientConfigAuth.
func (in *AccessControlOAuthIntroClientConfigAuth) DeepCopy() *AccessControlOAuthIntroClientConfigAuth {
	if in == nil {
		return nil
	}
	out := new(AccessControlOAuthIntroClientConfigAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicy) DeepCopyInto(out *AccessControlPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicy.
func (in *AccessControlPolicy) DeepCopy() *AccessControlPolicy {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessControlPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyAPIKey) DeepCopyInto(out *AccessControlPolicyAPIKey) {
	*out = *in
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]AccessControlPolicyAPIKeyKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyAPIKey.
func (in *AccessControlPolicyAPIKey) DeepCopy() *AccessControlPolicyAPIKey {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyAPIKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyAPIKeyKey) DeepCopyInto(out *AccessControlPolicyAPIKeyKey) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyAPIKeyKey.
func (in *AccessControlPolicyAPIKeyKey) DeepCopy() *AccessControlPolicyAPIKeyKey {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyAPIKeyKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyBasicAuth) DeepCopyInto(out *AccessControlPolicyBasicAuth) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyBasicAuth.
func (in *AccessControlPolicyBasicAuth) DeepCopy() *AccessControlPolicyBasicAuth {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyBasicAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyJWT) DeepCopyInto(out *AccessControlPolicyJWT) {
	*out = *in
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyJWT.
func (in *AccessControlPolicyJWT) DeepCopy() *AccessControlPolicyJWT {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyJWT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyList) DeepCopyInto(out *AccessControlPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessControlPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyList.
func (in *AccessControlPolicyList) DeepCopy() *AccessControlPolicyList {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessControlPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyOIDC) DeepCopyInto(out *AccessControlPolicyOIDC) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.AuthParams != nil {
		in, out := &in.AuthParams, &out.AuthParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StateCookie != nil {
		in, out := &in.StateCookie, &out.StateCookie
		*out = new(StateCookie)
		**out = **in
	}
	if in.Session != nil {
		in, out := &in.Session, &out.Session
		*out = new(Session)
		(*in).DeepCopyInto(*out)
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyOIDC.
func (in *AccessControlPolicyOIDC) DeepCopy() *AccessControlPolicyOIDC {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyOIDC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyOIDCGoogle) DeepCopyInto(out *AccessControlPolicyOIDCGoogle) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.AuthParams != nil {
		in, out := &in.AuthParams, &out.AuthParams
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.StateCookie != nil {
		in, out := &in.StateCookie, &out.StateCookie
		*out = new(StateCookie)
		**out = **in
	}
	if in.Session != nil {
		in, out := &in.Session, &out.Session
		*out = new(Session)
		(*in).DeepCopyInto(*out)
	}
	if in.ForwardHeaders != nil {
		in, out := &in.ForwardHeaders, &out.ForwardHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Emails != nil {
		in, out := &in.Emails, &out.Emails
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyOIDCGoogle.
func (in *AccessControlPolicyOIDCGoogle) DeepCopy() *AccessControlPolicyOIDCGoogle {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyOIDCGoogle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicySpec) DeepCopyInto(out *AccessControlPolicySpec) {
	*out = *in
	if in.JWT != nil {
		in, out := &in.JWT, &out.JWT
		*out = new(AccessControlPolicyJWT)
		(*in).DeepCopyInto(*out)
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(AccessControlPolicyBasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = new(AccessControlPolicyAPIKey)
		(*in).DeepCopyInto(*out)
	}
	if in.OIDC != nil {
		in, out := &in.OIDC, &out.OIDC
		*out = new(AccessControlPolicyOIDC)
		(*in).DeepCopyInto(*out)
	}
	if in.OIDCGoogle != nil {
		in, out := &in.OIDCGoogle, &out.OIDCGoogle
		*out = new(AccessControlPolicyOIDCGoogle)
		(*in).DeepCopyInto(*out)
	}
	if in.OAuthIntro != nil {
		in, out := &in.OAuthIntro, &out.OAuthIntro
		*out = new(AccessControlOAuthIntro)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicySpec.
func (in *AccessControlPolicySpec) DeepCopy() *AccessControlPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessControlPolicyStatus) DeepCopyInto(out *AccessControlPolicyStatus) {
	*out = *in
	in.SyncedAt.DeepCopyInto(&out.SyncedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessControlPolicyStatus.
func (in *AccessControlPolicyStatus) DeepCopy() *AccessControlPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AccessControlPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeIngress) DeepCopyInto(out *EdgeIngress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeIngress.
func (in *EdgeIngress) DeepCopy() *EdgeIngress {
	if in == nil {
		return nil
	}
	out := new(EdgeIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeIngress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeIngressACP) DeepCopyInto(out *EdgeIngressACP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeIngressACP.
func (in *EdgeIngressACP) DeepCopy() *EdgeIngressACP {
	if in == nil {
		return nil
	}
	out := new(EdgeIngressACP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeIngressList) DeepCopyInto(out *EdgeIngressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EdgeIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeIngressList.
func (in *EdgeIngressList) DeepCopy() *EdgeIngressList {
	if in == nil {
		return nil
	}
	out := new(EdgeIngressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EdgeIngressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeIngressService) DeepCopyInto(out *EdgeIngressService) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeIngressService.
func (in *EdgeIngressService) DeepCopy() *EdgeIngressService {
	if in == nil {
		return nil
	}
	out := new(EdgeIngressService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeIngressSpec) DeepCopyInto(out *EdgeIngressSpec) {
	*out = *in
	out.Service = in.Service
	if in.ACP != nil {
		in, out := &in.ACP, &out.ACP
		*out = new(EdgeIngressACP)
		**out = **in
	}
	if in.CustomDomains != nil {
		in, out := &in.CustomDomains, &out.CustomDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeIngressSpec.
func (in *EdgeIngressSpec) DeepCopy() *EdgeIngressSpec {
	if in == nil {
		return nil
	}
	out := new(EdgeIngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EdgeIngressStatus) DeepCopyInto(out *EdgeIngressStatus) {
	*out = *in
	in.SyncedAt.DeepCopyInto(&out.SyncedAt)
	if in.CustomDomains != nil {
		in, out := &in.CustomDomains, &out.CustomDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EdgeIngressStatus.
func (in *EdgeIngressStatus) DeepCopy() *EdgeIngressStatus {
	if in == nil {
		return nil
	}
	out := new(EdgeIngressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPClientConfig) DeepCopyInto(out *HTTPClientConfig) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(HTTPClientConfigTLS)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPClientConfig.
func (in *HTTPClientConfig) DeepCopy() *HTTPClientConfig {
	if in == nil {
		return nil
	}
	out := new(HTTPClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPClientConfigTLS) DeepCopyInto(out *HTTPClientConfigTLS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPClientConfigTLS.
func (in *HTTPClientConfigTLS) DeepCopy() *HTTPClientConfigTLS {
	if in == nil {
		return nil
	}
	out := new(HTTPClientConfigTLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressClass) DeepCopyInto(out *IngressClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressClass.
func (in *IngressClass) DeepCopy() *IngressClass {
	if in == nil {
		return nil
	}
	out := new(IngressClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressClassList) DeepCopyInto(out *IngressClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IngressClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressClassList.
func (in *IngressClassList) DeepCopy() *IngressClassList {
	if in == nil {
		return nil
	}
	out := new(IngressClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IngressClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressClassSpec) DeepCopyInto(out *IngressClassSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressClassSpec.
func (in *IngressClassSpec) DeepCopy() *IngressClassSpec {
	if in == nil {
		return nil
	}
	out := new(IngressClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenAPISpec) DeepCopyInto(out *OpenAPISpec) {
	*out = *in
	out.Port = in.Port
	out.Schema = in.Schema
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenAPISpec.
func (in *OpenAPISpec) DeepCopy() *OpenAPISpec {
	if in == nil {
		return nil
	}
	out := new(OpenAPISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Session) DeepCopyInto(out *Session) {
	*out = *in
	if in.Refresh != nil {
		in, out := &in.Refresh, &out.Refresh
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Session.
func (in *Session) DeepCopy() *Session {
	if in == nil {
		return nil
	}
	out := new(Session)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateCookie) DeepCopyInto(out *StateCookie) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateCookie.
func (in *StateCookie) DeepCopy() *StateCookie {
	if in == nil {
		return nil
	}
	out := new(StateCookie)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenSource) DeepCopyInto(out *TokenSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenSource.
func (in *TokenSource) DeepCopy() *TokenSource {
	if in == nil {
		return nil
	}
	out := new(TokenSource)
	in.DeepCopyInto(out)
	return out
}

//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BPFSocketLBHostnsOnly) DeepCopyInto(out *BPFSocketLBHostnsOnly) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BPFSocketLBHostnsOnly.
func (in *BPFSocketLBHostnsOnly) DeepCopy() *BPFSocketLBHostnsOnly {
	if in == nil {
		return nil
	}
	out := new(BPFSocketLBHostnsOnly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressGateway) DeepCopyInto(out *EgressGateway) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressGateway.
func (in *EgressGateway) DeepCopy() *EgressGateway {
	if in == nil {
		return nil
	}
	out := new(EgressGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hubble) DeepCopyInto(out *Hubble) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hubble.
func (in *Hubble) DeepCopy() *Hubble {
	if in == nil {
		return nil
	}
	out := new(Hubble)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv6) DeepCopyInto(out *IPv6) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv6.
func (in *IPv6) DeepCopy() *IPv6 {
	if in == nil {
		return nil
	}
	out := new(IPv6)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxy) DeepCopyInto(out *KubeProxy) {
	*out = *in
	if in.ServiceHost != nil {
		in, out := &in.ServiceHost, &out.ServiceHost
		*out = new(string)
		**out = **in
	}
	if in.ServicePort != nil {
		in, out := &in.ServicePort, &out.ServicePort
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxy.
func (in *KubeProxy) DeepCopy() *KubeProxy {
	if in == nil {
		return nil
	}
	out := new(KubeProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig) DeepCopyInto(out *NetworkConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Debug != nil {
		in, out := &in.Debug, &out.Debug
		*out = new(bool)
		**out = **in
	}
	if in.PSPEnabled != nil {
		in, out := &in.PSPEnabled, &out.PSPEnabled
		*out = new(bool)
		**out = **in
	}
	if in.KubeProxy != nil {
		in, out := &in.KubeProxy, &out.KubeProxy
		*out = new(KubeProxy)
		(*in).DeepCopyInto(*out)
	}
	if in.Hubble != nil {
		in, out := &in.Hubble, &out.Hubble
		*out = new(Hubble)
		**out = **in
	}
	if in.TunnelMode != nil {
		in, out := &in.TunnelMode, &out.TunnelMode
		*out = new(TunnelMode)
		**out = **in
	}
	if in.Store != nil {
		in, out := &in.Store, &out.Store
		*out = new(Store)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(IPv6)
		**out = **in
	}
	if in.BPFSocketLBHostnsOnly != nil {
		in, out := &in.BPFSocketLBHostnsOnly, &out.BPFSocketLBHostnsOnly
		*out = new(BPFSocketLBHostnsOnly)
		**out = **in
	}
	if in.EgressGateway != nil {
		in, out := &in.EgressGateway, &out.EgressGateway
		*out = new(EgressGateway)
		**out = **in
	}
	if in.MTU != nil {
		in, out := &in.MTU, &out.MTU
		*out = new(int)
		**out = **in
	}
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LoadBalancingMode != nil {
		in, out := &in.LoadBalancingMode, &out.LoadBalancingMode
		*out = new(LoadBalancingMode)
		**out = **in
	}
	if in.IPv4NativeRoutingCIDREnabled != nil {
		in, out := &in.IPv4NativeRoutingCIDREnabled, &out.IPv4NativeRoutingCIDREnabled
		*out = new(bool)
		**out = **in
	}
	if in.Overlay != nil {
		in, out := &in.Overlay, &out.Overlay
		*out = new(Overlay)
		(*in).DeepCopyInto(*out)
	}
	if in.SnatToUpstreamDNS != nil {
		in, out := &in.SnatToUpstreamDNS, &out.SnatToUpstreamDNS
		*out = new(SnatToUpstreamDNS)
		**out = **in
	}
	if in.SnatOutOfCluster != nil {
		in, out := &in.SnatOutOfCluster, &out.SnatOutOfCluster
		*out = new(SnatOutOfCluster)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfig.
func (in *NetworkConfig) DeepCopy() *NetworkConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nodeport) DeepCopyInto(out *Nodeport) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nodeport.
func (in *Nodeport) DeepCopy() *Nodeport {
	if in == nil {
		return nil
	}
	out := new(Nodeport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Overlay) DeepCopyInto(out *Overlay) {
	*out = *in
	if in.CreatePodRoutes != nil {
		in, out := &in.CreatePodRoutes, &out.CreatePodRoutes
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Overlay.
func (in *Overlay) DeepCopy() *Overlay {
	if in == nil {
		return nil
	}
	out := new(Overlay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatOutOfCluster) DeepCopyInto(out *SnatOutOfCluster) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatOutOfCluster.
func (in *SnatOutOfCluster) DeepCopy() *SnatOutOfCluster {
	if in == nil {
		return nil
	}
	out := new(SnatOutOfCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatToUpstreamDNS) DeepCopyInto(out *SnatToUpstreamDNS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatToUpstreamDNS.
func (in *SnatToUpstreamDNS) DeepCopy() *SnatToUpstreamDNS {
	if in == nil {
		return nil
	}
	out := new(SnatToUpstreamDNS)
	in.DeepCopyInto(out)
	return out
}

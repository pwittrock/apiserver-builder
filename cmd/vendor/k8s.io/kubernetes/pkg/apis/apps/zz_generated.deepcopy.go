// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package apps

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	reflect "reflect"
)

// Deprecated: register deep-copy functions.
func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// Deprecated: RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ControllerRevision).DeepCopyInto(out.(*ControllerRevision))
			return nil
		}, InType: reflect.TypeOf(&ControllerRevision{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ControllerRevisionList).DeepCopyInto(out.(*ControllerRevisionList))
			return nil
		}, InType: reflect.TypeOf(&ControllerRevisionList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RollingUpdateStatefulSetStrategy).DeepCopyInto(out.(*RollingUpdateStatefulSetStrategy))
			return nil
		}, InType: reflect.TypeOf(&RollingUpdateStatefulSetStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StatefulSet).DeepCopyInto(out.(*StatefulSet))
			return nil
		}, InType: reflect.TypeOf(&StatefulSet{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StatefulSetList).DeepCopyInto(out.(*StatefulSetList))
			return nil
		}, InType: reflect.TypeOf(&StatefulSetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StatefulSetSpec).DeepCopyInto(out.(*StatefulSetSpec))
			return nil
		}, InType: reflect.TypeOf(&StatefulSetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StatefulSetStatus).DeepCopyInto(out.(*StatefulSetStatus))
			return nil
		}, InType: reflect.TypeOf(&StatefulSetStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StatefulSetUpdateStrategy).DeepCopyInto(out.(*StatefulSetUpdateStrategy))
			return nil
		}, InType: reflect.TypeOf(&StatefulSetUpdateStrategy{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRevision) DeepCopyInto(out *ControllerRevision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Data == nil {
		out.Data = nil
	} else {
		out.Data = in.Data.DeepCopyObject()
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRevision.
func (x *ControllerRevision) DeepCopy() *ControllerRevision {
	if x == nil {
		return nil
	}
	out := new(ControllerRevision)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *ControllerRevision) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRevisionList) DeepCopyInto(out *ControllerRevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerRevision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRevisionList.
func (x *ControllerRevisionList) DeepCopy() *ControllerRevisionList {
	if x == nil {
		return nil
	}
	out := new(ControllerRevisionList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *ControllerRevisionList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateStatefulSetStrategy) DeepCopyInto(out *RollingUpdateStatefulSetStrategy) {
	*out = *in
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateStatefulSetStrategy.
func (x *RollingUpdateStatefulSetStrategy) DeepCopy() *RollingUpdateStatefulSetStrategy {
	if x == nil {
		return nil
	}
	out := new(RollingUpdateStatefulSetStrategy)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSet) DeepCopyInto(out *StatefulSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSet.
func (x *StatefulSet) DeepCopy() *StatefulSet {
	if x == nil {
		return nil
	}
	out := new(StatefulSet)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *StatefulSet) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetList) DeepCopyInto(out *StatefulSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StatefulSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetList.
func (x *StatefulSetList) DeepCopy() *StatefulSetList {
	if x == nil {
		return nil
	}
	out := new(StatefulSetList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *StatefulSetList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetSpec) DeepCopyInto(out *StatefulSetSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]api.PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetSpec.
func (x *StatefulSetSpec) DeepCopy() *StatefulSetSpec {
	if x == nil {
		return nil
	}
	out := new(StatefulSetSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetStatus) DeepCopyInto(out *StatefulSetStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetStatus.
func (x *StatefulSetStatus) DeepCopy() *StatefulSetStatus {
	if x == nil {
		return nil
	}
	out := new(StatefulSetStatus)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetUpdateStrategy) DeepCopyInto(out *StatefulSetUpdateStrategy) {
	*out = *in
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		if *in == nil {
			*out = nil
		} else {
			*out = new(RollingUpdateStatefulSetStrategy)
			**out = **in
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetUpdateStrategy.
func (x *StatefulSetUpdateStrategy) DeepCopy() *StatefulSetUpdateStrategy {
	if x == nil {
		return nil
	}
	out := new(StatefulSetUpdateStrategy)
	x.DeepCopyInto(out)
	return out
}
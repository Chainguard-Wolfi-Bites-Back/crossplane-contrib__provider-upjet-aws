//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleConfigInitParameters) DeepCopyInto(out *RuleConfigInitParameters) {
	*out = *in
	if in.Inverted != nil {
		in, out := &in.Inverted, &out.Inverted
		*out = new(bool)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleConfigInitParameters.
func (in *RuleConfigInitParameters) DeepCopy() *RuleConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(RuleConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleConfigObservation) DeepCopyInto(out *RuleConfigObservation) {
	*out = *in
	if in.Inverted != nil {
		in, out := &in.Inverted, &out.Inverted
		*out = new(bool)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleConfigObservation.
func (in *RuleConfigObservation) DeepCopy() *RuleConfigObservation {
	if in == nil {
		return nil
	}
	out := new(RuleConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleConfigParameters) DeepCopyInto(out *RuleConfigParameters) {
	*out = *in
	if in.Inverted != nil {
		in, out := &in.Inverted, &out.Inverted
		*out = new(bool)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleConfigParameters.
func (in *RuleConfigParameters) DeepCopy() *RuleConfigParameters {
	if in == nil {
		return nil
	}
	out := new(RuleConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SafetyRule) DeepCopyInto(out *SafetyRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SafetyRule.
func (in *SafetyRule) DeepCopy() *SafetyRule {
	if in == nil {
		return nil
	}
	out := new(SafetyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SafetyRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SafetyRuleInitParameters) DeepCopyInto(out *SafetyRuleInitParameters) {
	*out = *in
	if in.AssertedControls != nil {
		in, out := &in.AssertedControls, &out.AssertedControls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AssertedControlsRefs != nil {
		in, out := &in.AssertedControlsRefs, &out.AssertedControlsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AssertedControlsSelector != nil {
		in, out := &in.AssertedControlsSelector, &out.AssertedControlsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ControlPanelArn != nil {
		in, out := &in.ControlPanelArn, &out.ControlPanelArn
		*out = new(string)
		**out = **in
	}
	if in.ControlPanelArnRef != nil {
		in, out := &in.ControlPanelArnRef, &out.ControlPanelArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ControlPanelArnSelector != nil {
		in, out := &in.ControlPanelArnSelector, &out.ControlPanelArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.GatingControls != nil {
		in, out := &in.GatingControls, &out.GatingControls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RuleConfig != nil {
		in, out := &in.RuleConfig, &out.RuleConfig
		*out = new(RuleConfigInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetControls != nil {
		in, out := &in.TargetControls, &out.TargetControls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WaitPeriodMs != nil {
		in, out := &in.WaitPeriodMs, &out.WaitPeriodMs
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SafetyRuleInitParameters.
func (in *SafetyRuleInitParameters) DeepCopy() *SafetyRuleInitParameters {
	if in == nil {
		return nil
	}
	out := new(SafetyRuleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SafetyRuleList) DeepCopyInto(out *SafetyRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SafetyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SafetyRuleList.
func (in *SafetyRuleList) DeepCopy() *SafetyRuleList {
	if in == nil {
		return nil
	}
	out := new(SafetyRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SafetyRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SafetyRuleObservation) DeepCopyInto(out *SafetyRuleObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AssertedControls != nil {
		in, out := &in.AssertedControls, &out.AssertedControls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ControlPanelArn != nil {
		in, out := &in.ControlPanelArn, &out.ControlPanelArn
		*out = new(string)
		**out = **in
	}
	if in.GatingControls != nil {
		in, out := &in.GatingControls, &out.GatingControls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RuleConfig != nil {
		in, out := &in.RuleConfig, &out.RuleConfig
		*out = new(RuleConfigObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TargetControls != nil {
		in, out := &in.TargetControls, &out.TargetControls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WaitPeriodMs != nil {
		in, out := &in.WaitPeriodMs, &out.WaitPeriodMs
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SafetyRuleObservation.
func (in *SafetyRuleObservation) DeepCopy() *SafetyRuleObservation {
	if in == nil {
		return nil
	}
	out := new(SafetyRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SafetyRuleParameters) DeepCopyInto(out *SafetyRuleParameters) {
	*out = *in
	if in.AssertedControls != nil {
		in, out := &in.AssertedControls, &out.AssertedControls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AssertedControlsRefs != nil {
		in, out := &in.AssertedControlsRefs, &out.AssertedControlsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AssertedControlsSelector != nil {
		in, out := &in.AssertedControlsSelector, &out.AssertedControlsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ControlPanelArn != nil {
		in, out := &in.ControlPanelArn, &out.ControlPanelArn
		*out = new(string)
		**out = **in
	}
	if in.ControlPanelArnRef != nil {
		in, out := &in.ControlPanelArnRef, &out.ControlPanelArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ControlPanelArnSelector != nil {
		in, out := &in.ControlPanelArnSelector, &out.ControlPanelArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.GatingControls != nil {
		in, out := &in.GatingControls, &out.GatingControls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RuleConfig != nil {
		in, out := &in.RuleConfig, &out.RuleConfig
		*out = new(RuleConfigParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetControls != nil {
		in, out := &in.TargetControls, &out.TargetControls
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.WaitPeriodMs != nil {
		in, out := &in.WaitPeriodMs, &out.WaitPeriodMs
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SafetyRuleParameters.
func (in *SafetyRuleParameters) DeepCopy() *SafetyRuleParameters {
	if in == nil {
		return nil
	}
	out := new(SafetyRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SafetyRuleSpec) DeepCopyInto(out *SafetyRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SafetyRuleSpec.
func (in *SafetyRuleSpec) DeepCopy() *SafetyRuleSpec {
	if in == nil {
		return nil
	}
	out := new(SafetyRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SafetyRuleStatus) DeepCopyInto(out *SafetyRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SafetyRuleStatus.
func (in *SafetyRuleStatus) DeepCopy() *SafetyRuleStatus {
	if in == nil {
		return nil
	}
	out := new(SafetyRuleStatus)
	in.DeepCopyInto(out)
	return out
}
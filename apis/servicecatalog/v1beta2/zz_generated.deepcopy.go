//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionInitParameters) DeepCopyInto(out *DefinitionInitParameters) {
	*out = *in
	if in.AssumeRole != nil {
		in, out := &in.AssumeRole, &out.AssumeRole
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionInitParameters.
func (in *DefinitionInitParameters) DeepCopy() *DefinitionInitParameters {
	if in == nil {
		return nil
	}
	out := new(DefinitionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionObservation) DeepCopyInto(out *DefinitionObservation) {
	*out = *in
	if in.AssumeRole != nil {
		in, out := &in.AssumeRole, &out.AssumeRole
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionObservation.
func (in *DefinitionObservation) DeepCopy() *DefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(DefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionParameters) DeepCopyInto(out *DefinitionParameters) {
	*out = *in
	if in.AssumeRole != nil {
		in, out := &in.AssumeRole, &out.AssumeRole
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionParameters.
func (in *DefinitionParameters) DeepCopy() *DefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(DefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Product) DeepCopyInto(out *Product) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Product.
func (in *Product) DeepCopy() *Product {
	if in == nil {
		return nil
	}
	out := new(Product)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Product) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductInitParameters) DeepCopyInto(out *ProductInitParameters) {
	*out = *in
	if in.AcceptLanguage != nil {
		in, out := &in.AcceptLanguage, &out.AcceptLanguage
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Distributor != nil {
		in, out := &in.Distributor, &out.Distributor
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningArtifactParameters != nil {
		in, out := &in.ProvisioningArtifactParameters, &out.ProvisioningArtifactParameters
		*out = new(ProvisioningArtifactParametersInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.SupportDescription != nil {
		in, out := &in.SupportDescription, &out.SupportDescription
		*out = new(string)
		**out = **in
	}
	if in.SupportEmail != nil {
		in, out := &in.SupportEmail, &out.SupportEmail
		*out = new(string)
		**out = **in
	}
	if in.SupportURL != nil {
		in, out := &in.SupportURL, &out.SupportURL
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductInitParameters.
func (in *ProductInitParameters) DeepCopy() *ProductInitParameters {
	if in == nil {
		return nil
	}
	out := new(ProductInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductList) DeepCopyInto(out *ProductList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Product, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductList.
func (in *ProductList) DeepCopy() *ProductList {
	if in == nil {
		return nil
	}
	out := new(ProductList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProductList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductObservation) DeepCopyInto(out *ProductObservation) {
	*out = *in
	if in.AcceptLanguage != nil {
		in, out := &in.AcceptLanguage, &out.AcceptLanguage
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Distributor != nil {
		in, out := &in.Distributor, &out.Distributor
		*out = new(string)
		**out = **in
	}
	if in.HasDefaultPath != nil {
		in, out := &in.HasDefaultPath, &out.HasDefaultPath
		*out = new(bool)
		**out = **in
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
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningArtifactParameters != nil {
		in, out := &in.ProvisioningArtifactParameters, &out.ProvisioningArtifactParameters
		*out = new(ProvisioningArtifactParametersObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SupportDescription != nil {
		in, out := &in.SupportDescription, &out.SupportDescription
		*out = new(string)
		**out = **in
	}
	if in.SupportEmail != nil {
		in, out := &in.SupportEmail, &out.SupportEmail
		*out = new(string)
		**out = **in
	}
	if in.SupportURL != nil {
		in, out := &in.SupportURL, &out.SupportURL
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductObservation.
func (in *ProductObservation) DeepCopy() *ProductObservation {
	if in == nil {
		return nil
	}
	out := new(ProductObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductParameters) DeepCopyInto(out *ProductParameters) {
	*out = *in
	if in.AcceptLanguage != nil {
		in, out := &in.AcceptLanguage, &out.AcceptLanguage
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Distributor != nil {
		in, out := &in.Distributor, &out.Distributor
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningArtifactParameters != nil {
		in, out := &in.ProvisioningArtifactParameters, &out.ProvisioningArtifactParameters
		*out = new(ProvisioningArtifactParametersParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SupportDescription != nil {
		in, out := &in.SupportDescription, &out.SupportDescription
		*out = new(string)
		**out = **in
	}
	if in.SupportEmail != nil {
		in, out := &in.SupportEmail, &out.SupportEmail
		*out = new(string)
		**out = **in
	}
	if in.SupportURL != nil {
		in, out := &in.SupportURL, &out.SupportURL
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductParameters.
func (in *ProductParameters) DeepCopy() *ProductParameters {
	if in == nil {
		return nil
	}
	out := new(ProductParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductSpec) DeepCopyInto(out *ProductSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductSpec.
func (in *ProductSpec) DeepCopy() *ProductSpec {
	if in == nil {
		return nil
	}
	out := new(ProductSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductStatus) DeepCopyInto(out *ProductStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductStatus.
func (in *ProductStatus) DeepCopy() *ProductStatus {
	if in == nil {
		return nil
	}
	out := new(ProductStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisioningArtifactParametersInitParameters) DeepCopyInto(out *ProvisioningArtifactParametersInitParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableTemplateValidation != nil {
		in, out := &in.DisableTemplateValidation, &out.DisableTemplateValidation
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TemplatePhysicalID != nil {
		in, out := &in.TemplatePhysicalID, &out.TemplatePhysicalID
		*out = new(string)
		**out = **in
	}
	if in.TemplateURL != nil {
		in, out := &in.TemplateURL, &out.TemplateURL
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisioningArtifactParametersInitParameters.
func (in *ProvisioningArtifactParametersInitParameters) DeepCopy() *ProvisioningArtifactParametersInitParameters {
	if in == nil {
		return nil
	}
	out := new(ProvisioningArtifactParametersInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisioningArtifactParametersObservation) DeepCopyInto(out *ProvisioningArtifactParametersObservation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableTemplateValidation != nil {
		in, out := &in.DisableTemplateValidation, &out.DisableTemplateValidation
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TemplatePhysicalID != nil {
		in, out := &in.TemplatePhysicalID, &out.TemplatePhysicalID
		*out = new(string)
		**out = **in
	}
	if in.TemplateURL != nil {
		in, out := &in.TemplateURL, &out.TemplateURL
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisioningArtifactParametersObservation.
func (in *ProvisioningArtifactParametersObservation) DeepCopy() *ProvisioningArtifactParametersObservation {
	if in == nil {
		return nil
	}
	out := new(ProvisioningArtifactParametersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisioningArtifactParametersParameters) DeepCopyInto(out *ProvisioningArtifactParametersParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableTemplateValidation != nil {
		in, out := &in.DisableTemplateValidation, &out.DisableTemplateValidation
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TemplatePhysicalID != nil {
		in, out := &in.TemplatePhysicalID, &out.TemplatePhysicalID
		*out = new(string)
		**out = **in
	}
	if in.TemplateURL != nil {
		in, out := &in.TemplateURL, &out.TemplateURL
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisioningArtifactParametersParameters.
func (in *ProvisioningArtifactParametersParameters) DeepCopy() *ProvisioningArtifactParametersParameters {
	if in == nil {
		return nil
	}
	out := new(ProvisioningArtifactParametersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAction) DeepCopyInto(out *ServiceAction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAction.
func (in *ServiceAction) DeepCopy() *ServiceAction {
	if in == nil {
		return nil
	}
	out := new(ServiceAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceAction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceActionInitParameters) DeepCopyInto(out *ServiceActionInitParameters) {
	*out = *in
	if in.AcceptLanguage != nil {
		in, out := &in.AcceptLanguage, &out.AcceptLanguage
		*out = new(string)
		**out = **in
	}
	if in.Definition != nil {
		in, out := &in.Definition, &out.Definition
		*out = new(DefinitionInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceActionInitParameters.
func (in *ServiceActionInitParameters) DeepCopy() *ServiceActionInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceActionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceActionList) DeepCopyInto(out *ServiceActionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceActionList.
func (in *ServiceActionList) DeepCopy() *ServiceActionList {
	if in == nil {
		return nil
	}
	out := new(ServiceActionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceActionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceActionObservation) DeepCopyInto(out *ServiceActionObservation) {
	*out = *in
	if in.AcceptLanguage != nil {
		in, out := &in.AcceptLanguage, &out.AcceptLanguage
		*out = new(string)
		**out = **in
	}
	if in.Definition != nil {
		in, out := &in.Definition, &out.Definition
		*out = new(DefinitionObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceActionObservation.
func (in *ServiceActionObservation) DeepCopy() *ServiceActionObservation {
	if in == nil {
		return nil
	}
	out := new(ServiceActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceActionParameters) DeepCopyInto(out *ServiceActionParameters) {
	*out = *in
	if in.AcceptLanguage != nil {
		in, out := &in.AcceptLanguage, &out.AcceptLanguage
		*out = new(string)
		**out = **in
	}
	if in.Definition != nil {
		in, out := &in.Definition, &out.Definition
		*out = new(DefinitionParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceActionParameters.
func (in *ServiceActionParameters) DeepCopy() *ServiceActionParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceActionSpec) DeepCopyInto(out *ServiceActionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceActionSpec.
func (in *ServiceActionSpec) DeepCopy() *ServiceActionSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceActionStatus) DeepCopyInto(out *ServiceActionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceActionStatus.
func (in *ServiceActionStatus) DeepCopy() *ServiceActionStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceActionStatus)
	in.DeepCopyInto(out)
	return out
}
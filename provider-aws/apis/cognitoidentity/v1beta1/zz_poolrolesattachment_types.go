/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MappingRuleObservation struct {
}

type MappingRuleParameters struct {

	// +kubebuilder:validation:Required
	Claim *string `json:"claim" tf:"claim,omitempty"`

	// +kubebuilder:validation:Required
	MatchType *string `json:"matchType" tf:"match_type,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type PoolRolesAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PoolRolesAttachmentParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/cognitoidentity/v1beta1.Pool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IdentityPoolID *string `json:"identityPoolId,omitempty" tf:"identity_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityPoolIDRef *v1.Reference `json:"identityPoolIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	IdentityPoolIDSelector *v1.Selector `json:"identityPoolIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleMapping []RoleMappingParameters `json:"roleMapping,omitempty" tf:"role_mapping,omitempty"`

	// +kubebuilder:validation:Required
	Roles map[string]*string `json:"roles" tf:"roles,omitempty"`
}

type RoleMappingObservation struct {
}

type RoleMappingParameters struct {

	// +kubebuilder:validation:Optional
	AmbiguousRoleResolution *string `json:"ambiguousRoleResolution,omitempty" tf:"ambiguous_role_resolution,omitempty"`

	// +kubebuilder:validation:Required
	IdentityProvider *string `json:"identityProvider" tf:"identity_provider,omitempty"`

	// +kubebuilder:validation:Optional
	MappingRule []MappingRuleParameters `json:"mappingRule,omitempty" tf:"mapping_rule,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// PoolRolesAttachmentSpec defines the desired state of PoolRolesAttachment
type PoolRolesAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PoolRolesAttachmentParameters `json:"forProvider"`
}

// PoolRolesAttachmentStatus defines the observed state of PoolRolesAttachment.
type PoolRolesAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PoolRolesAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PoolRolesAttachment is the Schema for the PoolRolesAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PoolRolesAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PoolRolesAttachmentSpec   `json:"spec"`
	Status            PoolRolesAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PoolRolesAttachmentList contains a list of PoolRolesAttachments
type PoolRolesAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PoolRolesAttachment `json:"items"`
}

// Repository type metadata.
var (
	PoolRolesAttachment_Kind             = "PoolRolesAttachment"
	PoolRolesAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PoolRolesAttachment_Kind}.String()
	PoolRolesAttachment_KindAPIVersion   = PoolRolesAttachment_Kind + "." + CRDGroupVersion.String()
	PoolRolesAttachment_GroupVersionKind = CRDGroupVersion.WithKind(PoolRolesAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&PoolRolesAttachment{}, &PoolRolesAttachmentList{})
}

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

type EntryInitParameters struct {

	// Description of this entry. Due to API limitations, updating only the description of an existing entry requires temporarily removing and re-adding the entry.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type EntryObservation struct {

	// CIDR block of this entry.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Description of this entry. Due to API limitations, updating only the description of an existing entry requires temporarily removing and re-adding the entry.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type EntryParameters struct {

	// CIDR block of this entry.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPCIPv4CidrBlockAssociation
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("cidr_block",false)
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Reference to a VPCIPv4CidrBlockAssociation in ec2 to populate cidr.
	// +kubebuilder:validation:Optional
	CidrRef *v1.Reference `json:"cidrRef,omitempty" tf:"-"`

	// Selector for a VPCIPv4CidrBlockAssociation in ec2 to populate cidr.
	// +kubebuilder:validation:Optional
	CidrSelector *v1.Selector `json:"cidrSelector,omitempty" tf:"-"`

	// Description of this entry. Due to API limitations, updating only the description of an existing entry requires temporarily removing and re-adding the entry.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type ManagedPrefixListInitParameters struct {

	// Address family (IPv4 or IPv6) of this prefix list.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entry []EntryInitParameters `json:"entry,omitempty" tf:"entry,omitempty"`

	// Maximum number of entries that this prefix list can contain.
	MaxEntries *float64 `json:"maxEntries,omitempty" tf:"max_entries,omitempty"`

	// Name of this resource. The name must not start with com.amazonaws.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ManagedPrefixListObservation struct {

	// Address family (IPv4 or IPv6) of this prefix list.
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// ARN of the prefix list.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	Entry []EntryObservation `json:"entry,omitempty" tf:"entry,omitempty"`

	// ID of the prefix list.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Maximum number of entries that this prefix list can contain.
	MaxEntries *float64 `json:"maxEntries,omitempty" tf:"max_entries,omitempty"`

	// Name of this resource. The name must not start with com.amazonaws.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the AWS account that owns this prefix list.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Latest version of this prefix list.
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type ManagedPrefixListParameters struct {

	// Address family (IPv4 or IPv6) of this prefix list.
	// +kubebuilder:validation:Optional
	AddressFamily *string `json:"addressFamily,omitempty" tf:"address_family,omitempty"`

	// Configuration block for prefix list entry. Detailed below. Different entries may have overlapping CIDR blocks, but a particular CIDR should not be duplicated.
	// +kubebuilder:validation:Optional
	Entry []EntryParameters `json:"entry,omitempty" tf:"entry,omitempty"`

	// Maximum number of entries that this prefix list can contain.
	// +kubebuilder:validation:Optional
	MaxEntries *float64 `json:"maxEntries,omitempty" tf:"max_entries,omitempty"`

	// Name of this resource. The name must not start with com.amazonaws.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ManagedPrefixListSpec defines the desired state of ManagedPrefixList
type ManagedPrefixListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedPrefixListParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ManagedPrefixListInitParameters `json:"initProvider,omitempty"`
}

// ManagedPrefixListStatus defines the observed state of ManagedPrefixList.
type ManagedPrefixListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedPrefixListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPrefixList is the Schema for the ManagedPrefixLists API. Provides a managed prefix list resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ManagedPrefixList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.addressFamily) || has(self.initProvider.addressFamily)",message="addressFamily is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maxEntries) || has(self.initProvider.maxEntries)",message="maxEntries is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   ManagedPrefixListSpec   `json:"spec"`
	Status ManagedPrefixListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedPrefixListList contains a list of ManagedPrefixLists
type ManagedPrefixListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedPrefixList `json:"items"`
}

// Repository type metadata.
var (
	ManagedPrefixList_Kind             = "ManagedPrefixList"
	ManagedPrefixList_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedPrefixList_Kind}.String()
	ManagedPrefixList_KindAPIVersion   = ManagedPrefixList_Kind + "." + CRDGroupVersion.String()
	ManagedPrefixList_GroupVersionKind = CRDGroupVersion.WithKind(ManagedPrefixList_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedPrefixList{}, &ManagedPrefixListList{})
}

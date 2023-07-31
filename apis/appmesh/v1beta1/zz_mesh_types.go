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

type EgressFilterInitParameters struct {

	// Egress filter type. By default, the type is DROP_ALL.
	// Valid values are ALLOW_ALL and DROP_ALL.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EgressFilterObservation struct {

	// Egress filter type. By default, the type is DROP_ALL.
	// Valid values are ALLOW_ALL and DROP_ALL.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type EgressFilterParameters struct {

	// Egress filter type. By default, the type is DROP_ALL.
	// Valid values are ALLOW_ALL and DROP_ALL.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MeshInitParameters struct {

	// Service mesh specification to apply.
	Spec []MeshSpecInitParameters `json:"spec,omitempty" tf:"spec,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MeshObservation struct {

	// ARN of the service mesh.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Creation date of the service mesh.
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date,omitempty"`

	// ID of the service mesh.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Last update date of the service mesh.
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date,omitempty"`

	// AWS account ID of the service mesh's owner.
	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner,omitempty"`

	// Resource owner's AWS account ID.
	ResourceOwner *string `json:"resourceOwner,omitempty" tf:"resource_owner,omitempty"`

	// Service mesh specification to apply.
	Spec []MeshSpecObservation `json:"spec,omitempty" tf:"spec,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type MeshParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Service mesh specification to apply.
	// +kubebuilder:validation:Optional
	Spec []MeshSpecParameters `json:"spec,omitempty" tf:"spec,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MeshSpecInitParameters struct {

	// Egress filter rules for the service mesh.
	EgressFilter []EgressFilterInitParameters `json:"egressFilter,omitempty" tf:"egress_filter,omitempty"`
}

type MeshSpecObservation struct {

	// Egress filter rules for the service mesh.
	EgressFilter []EgressFilterObservation `json:"egressFilter,omitempty" tf:"egress_filter,omitempty"`
}

type MeshSpecParameters struct {

	// Egress filter rules for the service mesh.
	// +kubebuilder:validation:Optional
	EgressFilter []EgressFilterParameters `json:"egressFilter,omitempty" tf:"egress_filter,omitempty"`
}

// MeshSpec defines the desired state of Mesh
type MeshSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MeshParameters `json:"forProvider"`
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
	InitProvider MeshInitParameters `json:"initProvider,omitempty"`
}

// MeshStatus defines the observed state of Mesh.
type MeshStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MeshObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Mesh is the Schema for the Meshs API. Provides an AWS App Mesh service mesh resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Mesh struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MeshSpec   `json:"spec"`
	Status            MeshStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MeshList contains a list of Meshs
type MeshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mesh `json:"items"`
}

// Repository type metadata.
var (
	Mesh_Kind             = "Mesh"
	Mesh_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Mesh_Kind}.String()
	Mesh_KindAPIVersion   = Mesh_Kind + "." + CRDGroupVersion.String()
	Mesh_GroupVersionKind = CRDGroupVersion.WithKind(Mesh_Kind)
)

func init() {
	SchemeBuilder.Register(&Mesh{}, &MeshList{})
}

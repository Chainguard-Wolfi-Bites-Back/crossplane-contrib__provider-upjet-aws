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

type ContributorInsightsInitParameters struct {

	// The global secondary index name
	IndexName *string `json:"indexName,omitempty" tf:"index_name,omitempty"`
}

type ContributorInsightsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The global secondary index name
	IndexName *string `json:"indexName,omitempty" tf:"index_name,omitempty"`

	// The name of the table to enable contributor insights
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

type ContributorInsightsParameters struct {

	// The global secondary index name
	// +kubebuilder:validation:Optional
	IndexName *string `json:"indexName,omitempty" tf:"index_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the table to enable contributor insights
	// +crossplane:generate:reference:type=Table
	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Reference to a Table to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameRef *v1.Reference `json:"tableNameRef,omitempty" tf:"-"`

	// Selector for a Table to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameSelector *v1.Selector `json:"tableNameSelector,omitempty" tf:"-"`
}

// ContributorInsightsSpec defines the desired state of ContributorInsights
type ContributorInsightsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContributorInsightsParameters `json:"forProvider"`
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
	InitProvider ContributorInsightsInitParameters `json:"initProvider,omitempty"`
}

// ContributorInsightsStatus defines the observed state of ContributorInsights.
type ContributorInsightsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContributorInsightsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContributorInsights is the Schema for the ContributorInsightss API. Provides a DynamoDB contributor insights resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ContributorInsights struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContributorInsightsSpec   `json:"spec"`
	Status            ContributorInsightsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContributorInsightsList contains a list of ContributorInsightss
type ContributorInsightsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContributorInsights `json:"items"`
}

// Repository type metadata.
var (
	ContributorInsights_Kind             = "ContributorInsights"
	ContributorInsights_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContributorInsights_Kind}.String()
	ContributorInsights_KindAPIVersion   = ContributorInsights_Kind + "." + CRDGroupVersion.String()
	ContributorInsights_GroupVersionKind = CRDGroupVersion.WithKind(ContributorInsights_Kind)
)

func init() {
	SchemeBuilder.Register(&ContributorInsights{}, &ContributorInsightsList{})
}

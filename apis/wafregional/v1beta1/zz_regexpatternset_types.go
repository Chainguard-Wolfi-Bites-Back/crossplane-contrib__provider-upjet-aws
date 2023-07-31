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

type RegexPatternSetInitParameters struct {

	// The name or description of the Regex Pattern Set.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as B[a@]dB[o0]t.
	RegexPatternStrings []*string `json:"regexPatternStrings,omitempty" tf:"regex_pattern_strings,omitempty"`
}

type RegexPatternSetObservation struct {

	// The ID of the WAF Regional Regex Pattern Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or description of the Regex Pattern Set.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as B[a@]dB[o0]t.
	RegexPatternStrings []*string `json:"regexPatternStrings,omitempty" tf:"regex_pattern_strings,omitempty"`
}

type RegexPatternSetParameters struct {

	// The name or description of the Regex Pattern Set.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as B[a@]dB[o0]t.
	// +kubebuilder:validation:Optional
	RegexPatternStrings []*string `json:"regexPatternStrings,omitempty" tf:"regex_pattern_strings,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// RegexPatternSetSpec defines the desired state of RegexPatternSet
type RegexPatternSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegexPatternSetParameters `json:"forProvider"`
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
	InitProvider RegexPatternSetInitParameters `json:"initProvider,omitempty"`
}

// RegexPatternSetStatus defines the observed state of RegexPatternSet.
type RegexPatternSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegexPatternSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegexPatternSet is the Schema for the RegexPatternSets API. Provides a AWS WAF Regional Regex Pattern Set resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RegexPatternSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	Spec   RegexPatternSetSpec   `json:"spec"`
	Status RegexPatternSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegexPatternSetList contains a list of RegexPatternSets
type RegexPatternSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegexPatternSet `json:"items"`
}

// Repository type metadata.
var (
	RegexPatternSet_Kind             = "RegexPatternSet"
	RegexPatternSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegexPatternSet_Kind}.String()
	RegexPatternSet_KindAPIVersion   = RegexPatternSet_Kind + "." + CRDGroupVersion.String()
	RegexPatternSet_GroupVersionKind = CRDGroupVersion.WithKind(RegexPatternSet_Kind)
)

func init() {
	SchemeBuilder.Register(&RegexPatternSet{}, &RegexPatternSetList{})
}

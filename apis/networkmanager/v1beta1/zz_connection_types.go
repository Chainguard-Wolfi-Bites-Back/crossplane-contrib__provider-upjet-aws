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

type ConnectionInitParameters struct {

	// The ID of the link for the second device.
	ConnectedLinkID *string `json:"connectedLinkId,omitempty" tf:"connected_link_id,omitempty"`

	// A description of the connection.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the link for the first device.
	LinkID *string `json:"linkId,omitempty" tf:"link_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConnectionObservation struct {

	// The Amazon Resource Name (ARN) of the connection.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the second device in the connection.
	ConnectedDeviceID *string `json:"connectedDeviceId,omitempty" tf:"connected_device_id,omitempty"`

	// The ID of the link for the second device.
	ConnectedLinkID *string `json:"connectedLinkId,omitempty" tf:"connected_link_id,omitempty"`

	// A description of the connection.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the first device in the connection.
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// The ID of the global network.
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the link for the first device.
	LinkID *string `json:"linkId,omitempty" tf:"link_id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ConnectionParameters struct {

	// The ID of the second device in the connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.Device
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ConnectedDeviceID *string `json:"connectedDeviceId,omitempty" tf:"connected_device_id,omitempty"`

	// Reference to a Device in networkmanager to populate connectedDeviceId.
	// +kubebuilder:validation:Optional
	ConnectedDeviceIDRef *v1.Reference `json:"connectedDeviceIdRef,omitempty" tf:"-"`

	// Selector for a Device in networkmanager to populate connectedDeviceId.
	// +kubebuilder:validation:Optional
	ConnectedDeviceIDSelector *v1.Selector `json:"connectedDeviceIdSelector,omitempty" tf:"-"`

	// The ID of the link for the second device.
	// +kubebuilder:validation:Optional
	ConnectedLinkID *string `json:"connectedLinkId,omitempty" tf:"connected_link_id,omitempty"`

	// A description of the connection.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the first device in the connection.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.Device
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// Reference to a Device in networkmanager to populate deviceId.
	// +kubebuilder:validation:Optional
	DeviceIDRef *v1.Reference `json:"deviceIdRef,omitempty" tf:"-"`

	// Selector for a Device in networkmanager to populate deviceId.
	// +kubebuilder:validation:Optional
	DeviceIDSelector *v1.Selector `json:"deviceIdSelector,omitempty" tf:"-"`

	// The ID of the global network.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.GlobalNetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	// Reference to a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDRef *v1.Reference `json:"globalNetworkIdRef,omitempty" tf:"-"`

	// Selector for a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDSelector *v1.Selector `json:"globalNetworkIdSelector,omitempty" tf:"-"`

	// The ID of the link for the first device.
	// +kubebuilder:validation:Optional
	LinkID *string `json:"linkId,omitempty" tf:"link_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ConnectionSpec defines the desired state of Connection
type ConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionParameters `json:"forProvider"`
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
	InitProvider ConnectionInitParameters `json:"initProvider,omitempty"`
}

// ConnectionStatus defines the observed state of Connection.
type ConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Connection is the Schema for the Connections API. Creates a connection between two devices.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Connection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionSpec   `json:"spec"`
	Status            ConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionList contains a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Connection `json:"items"`
}

// Repository type metadata.
var (
	Connection_Kind             = "Connection"
	Connection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Connection_Kind}.String()
	Connection_KindAPIVersion   = Connection_Kind + "." + CRDGroupVersion.String()
	Connection_GroupVersionKind = CRDGroupVersion.WithKind(Connection_Kind)
)

func init() {
	SchemeBuilder.Register(&Connection{}, &ConnectionList{})
}

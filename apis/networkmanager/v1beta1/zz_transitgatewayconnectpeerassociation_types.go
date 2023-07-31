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

type TransitGatewayConnectPeerAssociationInitParameters struct {

	// The ID of the link.
	LinkID *string `json:"linkId,omitempty" tf:"link_id,omitempty"`
}

type TransitGatewayConnectPeerAssociationObservation struct {

	// The ID of the device.
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// The ID of the global network.
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the link.
	LinkID *string `json:"linkId,omitempty" tf:"link_id,omitempty"`

	// The Amazon Resource Name (ARN) of the Connect peer.
	TransitGatewayConnectPeerArn *string `json:"transitGatewayConnectPeerArn,omitempty" tf:"transit_gateway_connect_peer_arn,omitempty"`
}

type TransitGatewayConnectPeerAssociationParameters struct {

	// The ID of the device.
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

	// The ID of the link.
	// +kubebuilder:validation:Optional
	LinkID *string `json:"linkId,omitempty" tf:"link_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) of the Connect peer.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.TransitGatewayConnectPeer
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	TransitGatewayConnectPeerArn *string `json:"transitGatewayConnectPeerArn,omitempty" tf:"transit_gateway_connect_peer_arn,omitempty"`

	// Reference to a TransitGatewayConnectPeer in ec2 to populate transitGatewayConnectPeerArn.
	// +kubebuilder:validation:Optional
	TransitGatewayConnectPeerArnRef *v1.Reference `json:"transitGatewayConnectPeerArnRef,omitempty" tf:"-"`

	// Selector for a TransitGatewayConnectPeer in ec2 to populate transitGatewayConnectPeerArn.
	// +kubebuilder:validation:Optional
	TransitGatewayConnectPeerArnSelector *v1.Selector `json:"transitGatewayConnectPeerArnSelector,omitempty" tf:"-"`
}

// TransitGatewayConnectPeerAssociationSpec defines the desired state of TransitGatewayConnectPeerAssociation
type TransitGatewayConnectPeerAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayConnectPeerAssociationParameters `json:"forProvider"`
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
	InitProvider TransitGatewayConnectPeerAssociationInitParameters `json:"initProvider,omitempty"`
}

// TransitGatewayConnectPeerAssociationStatus defines the observed state of TransitGatewayConnectPeerAssociation.
type TransitGatewayConnectPeerAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayConnectPeerAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayConnectPeerAssociation is the Schema for the TransitGatewayConnectPeerAssociations API. Associates a transit gateway Connect peer with a device, and optionally, with a link.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayConnectPeerAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayConnectPeerAssociationSpec   `json:"spec"`
	Status            TransitGatewayConnectPeerAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayConnectPeerAssociationList contains a list of TransitGatewayConnectPeerAssociations
type TransitGatewayConnectPeerAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayConnectPeerAssociation `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayConnectPeerAssociation_Kind             = "TransitGatewayConnectPeerAssociation"
	TransitGatewayConnectPeerAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayConnectPeerAssociation_Kind}.String()
	TransitGatewayConnectPeerAssociation_KindAPIVersion   = TransitGatewayConnectPeerAssociation_Kind + "." + CRDGroupVersion.String()
	TransitGatewayConnectPeerAssociation_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayConnectPeerAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayConnectPeerAssociation{}, &TransitGatewayConnectPeerAssociationList{})
}

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

type LoggingConfigurationInitParameters struct {

	// Determines whether execution data is included in your log. When set to false, data is excluded.
	IncludeExecutionData *bool `json:"includeExecutionData,omitempty" tf:"include_execution_data,omitempty"`

	// Defines which category of execution history events are logged. Valid values: ALL, ERROR, FATAL, OFF
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// Amazon Resource Name (ARN) of a CloudWatch log group. Make sure the State Machine has the correct IAM policies for logging. The ARN must end with :*
	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`
}

type LoggingConfigurationObservation struct {

	// Determines whether execution data is included in your log. When set to false, data is excluded.
	IncludeExecutionData *bool `json:"includeExecutionData,omitempty" tf:"include_execution_data,omitempty"`

	// Defines which category of execution history events are logged. Valid values: ALL, ERROR, FATAL, OFF
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// Amazon Resource Name (ARN) of a CloudWatch log group. Make sure the State Machine has the correct IAM policies for logging. The ARN must end with :*
	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`
}

type LoggingConfigurationParameters struct {

	// Determines whether execution data is included in your log. When set to false, data is excluded.
	// +kubebuilder:validation:Optional
	IncludeExecutionData *bool `json:"includeExecutionData,omitempty" tf:"include_execution_data,omitempty"`

	// Defines which category of execution history events are logged. Valid values: ALL, ERROR, FATAL, OFF
	// +kubebuilder:validation:Optional
	Level *string `json:"level,omitempty" tf:"level,omitempty"`

	// Amazon Resource Name (ARN) of a CloudWatch log group. Make sure the State Machine has the correct IAM policies for logging. The ARN must end with :*
	// +kubebuilder:validation:Optional
	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`
}

type StateMachineInitParameters struct {

	// The Amazon States Language definition of the state machine.
	Definition *string `json:"definition,omitempty" tf:"definition,omitempty"`

	// Defines what execution history events are logged and where they are logged. The logging_configuration parameter is only valid when type is set to EXPRESS. Defaults to OFF. For more information see Logging Express Workflows and Log Levels in the AWS Step Functions User Guide.
	LoggingConfiguration []LoggingConfigurationInitParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Selects whether AWS X-Ray tracing is enabled.
	TracingConfiguration []TracingConfigurationInitParameters `json:"tracingConfiguration,omitempty" tf:"tracing_configuration,omitempty"`

	// Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid values: STANDARD, EXPRESS.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StateMachineObservation struct {

	// The ARN of the state machine.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The date the state machine was created.
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// The Amazon States Language definition of the state machine.
	Definition *string `json:"definition,omitempty" tf:"definition,omitempty"`

	// The ARN of the state machine.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Defines what execution history events are logged and where they are logged. The logging_configuration parameter is only valid when type is set to EXPRESS. Defaults to OFF. For more information see Logging Express Workflows and Log Levels in the AWS Step Functions User Guide.
	LoggingConfiguration []LoggingConfigurationObservation `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// The current status of the state machine. Either ACTIVE or DELETING.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Selects whether AWS X-Ray tracing is enabled.
	TracingConfiguration []TracingConfigurationObservation `json:"tracingConfiguration,omitempty" tf:"tracing_configuration,omitempty"`

	// Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid values: STANDARD, EXPRESS.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StateMachineParameters struct {

	// The Amazon States Language definition of the state machine.
	// +kubebuilder:validation:Optional
	Definition *string `json:"definition,omitempty" tf:"definition,omitempty"`

	// Defines what execution history events are logged and where they are logged. The logging_configuration parameter is only valid when type is set to EXPRESS. Defaults to OFF. For more information see Logging Express Workflows and Log Levels in the AWS Step Functions User Guide.
	// +kubebuilder:validation:Optional
	LoggingConfiguration []LoggingConfigurationParameters `json:"loggingConfiguration,omitempty" tf:"logging_configuration,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Selects whether AWS X-Ray tracing is enabled.
	// +kubebuilder:validation:Optional
	TracingConfiguration []TracingConfigurationParameters `json:"tracingConfiguration,omitempty" tf:"tracing_configuration,omitempty"`

	// Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid values: STANDARD, EXPRESS.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TracingConfigurationInitParameters struct {

	// When set to true, AWS X-Ray tracing is enabled. Make sure the State Machine has the correct IAM policies for logging. See the AWS Step Functions Developer Guide for details.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type TracingConfigurationObservation struct {

	// When set to true, AWS X-Ray tracing is enabled. Make sure the State Machine has the correct IAM policies for logging. See the AWS Step Functions Developer Guide for details.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type TracingConfigurationParameters struct {

	// When set to true, AWS X-Ray tracing is enabled. Make sure the State Machine has the correct IAM policies for logging. See the AWS Step Functions Developer Guide for details.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

// StateMachineSpec defines the desired state of StateMachine
type StateMachineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StateMachineParameters `json:"forProvider"`
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
	InitProvider StateMachineInitParameters `json:"initProvider,omitempty"`
}

// StateMachineStatus defines the observed state of StateMachine.
type StateMachineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StateMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StateMachine is the Schema for the StateMachines API. Provides a Step Function State Machine resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StateMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.definition) || has(self.initProvider.definition)",message="definition is a required parameter"
	Spec   StateMachineSpec   `json:"spec"`
	Status StateMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StateMachineList contains a list of StateMachines
type StateMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StateMachine `json:"items"`
}

// Repository type metadata.
var (
	StateMachine_Kind             = "StateMachine"
	StateMachine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StateMachine_Kind}.String()
	StateMachine_KindAPIVersion   = StateMachine_Kind + "." + CRDGroupVersion.String()
	StateMachine_GroupVersionKind = CRDGroupVersion.WithKind(StateMachine_Kind)
)

func init() {
	SchemeBuilder.Register(&StateMachine{}, &StateMachineList{})
}

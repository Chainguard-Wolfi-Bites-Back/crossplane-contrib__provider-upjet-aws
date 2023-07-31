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

type ActionDefinitionInitParameters struct {

	// A configuration block describing the stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the matching packet. You can pair this custom action with any of the standard stateless rule actions. See Publish Metric Action below for details.
	PublishMetricAction []PublishMetricActionInitParameters `json:"publishMetricAction,omitempty" tf:"publish_metric_action,omitempty"`
}

type ActionDefinitionObservation struct {

	// A configuration block describing the stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the matching packet. You can pair this custom action with any of the standard stateless rule actions. See Publish Metric Action below for details.
	PublishMetricAction []PublishMetricActionObservation `json:"publishMetricAction,omitempty" tf:"publish_metric_action,omitempty"`
}

type ActionDefinitionParameters struct {

	// A configuration block describing the stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the matching packet. You can pair this custom action with any of the standard stateless rule actions. See Publish Metric Action below for details.
	// +kubebuilder:validation:Optional
	PublishMetricAction []PublishMetricActionParameters `json:"publishMetricAction,omitempty" tf:"publish_metric_action,omitempty"`
}

type DimensionInitParameters struct {

	// The string value to use in the custom metric dimension.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DimensionObservation struct {

	// The string value to use in the custom metric dimension.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type DimensionParameters struct {

	// The string value to use in the custom metric dimension.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type FirewallPolicyEncryptionConfigurationInitParameters struct {

	// The ID of the customer managed key. You can use any of the key identifiers that KMS supports, unless you're using a key that's managed by another account. If you're using a key managed by another account, then specify the key ARN.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The type of AWS KMS key to use for encryption of your Network Firewall resources. Valid values are CUSTOMER_KMS and AWS_OWNED_KMS_KEY.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FirewallPolicyEncryptionConfigurationObservation struct {

	// The ID of the customer managed key. You can use any of the key identifiers that KMS supports, unless you're using a key that's managed by another account. If you're using a key managed by another account, then specify the key ARN.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The type of AWS KMS key to use for encryption of your Network Firewall resources. Valid values are CUSTOMER_KMS and AWS_OWNED_KMS_KEY.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FirewallPolicyEncryptionConfigurationParameters struct {

	// The ID of the customer managed key. You can use any of the key identifiers that KMS supports, unless you're using a key that's managed by another account. If you're using a key managed by another account, then specify the key ARN.
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// The type of AWS KMS key to use for encryption of your Network Firewall resources. Valid values are CUSTOMER_KMS and AWS_OWNED_KMS_KEY.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type FirewallPolicyFirewallPolicyInitParameters struct {

	// Set of actions to take on a packet if it does not match any stateful rules in the policy. This can only be specified if the policy has a stateful_engine_options block with a rule_order value of STRICT_ORDER. You can specify one of either or neither values of aws:drop_strict or aws:drop_established, as well as any combination of aws:alert_strict and aws:alert_established.
	StatefulDefaultActions []*string `json:"statefulDefaultActions,omitempty" tf:"stateful_default_actions,omitempty"`

	// A configuration block that defines options on how the policy handles stateful rules. See Stateful Engine Options below for details.
	StatefulEngineOptions []StatefulEngineOptionsInitParameters `json:"statefulEngineOptions,omitempty" tf:"stateful_engine_options,omitempty"`

	// Set of configuration blocks containing references to the stateful rule groups that are used in the policy. See Stateful Rule Group Reference below for details.
	StatefulRuleGroupReference []StatefulRuleGroupReferenceInitParameters `json:"statefulRuleGroupReference,omitempty" tf:"stateful_rule_group_reference,omitempty"`

	// Set of configuration blocks describing the custom action definitions that are available for use in the firewall policy's stateless_default_actions. See Stateless Custom Action below for details.
	StatelessCustomAction []StatelessCustomActionInitParameters `json:"statelessCustomAction,omitempty" tf:"stateless_custom_action,omitempty"`

	// Set of actions to take on a packet if it does not match any of the stateless rules in the policy. You must specify one of the standard actions including: aws:drop, aws:pass, or aws:forward_to_sfe.
	// In addition, you can specify custom actions that are compatible with your standard action choice. If you want non-matching packets to be forwarded for stateful inspection, specify aws:forward_to_sfe.
	StatelessDefaultActions []*string `json:"statelessDefaultActions,omitempty" tf:"stateless_default_actions,omitempty"`

	// Set of actions to take on a fragmented packet if it does not match any of the stateless rules in the policy. You must specify one of the standard actions including: aws:drop, aws:pass, or aws:forward_to_sfe.
	// In addition, you can specify custom actions that are compatible with your standard action choice. If you want non-matching packets to be forwarded for stateful inspection, specify aws:forward_to_sfe.
	StatelessFragmentDefaultActions []*string `json:"statelessFragmentDefaultActions,omitempty" tf:"stateless_fragment_default_actions,omitempty"`

	// Set of configuration blocks containing references to the stateless rule groups that are used in the policy. See Stateless Rule Group Reference below for details.
	StatelessRuleGroupReference []StatelessRuleGroupReferenceInitParameters `json:"statelessRuleGroupReference,omitempty" tf:"stateless_rule_group_reference,omitempty"`
}

type FirewallPolicyFirewallPolicyObservation struct {

	// Set of actions to take on a packet if it does not match any stateful rules in the policy. This can only be specified if the policy has a stateful_engine_options block with a rule_order value of STRICT_ORDER. You can specify one of either or neither values of aws:drop_strict or aws:drop_established, as well as any combination of aws:alert_strict and aws:alert_established.
	StatefulDefaultActions []*string `json:"statefulDefaultActions,omitempty" tf:"stateful_default_actions,omitempty"`

	// A configuration block that defines options on how the policy handles stateful rules. See Stateful Engine Options below for details.
	StatefulEngineOptions []StatefulEngineOptionsObservation `json:"statefulEngineOptions,omitempty" tf:"stateful_engine_options,omitempty"`

	// Set of configuration blocks containing references to the stateful rule groups that are used in the policy. See Stateful Rule Group Reference below for details.
	StatefulRuleGroupReference []StatefulRuleGroupReferenceObservation `json:"statefulRuleGroupReference,omitempty" tf:"stateful_rule_group_reference,omitempty"`

	// Set of configuration blocks describing the custom action definitions that are available for use in the firewall policy's stateless_default_actions. See Stateless Custom Action below for details.
	StatelessCustomAction []StatelessCustomActionObservation `json:"statelessCustomAction,omitempty" tf:"stateless_custom_action,omitempty"`

	// Set of actions to take on a packet if it does not match any of the stateless rules in the policy. You must specify one of the standard actions including: aws:drop, aws:pass, or aws:forward_to_sfe.
	// In addition, you can specify custom actions that are compatible with your standard action choice. If you want non-matching packets to be forwarded for stateful inspection, specify aws:forward_to_sfe.
	StatelessDefaultActions []*string `json:"statelessDefaultActions,omitempty" tf:"stateless_default_actions,omitempty"`

	// Set of actions to take on a fragmented packet if it does not match any of the stateless rules in the policy. You must specify one of the standard actions including: aws:drop, aws:pass, or aws:forward_to_sfe.
	// In addition, you can specify custom actions that are compatible with your standard action choice. If you want non-matching packets to be forwarded for stateful inspection, specify aws:forward_to_sfe.
	StatelessFragmentDefaultActions []*string `json:"statelessFragmentDefaultActions,omitempty" tf:"stateless_fragment_default_actions,omitempty"`

	// Set of configuration blocks containing references to the stateless rule groups that are used in the policy. See Stateless Rule Group Reference below for details.
	StatelessRuleGroupReference []StatelessRuleGroupReferenceObservation `json:"statelessRuleGroupReference,omitempty" tf:"stateless_rule_group_reference,omitempty"`
}

type FirewallPolicyFirewallPolicyParameters struct {

	// Set of actions to take on a packet if it does not match any stateful rules in the policy. This can only be specified if the policy has a stateful_engine_options block with a rule_order value of STRICT_ORDER. You can specify one of either or neither values of aws:drop_strict or aws:drop_established, as well as any combination of aws:alert_strict and aws:alert_established.
	// +kubebuilder:validation:Optional
	StatefulDefaultActions []*string `json:"statefulDefaultActions,omitempty" tf:"stateful_default_actions,omitempty"`

	// A configuration block that defines options on how the policy handles stateful rules. See Stateful Engine Options below for details.
	// +kubebuilder:validation:Optional
	StatefulEngineOptions []StatefulEngineOptionsParameters `json:"statefulEngineOptions,omitempty" tf:"stateful_engine_options,omitempty"`

	// Set of configuration blocks containing references to the stateful rule groups that are used in the policy. See Stateful Rule Group Reference below for details.
	// +kubebuilder:validation:Optional
	StatefulRuleGroupReference []StatefulRuleGroupReferenceParameters `json:"statefulRuleGroupReference,omitempty" tf:"stateful_rule_group_reference,omitempty"`

	// Set of configuration blocks describing the custom action definitions that are available for use in the firewall policy's stateless_default_actions. See Stateless Custom Action below for details.
	// +kubebuilder:validation:Optional
	StatelessCustomAction []StatelessCustomActionParameters `json:"statelessCustomAction,omitempty" tf:"stateless_custom_action,omitempty"`

	// Set of actions to take on a packet if it does not match any of the stateless rules in the policy. You must specify one of the standard actions including: aws:drop, aws:pass, or aws:forward_to_sfe.
	// In addition, you can specify custom actions that are compatible with your standard action choice. If you want non-matching packets to be forwarded for stateful inspection, specify aws:forward_to_sfe.
	// +kubebuilder:validation:Optional
	StatelessDefaultActions []*string `json:"statelessDefaultActions,omitempty" tf:"stateless_default_actions,omitempty"`

	// Set of actions to take on a fragmented packet if it does not match any of the stateless rules in the policy. You must specify one of the standard actions including: aws:drop, aws:pass, or aws:forward_to_sfe.
	// In addition, you can specify custom actions that are compatible with your standard action choice. If you want non-matching packets to be forwarded for stateful inspection, specify aws:forward_to_sfe.
	// +kubebuilder:validation:Optional
	StatelessFragmentDefaultActions []*string `json:"statelessFragmentDefaultActions,omitempty" tf:"stateless_fragment_default_actions,omitempty"`

	// Set of configuration blocks containing references to the stateless rule groups that are used in the policy. See Stateless Rule Group Reference below for details.
	// +kubebuilder:validation:Optional
	StatelessRuleGroupReference []StatelessRuleGroupReferenceParameters `json:"statelessRuleGroupReference,omitempty" tf:"stateless_rule_group_reference,omitempty"`
}

type FirewallPolicyInitParameters struct {

	// A friendly description of the firewall policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// KMS encryption configuration settings. See Encryption Configuration below for details.
	EncryptionConfiguration []FirewallPolicyEncryptionConfigurationInitParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
	FirewallPolicy []FirewallPolicyFirewallPolicyInitParameters `json:"firewallPolicy,omitempty" tf:"firewall_policy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FirewallPolicyObservation struct {

	// The Amazon Resource Name (ARN) that identifies the firewall policy.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A friendly description of the firewall policy.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// KMS encryption configuration settings. See Encryption Configuration below for details.
	EncryptionConfiguration []FirewallPolicyEncryptionConfigurationObservation `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
	FirewallPolicy []FirewallPolicyFirewallPolicyObservation `json:"firewallPolicy,omitempty" tf:"firewall_policy,omitempty"`

	// The Amazon Resource Name (ARN) that identifies the firewall policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// A string token used when updating a firewall policy.
	UpdateToken *string `json:"updateToken,omitempty" tf:"update_token,omitempty"`
}

type FirewallPolicyParameters struct {

	// A friendly description of the firewall policy.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// KMS encryption configuration settings. See Encryption Configuration below for details.
	// +kubebuilder:validation:Optional
	EncryptionConfiguration []FirewallPolicyEncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// A configuration block describing the rule groups and policy actions to use in the firewall policy. See Firewall Policy below for details.
	// +kubebuilder:validation:Optional
	FirewallPolicy []FirewallPolicyFirewallPolicyParameters `json:"firewallPolicy,omitempty" tf:"firewall_policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type OverrideInitParameters struct {

	// The action that changes the rule group from DROP to ALERT . This only applies to managed rule groups.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`
}

type OverrideObservation struct {

	// The action that changes the rule group from DROP to ALERT . This only applies to managed rule groups.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`
}

type OverrideParameters struct {

	// The action that changes the rule group from DROP to ALERT . This only applies to managed rule groups.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`
}

type PublishMetricActionInitParameters struct {

	// Set of configuration blocks describing dimension settings to use for Amazon CloudWatch custom metrics. See Dimension below for more details.
	Dimension []DimensionInitParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`
}

type PublishMetricActionObservation struct {

	// Set of configuration blocks describing dimension settings to use for Amazon CloudWatch custom metrics. See Dimension below for more details.
	Dimension []DimensionObservation `json:"dimension,omitempty" tf:"dimension,omitempty"`
}

type PublishMetricActionParameters struct {

	// Set of configuration blocks describing dimension settings to use for Amazon CloudWatch custom metrics. See Dimension below for more details.
	// +kubebuilder:validation:Optional
	Dimension []DimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`
}

type StatefulEngineOptionsInitParameters struct {

	// Indicates how to manage the order of stateful rule evaluation for the policy. Default value: DEFAULT_ACTION_ORDER. Valid values: DEFAULT_ACTION_ORDER, STRICT_ORDER.
	RuleOrder *string `json:"ruleOrder,omitempty" tf:"rule_order,omitempty"`
}

type StatefulEngineOptionsObservation struct {

	// Indicates how to manage the order of stateful rule evaluation for the policy. Default value: DEFAULT_ACTION_ORDER. Valid values: DEFAULT_ACTION_ORDER, STRICT_ORDER.
	RuleOrder *string `json:"ruleOrder,omitempty" tf:"rule_order,omitempty"`
}

type StatefulEngineOptionsParameters struct {

	// Indicates how to manage the order of stateful rule evaluation for the policy. Default value: DEFAULT_ACTION_ORDER. Valid values: DEFAULT_ACTION_ORDER, STRICT_ORDER.
	// +kubebuilder:validation:Optional
	RuleOrder *string `json:"ruleOrder,omitempty" tf:"rule_order,omitempty"`
}

type StatefulRuleGroupReferenceInitParameters struct {

	// Configuration block for override values
	Override []OverrideInitParameters `json:"override,omitempty" tf:"override,omitempty"`

	// An integer setting that indicates the order in which to run the stateless rule groups in a single policy. AWS Network Firewall applies each stateless rule group to a packet starting with the group that has the lowest priority setting.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type StatefulRuleGroupReferenceObservation struct {

	// Configuration block for override values
	Override []OverrideObservation `json:"override,omitempty" tf:"override,omitempty"`

	// An integer setting that indicates the order in which to run the stateless rule groups in a single policy. AWS Network Firewall applies each stateless rule group to a packet starting with the group that has the lowest priority setting.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The Amazon Resource Name (ARN) of the stateless rule group.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`
}

type StatefulRuleGroupReferenceParameters struct {

	// Configuration block for override values
	// +kubebuilder:validation:Optional
	Override []OverrideParameters `json:"override,omitempty" tf:"override,omitempty"`

	// An integer setting that indicates the order in which to run the stateless rule groups in a single policy. AWS Network Firewall applies each stateless rule group to a packet starting with the group that has the lowest priority setting.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The Amazon Resource Name (ARN) of the stateless rule group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkfirewall/v1beta1.RuleGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a RuleGroup in networkfirewall to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a RuleGroup in networkfirewall to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`
}

type StatelessCustomActionInitParameters struct {

	// A configuration block describing the custom action associated with the action_name. See Action Definition below for details.
	ActionDefinition []ActionDefinitionInitParameters `json:"actionDefinition,omitempty" tf:"action_definition,omitempty"`

	// A friendly name of the custom action.
	ActionName *string `json:"actionName,omitempty" tf:"action_name,omitempty"`
}

type StatelessCustomActionObservation struct {

	// A configuration block describing the custom action associated with the action_name. See Action Definition below for details.
	ActionDefinition []ActionDefinitionObservation `json:"actionDefinition,omitempty" tf:"action_definition,omitempty"`

	// A friendly name of the custom action.
	ActionName *string `json:"actionName,omitempty" tf:"action_name,omitempty"`
}

type StatelessCustomActionParameters struct {

	// A configuration block describing the custom action associated with the action_name. See Action Definition below for details.
	// +kubebuilder:validation:Optional
	ActionDefinition []ActionDefinitionParameters `json:"actionDefinition,omitempty" tf:"action_definition,omitempty"`

	// A friendly name of the custom action.
	// +kubebuilder:validation:Optional
	ActionName *string `json:"actionName,omitempty" tf:"action_name,omitempty"`
}

type StatelessRuleGroupReferenceInitParameters struct {

	// An integer setting that indicates the order in which to run the stateless rule groups in a single policy. AWS Network Firewall applies each stateless rule group to a packet starting with the group that has the lowest priority setting.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type StatelessRuleGroupReferenceObservation struct {

	// An integer setting that indicates the order in which to run the stateless rule groups in a single policy. AWS Network Firewall applies each stateless rule group to a packet starting with the group that has the lowest priority setting.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The Amazon Resource Name (ARN) of the stateless rule group.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`
}

type StatelessRuleGroupReferenceParameters struct {

	// An integer setting that indicates the order in which to run the stateless rule groups in a single policy. AWS Network Firewall applies each stateless rule group to a packet starting with the group that has the lowest priority setting.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The Amazon Resource Name (ARN) of the stateless rule group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkfirewall/v1beta1.RuleGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a RuleGroup in networkfirewall to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a RuleGroup in networkfirewall to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`
}

// FirewallPolicySpec defines the desired state of FirewallPolicy
type FirewallPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallPolicyParameters `json:"forProvider"`
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
	InitProvider FirewallPolicyInitParameters `json:"initProvider,omitempty"`
}

// FirewallPolicyStatus defines the observed state of FirewallPolicy.
type FirewallPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicy is the Schema for the FirewallPolicys API. Provides an AWS Network Firewall Policy resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.firewallPolicy) || has(self.initProvider.firewallPolicy)",message="firewallPolicy is a required parameter"
	Spec   FirewallPolicySpec   `json:"spec"`
	Status FirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicyList contains a list of FirewallPolicys
type FirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	FirewallPolicy_Kind             = "FirewallPolicy"
	FirewallPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallPolicy_Kind}.String()
	FirewallPolicy_KindAPIVersion   = FirewallPolicy_Kind + "." + CRDGroupVersion.String()
	FirewallPolicy_GroupVersionKind = CRDGroupVersion.WithKind(FirewallPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallPolicy{}, &FirewallPolicyList{})
}

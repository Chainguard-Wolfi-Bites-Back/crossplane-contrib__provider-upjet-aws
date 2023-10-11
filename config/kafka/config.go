package kafka

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the kafka group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_msk_cluster", func(r *config.Resource) {
		r.References["encryption_info.encryption_at_rest_kms_key_arn"] = config.Reference{
			Type:      "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
			Extractor: common.PathARNExtractor,
		}
		r.References["logging_info.broker_logs.s3.bucket"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
		}
		r.References["logging_info.broker_logs.cloudwatch_logs.log_group"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/cloudwatchlogs/v1beta1.Group",
		}
		r.References["broker_node_group_info.client_subnets"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
		}
		r.References["broker_node_group_info.security_groups"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
		}
		r.References["configuration_info.arn"] = config.Reference{
			Type:      "Configuration",
			Extractor: common.PathARNExtractor,
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_msk_scram_secret_association", func(r *config.Resource) {
		r.References["secret_arn_list"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/secretsmanager/v1beta1.Secret",
			RefFieldName:      "SecretArnRefs",
			SelectorFieldName: "SecretArnSelector",
		}
		r.MetaResource.ArgumentDocs["secret_arn_list"] = "- (Required) List of all AWS Secrets Manager secret ARNs to associate with the cluster. Secrets not referenced, selected or listed here will be disassociated from the cluster."
	})
	p.AddResourceConfigurator("aws_msk_serverless_cluster", func(r *config.Resource) {
		r.UseAsync = true
		r.References["vpc_config.security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}
		r.References["vpc_config.subnet_ids"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.Subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}
	})
}

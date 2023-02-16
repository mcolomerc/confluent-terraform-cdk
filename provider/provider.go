package provider

import (
	confluentprovider "cdk.tf/go/stack/generated/confluentinc/confluent/provider"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConfluentProvider struct {
}

func NewConfluentProvider(stack cdktf.TerraformStack) *ConfluentProvider {
	config := providerConfig(stack)
	confluentprovider.NewConfluentProvider(stack, jsii.String("confluent"), &confluentprovider.ConfluentProviderConfig{
		CloudApiKey:    config.CloudApiKey.StringValue(),
		CloudApiSecret: config.CloudApiSecret.StringValue(),
	})
	return &ConfluentProvider{}
}

type ProviderConfig struct {
	CloudApiKey    cdktf.TerraformVariable
	CloudApiSecret cdktf.TerraformVariable
}

func providerConfig(stack cdktf.TerraformStack) *ProviderConfig {
	return &ProviderConfig{
		CloudApiKey: cdktf.NewTerraformVariable(stack, jsii.String("confluent_cloud_api_key"), &cdktf.TerraformVariableConfig{
			Type:        jsii.String("string"),
			Description: jsii.String("Confluent Cloud API_KEY"),
		}),
		CloudApiSecret: cdktf.NewTerraformVariable(stack, jsii.String("confluent_cloud_api_secret"), &cdktf.TerraformVariableConfig{
			Type:        jsii.String("string"),
			Description: jsii.String("Confluent Cloud API_SECRET"),
		}),
	}
}

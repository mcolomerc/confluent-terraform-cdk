package main

import (
	"cdk.tf/go/stack/generated/confluentinc/confluent/kafkacluster"
	confluentprovider "cdk.tf/go/stack/generated/confluentinc/confluent/provider"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	cloudApiKey := ""    //CLOUD_API_KEY
	cloudApiSecret := "" //CLOUD_API_SECRET
	environmentId := ""  //ENVIRONMENT_ID
	var cku float64 = 1  //CKUs

	confluentprovider.NewConfluentProvider(stack, jsii.String("confluent"), &confluentprovider.ConfluentProviderConfig{
		CloudApiKey:    &cloudApiKey,
		CloudApiSecret: &cloudApiSecret,
	})

	kafkacluster.NewKafkaCluster(stack, jsii.String("kafka_cluster"), &kafkacluster.KafkaClusterConfig{
		Availability: jsii.String("SINGLE_ZONE"),
		Cloud:        jsii.String("GCP"),
		DisplayName:  jsii.String("cdktf-cluster"),
		Environment: &kafkacluster.KafkaClusterEnvironment{
			Id: &environmentId,
		},
		Region: jsii.String("europe-west3"),
		Dedicated: &kafkacluster.KafkaClusterDedicated{
			Cku: &cku,
		},
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "confluent-dedicated-cluster")

	app.Synth()
}

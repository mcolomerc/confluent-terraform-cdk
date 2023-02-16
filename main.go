package main

import (
	"cdk.tf/go/stack/cluster"
	"cdk.tf/go/stack/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(nil)

	name := "confluent-cloud-kafka-cluster"

	//Stacks represent a collection of infrastructure that CDK for Terraform (CDKTF) synthesizes as a dedicated Terraform configuration.
	stack := cdktf.NewTerraformStack(app, &name)

	// Provider
	provider.NewConfluentProvider(stack)

	// Cluster
	cluster.NewKafkaCluster(stack, &name)

	app.Synth()

}

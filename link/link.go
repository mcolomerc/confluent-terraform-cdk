package link

import (
	"cdk.tf/go/stack/cluster"
	"cdk.tf/go/stack/config"
	"cdk.tf/go/stack/generated/confluentinc/confluent/clusterlink"
	"cdk.tf/go/stack/generated/confluentinc/confluent/dataconfluentkafkacluster"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterLink struct {
}

var source cluster.KafkaCluster

func CreateClusterLink(sourceCluster cluster.KafkaCluster, scope constructs.Construct, name *string) *ClusterLink {
	source = sourceCluster
	return NewClusterLink(scope, name)
}

func NewClusterLink(scope constructs.Construct, name *string) *ClusterLink {
	c := constructs.NewConstruct(scope, name)

	conf := config.GetConfing()

	sourceCluster := dataconfluentkafkacluster.NewDataConfluentKafkaCluster(c, jsii.String("source-cluster"), &dataconfluentkafkacluster.DataConfluentKafkaClusterConfig{
		Id: jsii.String(conf.Link.Source),
		Environment: &dataconfluentkafkacluster.DataConfluentKafkaClusterEnvironment{
			Id: jsii.String(conf.Environment),
		},
	})

	link := clusterlink.NewClusterLink(c, jsii.String("cluster-link"), &clusterlink.ClusterLinkConfig{
		LinkName: jsii.String(conf.Link.Source + "-link-" + *source.Kafka.Id()),
		SourceKafkaCluster: &clusterlink.ClusterLinkSourceKafkaCluster{
			Id:                jsii.String(conf.Link.Source),
			BootstrapEndpoint: jsii.String(*sourceCluster.BootstrapEndpoint()),
			Credentials: &clusterlink.ClusterLinkSourceKafkaClusterCredentials{
				Key:    jsii.String(conf.Link.Key),
				Secret: jsii.String(conf.Link.Secret),
			},
		},
		DestinationKafkaCluster: &clusterlink.ClusterLinkDestinationKafkaCluster{
			Id:           jsii.String(*source.Kafka.Id()),
			RestEndpoint: jsii.String(*source.Kafka.RestEndpoint()),
			Credentials: &clusterlink.ClusterLinkDestinationKafkaClusterCredentials{
				Key:    jsii.String(*source.Account.ApiKey.Id()),
				Secret: jsii.String(*source.Account.ApiKey.Secret()),
			},
		},
		Config: &map[string]*string{
			"auto.create.mirror.topics.enable":  jsii.String("true"),
			"consumer.offset.sync.enable":       jsii.String("true"),
			"auto.create.mirror.topics.filters": jsii.String("{ \"topicFilters\":[ { \"name\": \"*\", \"patternType\": \"LITERAL\", \"filterType\": \"INCLUDE\" } ] }"),
			"acl.sync.enable":                   jsii.String("true"),
			"acl.sync.ms":                       jsii.String("1000"),
			"acl.filters":                       jsii.String("{ \"aclFilters\": [ { \"resourceFilter\": {\"resourceType\": \"any\", \"patternType\": \"any\" }, \"accessFilter\": { \"operation\": \"any\", \"permissionType\": \"any\"}}]}"),
		},
	})

	cdktf.NewTerraformOutput(scope, jsii.String("cluster_link_id"), &cdktf.TerraformOutputConfig{
		Value: link.Id(),
	})

	return &ClusterLink{}
}

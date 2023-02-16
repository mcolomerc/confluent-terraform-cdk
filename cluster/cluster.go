package cluster

import (
	"fmt"

	"cdk.tf/go/stack/config"
	"cdk.tf/go/stack/generated/confluentinc/confluent/apikey"
	"cdk.tf/go/stack/generated/confluentinc/confluent/dataconfluentserviceaccount"
	"cdk.tf/go/stack/generated/confluentinc/confluent/kafkacluster"
	"cdk.tf/go/stack/generated/confluentinc/confluent/rolebinding"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClusterInputs struct {
	Config cdktf.TerraformVariable
}
type KafkaCluster struct {
	Kafka   kafkacluster.KafkaCluster
	Account ClusterAccount
}
type ClusterAccount struct {
	ApiKey apikey.ApiKey
}

const ClusterRole = "CloudClusterAdmin"

var kafkaCluster KafkaCluster

func NewKafkaCluster(scope constructs.Construct, name *string) *KafkaCluster {

	c := constructs.NewConstruct(scope, name)

	conf := config.GetConfing()

	// Kafka cluster
	cluster := kafkacluster.NewKafkaCluster(c, jsii.String("kafka_cluster"), &kafkacluster.KafkaClusterConfig{
		Availability: jsii.String(string(conf.Availability)), // SINGLE_ZONE or MULTI_ZONE
		Cloud:        jsii.String(string(conf.Cloud)),        // AWS, GCP or AZURE
		DisplayName:  jsii.String(conf.DisplayName),          // Display name of the cluster
		Environment: &kafkacluster.KafkaClusterEnvironment{
			Id: jsii.String(conf.Environment), // Environment ID
		},
		Region: jsii.String(conf.Region), // Region
	})

	switch conf.Type {
	case config.Basic:
		cluster.Basic().SetInternalValue(&[]*kafkacluster.KafkaClusterBasic{{}})
	case config.Standard:
		cluster.Standard().SetInternalValue(&[]*kafkacluster.KafkaClusterStandard{{}})
	case config.Dedicated:
		cluster.Dedicated().SetInternalValue(&kafkacluster.KafkaClusterDedicated{
			Cku: jsii.Number(float64(conf.Cku)),
		})
	default:
		fmt.Println("Invalid cluster type")
	}

	cdktf.NewTerraformOutput(c, jsii.String("cluster_"), &cdktf.TerraformOutputConfig{
		Value: cluster.Id(),
	})

	kafkaCluster.Kafka = cluster
	if conf.ServiceAccount != "" {
		kafkaCluster.Account = *NewClusterApiKey(c, jsii.String("cluster-admin-service-account")) // Create a new cluster API key for the service account
	}
	return &kafkaCluster
}

func NewClusterApiKey(scope constructs.Construct, name *string) *ClusterAccount {
	c := constructs.NewConstruct(scope, name)

	conf := config.GetConfing()

	serviceAccount := dataconfluentserviceaccount.NewDataConfluentServiceAccount(c, jsii.String("saccount"),
		&dataconfluentserviceaccount.DataConfluentServiceAccountConfig{
			DisplayName: jsii.String(conf.ServiceAccount),
		})

	// Service account role binding
	rolebinding.NewRoleBinding(scope, jsii.String("saccount_role"), &rolebinding.RoleBindingConfig{
		Principal:  jsii.String(fmt.Sprintf("User:%s", *serviceAccount.Id())),
		RoleName:   jsii.String(ClusterRole),
		CrnPattern: jsii.String(*kafkaCluster.Kafka.RbacCrn()),
	})

	apiKey := apikey.NewApiKey(scope, jsii.String("saccount_api_key"), &apikey.ApiKeyConfig{
		DisplayName: jsii.String(fmt.Sprintf("%s_api_key", conf.ServiceAccount)),
		Description: jsii.String(fmt.Sprintf("API Key that is owned by %s service account", conf.ServiceAccount)),
		Owner: &apikey.ApiKeyOwner{
			Id:         jsii.String(*serviceAccount.Id()),
			ApiVersion: jsii.String(*serviceAccount.ApiVersion()),
			Kind:       jsii.String(*serviceAccount.Kind()),
		},
		ManagedResource: &apikey.ApiKeyManagedResource{
			Id:         jsii.String(*kafkaCluster.Kafka.Id()),
			ApiVersion: jsii.String(*kafkaCluster.Kafka.ApiVersion()),
			Kind:       jsii.String(*kafkaCluster.Kafka.Kind()),
			Environment: &apikey.ApiKeyManagedResourceEnvironment{
				Id: jsii.String(conf.Environment),
			},
		},
	})
	cdktf.NewTerraformOutput(scope, jsii.String("key_secret"), &cdktf.TerraformOutputConfig{
		Value:     apiKey.Secret(),
		Sensitive: jsii.Bool(true),
	})

	cdktf.NewTerraformOutput(scope, jsii.String("key_id"), &cdktf.TerraformOutputConfig{
		Value: apiKey.Id(),
	})

	return &ClusterAccount{
		ApiKey: apiKey,
	}
}

// Path: cluster/cluster.go

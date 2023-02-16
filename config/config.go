package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Cluster struct {
	Cloud          Cloud        `yaml:"cloud"`
	Region         string       `yaml:"region"`
	DisplayName    string       `yaml:"display_name"`
	Availability   Availability `yaml:"availability"`
	Type           Type         `yaml:"type"`
	Cku            int          `yaml:"cku"`
	ServiceAccount string       `yaml:"serviceAccount"`
}
type Config struct {
	Environment string `yaml:"environment"`
	Cluster     `yaml:"cluster"`
}

var config *Config

func GetConfing() *Config {
	if config == nil {
		config = NewConfing()
	}
	return config
}

func NewConfing() *Config {
	c := &Config{
		Environment: "dev",
		Cluster: Cluster{
			Cloud:        GCP,
			Region:       "us-central1",
			DisplayName:  "test-cluster",
			Availability: SingleZone,
			Type:         Basic,
			Cku:          0,
		},
	}
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	if c.Validate() {
		return c
	} else {
		return nil
	}
}

func (c Config) Validate() bool {
	if c.Cluster.Type != Dedicated && c.Cluster.Type != Basic && c.Cluster.Type != Standard {
		log.Println("Invalid cluster type. Must be basic, standard or dedicated")
		return false
	}
	if c.Cluster.Cloud != AWS && c.Cluster.Cloud != GCP && c.Cluster.Cloud != Azure {
		log.Println("Invalid cloud provider. Must be AWS, GCP or Azure")
		return false
	}
	if c.Cluster.Availability != SingleZone && c.Cluster.Availability != MultiZone {
		log.Println("Invalid Availability type")
		return false
	}
	if c.Cluster.Type == Dedicated && c.Cluster.Cku == 0 {
		log.Println("Dedicated cluster must have CKU set")
		log.Println("Setting CKU to 1")
		c.Cluster.Cku = 1
	}
	if c.Cluster.Type == Basic && c.Cluster.Availability != SingleZone {
		log.Println("Basic cluster must be single zone")
		c.Cluster.Availability = SingleZone
	}
	return true
}

type Type string

const (
	Basic     Type = "basic"
	Standard  Type = "standard"
	Dedicated Type = "dedicated"
)

type Availability string

const (
	SingleZone Availability = "SINGLE_ZONE"
	MultiZone  Availability = "MULTI_ZONE"
)

type Cloud string

const (
	GCP   Cloud = "GCP"
	AWS   Cloud = "AWS"
	Azure Cloud = "AZURE"
)

package config

import (
	"fmt"
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
	Link           `yaml:"link"`
}

type Link struct {
	Source string `yaml:"source"`
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
}

type Config struct {
	Environment string `yaml:"environment"`
	Cluster     `yaml:"cluster"`
}

var config *Config

func GetConfing() *Config {
	if config == nil {
		log.Println("Building new config")
		config = NewConfing()
	}
	return config
}

func NewConfing() *Config {
	log.Println("Building config")
	c := &Config{
		Environment: "dev",
		Cluster: Cluster{
			Cloud:        GCP,
			Region:       "us-central1",
			DisplayName:  "test-cluster",
			Availability: SingleZone,
			Type:         Basic,
			Cku:          0,
			Link:         Link{},
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
		fmt.Println("validation -- ", c)
		return c
	} else {
		log.Println("Invalid config")
		return nil
	}
}

func (c *Config) Validate() bool {
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
		c.Cluster.Availability = SingleZone
	}

	if c.Cluster.Type == Basic && c.Cluster.Availability != SingleZone {
		log.Println("Basic cluster must be single zone")
		c.Cluster.Availability = SingleZone
	}
	if c.Link != (Link{}) {
		c.Cluster.Type = Dedicated
		if c.Link.Source == "" {
			log.Println("Link source must be set")
			return false
		}
		if c.Link.Key == "" {
			log.Println("Link key must be set")
			return false
		}
		if c.Link.Secret == "" {
			log.Println("Link secret must be set")
			return false
		}
	}
	if c.Cluster.Type == Dedicated && c.Cluster.Cku == 0 {
		log.Println("Dedicated cluster must have CKU set")
		log.Println("Setting CKU to 1")
		c.Cluster.Cku = 1
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

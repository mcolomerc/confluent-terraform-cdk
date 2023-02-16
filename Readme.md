# Confluent Terraform Provider & Terraform CDK

**Stack > Create a Confluent Cloud Cluster.**

Example: [Confluent Terraform Provider](https://registry.terraform.io/providers/confluentinc/confluent/latest/docs) with [Terraform CDK](https://developer.hashicorp.com/terraform/cdktf)

Requires [Go 1.16+](https://golang.org/doc/install) and [Terraform 0.15+](https://www.terraform.io/downloads.html)

- Install Terraform CDK: `brew install cdktf`

- Clone this repo.

- Add the Confluent Provider

`cdktf provider add confluentinc/confluent`

```sh
Checking whether pre-built provider exists for the following constraints:
  provider: confluentinc/confluent
  version : latest
  language: go
  cdktf   : 0.15.2

Pre-built provider does not exist for the given constraints.
Adding local provider registry.terraform.io/confluentinc/confluent with version constraint undefined to cdktf.json
Local providers have been updated. Running cdktf get to update...
Generated go constructs in the output directory: generated

The generated code depends on jsii-runtime-go. If you haven't yet installed it, you can run go mod tidy to automatically install it.
```

`cdktf provider list`

```sh
┌────────────────────────┬──────────────────┬─────────┬────────────┬──────────────────────────────────────────────────┬─────────────────┐
│ Provider Name          │ Provider Version │ CDKTF   │ Constraint │ Package Name                                     │ Package Version │
├────────────────────────┼──────────────────┼─────────┼────────────┼──────────────────────────────────────────────────┼─────────────────┤
│ confluentinc/confluent │ 1.31.0           │         │ ~> 1.31    │                                                  │                 │
└────────────────────────┴──────────────────┴─────────┴────────────┴──────────────────────────────────────────────────┴─────────────────┘
```

---

Configure the Confluent Provider, provide Confluent Cloud credentials

Using tfvars file:  **sample.tfvars**

```hcl
confluent_cloud_api_key="CONFLUENT_CLOUD_API_KEY" 

confluent_cloud_api_secret="CONFLUENT_CLOUD_API_SECRET"
```

**Deploy:** `cdktf deploy --var-file=./sample.tfvars` 

Alternatives:

- Define environment variables: `TF_VAR_imageId=ami-abcde123`
- `--var` CLI option: `cdktf deploy --var='imageId=ami-abcde123'`
- `--var-file` CLI option: `cdktf deploy --var-file=/path/to/variables.tfvars`

Deploys all stacks & auto approve: `cdktf deploy --auto-approve '*'`  

**Destroy** `cdktf destroy`

## Stack Configuration Options

Create a `config.yaml` file with the following options:

```yaml
environment:  #confluent cloud environment id
# new cluster configuration
cluster: 
  cloud: # GCP, AWS or Azure
  region: # cloud region
  display_name: # cluster name
  availability: # SINGLE_ZONE or MULTI_ZONE - basic are SINGLE_ZONE only. 
  type: # basic, standard or dedicated
  cku: # optional: for dedicated clusters, default = 1
  serviceAccount: # optional: create a cluster API_KEY for the given service account name.
```

---

Mirror a source Confluent Cloud cluster with Cluster Link 

```yaml
environment: #confluent cloud environment id
# new cluster configuration
cluster: 
  cloud: # GCP, AWS or Azure
  region: # cloud region
  display_name: # cluster name 
  serviceAccount: # optional: create a cluster API_KEY for the SACC.
  # when using source mirror, destination is a dedicated cluster
  link:
    source: # source cluster id 
    key: # source cluster api key 
    secret: # source cluster api secret
```

Destroy: 

`confluent kafka mirror promote <topic_1> <topic_2> ... <topic_n>  --link <link_name> --cluster <cluster_id>`

`confluent kafka link delete <link_name> --cluster <cluster_id>`




## Terraform CDK help

- Create Terraform CDK - Golang project

`cdktf init --template=go --local`

Your cdktf go project is ready!

  `cat help`                Prints this message

  Compile:
    `go build`              Builds your go project

  Synthesize:
    `cdktf synth [stack]`   Synthesize Terraform resources to cdktf.out/

  Diff:
    `cdktf diff [stack]`    Perform a diff (terraform plan) for the given stack

  Deploy:
    `cdktf deploy [stack]`  Deploy the given stack

  Destroy:
    `cdktf destroy [stack]` Destroy the given stack

  Learn more about using modules and providers <https://cdk.tf/modules-and-providers>





# Confluent Terraform Provider & Terraform CDK

Example: [Confluent Terraform Provider](https://registry.terraform.io/providers/confluentinc/confluent/latest/docs) with [Terraform CDK](https://developer.hashicorp.com/terraform/cdktf)

**Stack > Create a Dedicated Cluster.**

- Install Terraform CDK

`brew install cdktf`


- Create Terraform CDK - Golang project 

`cdktf init --template=go --local`

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
│ confluentinc/confluent │ 1.28.0           │         │ ~> 1.28    │                                                  │                 │
└────────────────────────┴──────────────────┴─────────┴────────────┴──────────────────────────────────────────────────┴─────────────────┘
```

- Deploy: `cdktf deploy`

- Destroy: `cdktf destroy`

## Terraform CDK help

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

  Learn more about using modules and providers https://cdk.tf/modules-and-providers

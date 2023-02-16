package main

import (
	"testing"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// The tests below are example tests, you can find more information at
// https://cdk.tf/testing
var stackS string = "confluent-cluster"
var stack = cdktf.NewTerraformStack(cdktf.Testing_App(nil), &stackS)

func TestCheckValidity(t *testing.T) {
	// We need to do a full synth to validate the terraform configuration
	assertion := cdktf.Testing_ToBeValidTerraform(cdktf.Testing_FullSynth(stack))

	if !*assertion {
		t.Error(assertion)
	}
}

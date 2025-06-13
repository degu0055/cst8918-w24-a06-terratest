package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAzureLinuxVMCreation(t *testing.T) {
	labelPrefix := "degu0055"

	terraformOptions := &terraform.Options{
		// The path to where your Terraform code is located
		TerraformDir: "../",

		Vars: map[string]interface{}{
			"labelPrefix": labelPrefix,
		},
	}

	// Run terraform init and apply
	terraform.InitAndApply(t, terraformOptions)

	// Add your tests here — example: check output variables, etc.
	output := terraform.Output(t, terraformOptions, "vm_name")
	assert.NotEmpty(t, output)

	// Clean up
	terraform.Destroy(t, terraformOptions)
}

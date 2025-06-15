package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNicConnectedToVM(t *testing.T) {
	// Load Terraform output from root directory
	terraformOptions := &terraform.Options{
		TerraformDir: "../", // adjust if your test directory is elsewhere
	}

	// Pull values from Terraform output
	subscriptionID := "cdb9bdf3-e7ee-43e9-8f6c-ba7327868df1" // <-- replace with your actual subscription ID
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	expectedNicName := terraform.Output(t, terraformOptions, "nic_name")

	// Get NICs attached to the VM
	nicList := azure.GetVirtualMachineNics(t, vmName, resourceGroupName, subscriptionID)

	assert.NotEmpty(t, nicList, "NIC list should not be empty")
	assert.True(t, containsIgnoreCase(nicList, expectedNicName), "Expected NIC is not connected to the VM")
}

// Helper function to check case-insensitive match
func containsIgnoreCase(list []string, target string) bool {
	for _, item := range list {
		if strings.EqualFold(item, target) {
			return true
		}
	}
	return false
}

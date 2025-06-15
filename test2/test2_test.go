package test2

import (
	"io/ioutil"
	"testing"

	"github.com/gruntwork-io/terratest/modules/ssh"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestUbuntuVersion(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
	}

	vmPublicIP := terraform.Output(t, terraformOptions, "public_ip")

	// Read private key file contents
	privateKeyPath := "/Users/romeodeguzmanii/.ssh/id_rsa" // <--- update to your real key path
	privateKeyBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		t.Fatalf("Failed to read private key file: %v", err)
	}

	sshHost := ssh.Host{
		Hostname:    vmPublicIP,
		SshUserName: "azureadmin",
		SshKeyPair: &ssh.KeyPair{
			PrivateKey: string(privateKeyBytes),
		},
	}

	output, err := ssh.CheckSshCommandE(t, sshHost, "lsb_release -a")
	assert.NoError(t, err)
	assert.Contains(t, output, "Ubuntu 22.04")
}

package testvNetCreation

// importing the Golang native library called testing from github repository
import (
	"github.com/gruntwork-io/terratest/modules/terraform"
	"testing"
)
// creating the function to run the tests
func TestTerraformAzureNetworkingExample(t *testing.T) {
	t.Parrell()

	terraformOptions := &terraform.Options{		
	}
	// At the end of the test, run 'terraform destroy' to clean up any resources
	defer terraform.Destroy(t, terraformOptions)

	// This will run 'terraform init' and 'terraform apply' and failt the test if not valid
	terraform.InitAndAppy(t, terraformOptions)
}

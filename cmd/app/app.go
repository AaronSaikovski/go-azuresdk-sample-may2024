package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/AaronSaikovski/go-azuresdk-sample-may2024/internal/pkg/auth"
	"github.com/AaronSaikovski/go-azuresdk-sample-may2024/internal/pkg/resourcegroups"
	"github.com/logrusorgru/aurora"
)

// Define key global variables.
var (
	subscriptionId    = os.Getenv("AZURE_SUBSCRIPTION_ID")
	location          = "australiaeast"
	resourceGroupName = "test-go-rsg"
)

// run - main run method
func Run(versionString string) error {

	/* ********************************************************************** */

	// Create a context with cancellation capability
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	/* ********************************************************************** */

	// Get default cred
	cred, err := auth.GetAzureDefaultCredential()
	if err != nil {
		return err
	}

	/* ********************************************************************** */
	// check we are logged into the Azure source subscription
	if !auth.GetLogin(ctx, subscriptionId) {
		return fmt.Errorf("you are not logged into the azure subscription '%s', please login and retry operation", subscriptionId)
	}
	fmt.Println(aurora.Sprintf(aurora.Yellow("Logged into Subscription Id: %s\n"), subscriptionId))

	/* ********************************************************************** */

	//Get the resource group client
	resourceGroupClient, err := resourcegroups.GetResourceGroupClient(cred, subscriptionId)
	if err != nil {
		return err
	}

	/* ********************************************************************** */

	// create resource group
	resourceGroup, err := resourcegroups.CreateResourceGroup(ctx, location, resourceGroupName, resourceGroupClient)
	if err != nil {
		log.Fatalf("Creation of resource group failed: %+v", err)
	}

	// Print the name of the new resource group.
	fmt.Println(aurora.Sprintf(aurora.Yellow("Resource group %s created"), *resourceGroup.ResourceGroup.ID))

	return nil
}

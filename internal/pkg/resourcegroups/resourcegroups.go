package resourcegroups

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// GetResourceGroupClient creates a new instance of the armresources.ResourceGroupsClient for the given azidentity.DefaultAzureCredential and subscriptionID.
//
// Parameters:
// - cred: The azidentity.DefaultAzureCredential used to authenticate the client.
// - subscriptionID: The ID of the subscription to create the client for.
//
// Returns:
// - *armresources.ResourceGroupsClient: The created client instance.
// - error: An error if the client creation fails.
func GetResourceGroupClient(cred *azidentity.DefaultAzureCredential, subscriptionID string) (*armresources.ResourceGroupsClient, error) {
	// Create a new Resource Groups client factory
	resourcesClientFactory, err := armresources.NewClientFactory(subscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}

	// Create a new Resource Groups client
	resourceGroupClient := resourcesClientFactory.NewResourceGroupsClient()
	if resourceGroupClient == nil {
		return nil, errors.New("failed to create resource group client")
	}

	return resourceGroupClient, nil
}

// GetResourceGroup retrieves a resource group using the provided resource group name.
//
// ctx: The context within which the function is being executed.
// resourceGroupName: The name of the resource group to retrieve.
// Returns a pointer to armresources.ResourceGroup and an error.
// func GetResourceGroup(ctx context.Context, resourceGroupClient *armresources.ResourceGroupsClient, resourceGroupName string) (*armresources.ResourceGroup, error) {
func GetResourceGroup(ctx context.Context, resourceGroupClient *armresources.ResourceGroupsClient, resourceGroupName string) (*armresources.ResourceGroup, error) {
	resourceGroupResp, err := resourceGroupClient.Get(ctx, resourceGroupName, nil)
	if err != nil {
		return nil, err
	}
	return &resourceGroupResp.ResourceGroup, nil
}

// GetResourceGroupId retrieves the ID of a resource group.
//
// ctx: The context within which the function is being executed.
// resourceGroupClient: The client for interacting with Azure Resource Groups.
// resourceGroupName: The name of the resource group.
// *string, error: The ID of the resource group and an error if any.
func GetResourceGroupId(ctx context.Context, resourceGroupClient *armresources.ResourceGroupsClient, resourceGroupName string) (*string, error) {

	resourceGroupResp, err := resourceGroupClient.Get(ctx, resourceGroupName, nil)
	if err != nil {
		return nil, err
	}
	return resourceGroupResp.ResourceGroup.ID, nil
}

// ListResourceGroup fetches a list of resource groups.
//
// ctx - the context within which the function is executed.
// []*armresources.ResourceGroup, error - returns a slice of resource groups and an error if any.
// func ListResourceGroup(ctx context.Context, resourceGroupClient *armresources.ResourceGroupsClient) ([]*armresources.ResourceGroup, error) {
func ListResourceGroup(ctx context.Context, resourceGroupClient *armresources.ResourceGroupsClient) ([]*armresources.ResourceGroup, error) {
	resultPager := resourceGroupClient.NewListPager(nil)

	var resourceGroups []*armresources.ResourceGroup
	for resultPager.More() {
		pageResp, err := resultPager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		resourceGroups = append(resourceGroups, pageResp.ResourceGroupListResult.Value...)
	}
	return resourceGroups, nil

}

// CheckResourceGroupExists checks if a resource group exists.
//
// ctx: the context for the request.
// resourceGroupName: the name of the resource group to check.
// (bool, error): returns a boolean indicating if the resource group exists and an error if any.
func CheckResourceGroupExists(ctx context.Context, resourceGroupClient *armresources.ResourceGroupsClient, resourceGroupName string) (bool, error) {

	boolResp, err := resourceGroupClient.CheckExistence(ctx, resourceGroupName, nil)
	if err != nil {
		return false, err
	}
	return boolResp.Success, nil
}

// CreateResourceGroup creates a new resource group in Azure.
//
// ctx: The context within which the function is being executed.
// location: The location where the resource group should be created.
// resourceGroupName: The name of the resource group to be created.
// *armresources.ResourceGroupsClient: The client for interacting with Azure Resource Groups.
// (armresources.ResourceGroupsClientCreateOrUpdateResponse, error): The response from the CreateOrUpdate method and an error if any.
func CreateResourceGroup(ctx context.Context, location string, resourceGroupName string, rgClient *armresources.ResourceGroupsClient) (armresources.ResourceGroupsClientCreateOrUpdateResponse, error) {
	param := armresources.ResourceGroup{
		Location: to.Ptr(location),
	}

	return rgClient.CreateOrUpdate(ctx, resourceGroupName, param, nil)
}

// graphbeta_device_categories.go
// Graph Beta Api - Intune: Device Categories
// Documentation:
// Intune location: https://intune.microsoft.com/#view/Microsoft_Intune_DeviceSettings/DevicesMenu/~/deviceCategories
// API reference: https://learn.microsoft.com/en-us/graph/api/resources/intune-shared-devicecategory?view=graph-rest-beta
// Microsoft Graph requires the structs to support a JSON data structure.

package intune

import (
	"fmt"

	shared "github.com/deploymenttheory/go-api-sdk-m365/sdk/shared"
)

const uriBetaDeviceCategories = "/beta/deviceManagement/deviceCategories"

// ResponseDeviceCategoriesList is used to parse the list response of Device Categories from Microsoft Graph API.
type ResponseDeviceCategoriesList struct {
	ODataContext string                   `json:"@odata.context"`
	Value        []ResourceDeviceCategory `json:"value"`
}

// ResourceDeviceCategory represents an individual Device Category resource from Microsoft Graph API.
type ResourceDeviceCategory struct {
	OdataType   string `json:"@odata.type"`
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}

// GetDeviceCategories retrieves a list of Intune Device Categories from Microsoft Graph API.
func (c *Client) GetDeviceCategories() (*ResponseDeviceCategoriesList, error) {
	endpoint := uriBetaDeviceCategories

	var responseDeviceCategories ResponseDeviceCategoriesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &responseDeviceCategories)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device categories", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseDeviceCategories, nil
}

// GetDeviceCategoryByID retrieves a specific Device Category by its ID from Microsoft Graph API.
func (c *Client) GetDeviceCategoryByID(deviceCategoryId string) (*ResourceDeviceCategory, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceCategories, deviceCategoryId)

	var deviceCategory ResourceDeviceCategory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &deviceCategory)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByID, "device category", deviceCategoryId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &deviceCategory, nil
}

// GetDeviceCategoryByDisplayName retrieves a specific Device Category by its name from Microsoft Graph API.
func (c *Client) GetDeviceCategoryByDisplayName(categoryDisplayName string) (*ResourceDeviceCategory, error) {
	// Retrieve all device categories
	categoriesList, err := c.GetDeviceCategories()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device categories", err)
	}

	// Search for the category with the matching name
	var categoryID string
	for _, category := range categoriesList.Value {
		if category.DisplayName == categoryDisplayName {
			categoryID = category.ID
			break
		}
	}

	if categoryID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "device category", categoryDisplayName, "Category not found")
	}

	// Retrieve the full details of the category using its ID
	return c.GetDeviceCategoryByID(categoryID)
}

// CreateDeviceCategory creates a new Device Category in Microsoft Graph API.
func (c *Client) CreateDeviceCategory(request *ResourceDeviceCategory) (*ResourceDeviceCategory, error) {
	endpoint := uriBetaDeviceCategories

	request.OdataType = "#microsoft.graph.deviceCategory"

	var responseCreatedCategory ResourceDeviceCategory
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &responseCreatedCategory)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedCreate, "device category", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseCreatedCategory, nil
}

// UpdateDeviceCategoryByID updates a specific Device Category identified by its ID.
func (c *Client) UpdateDeviceCategoryByID(deviceCategoryId string, updateRequest *ResourceDeviceCategory) (*ResourceDeviceCategory, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceCategories, deviceCategoryId)

	// Set OdataType to empty since it's not required for update requests
	updateRequest.OdataType = "#microsoft.graph.deviceCategory"

	var updatedCategory ResourceDeviceCategory
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, updateRequest, &updatedCategory)
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedUpdate, "device category", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedCategory, nil
}

// UpdateDeviceCategoryByDisplayName updates a specific Device Category identified by its name.
func (c *Client) UpdateDeviceCategoryByDisplayName(categoryName string, updateRequest *ResourceDeviceCategory) (*ResourceDeviceCategory, error) {
	// Retrieve all device categories
	categoriesList, err := c.GetDeviceCategories()
	if err != nil {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGet, "device categories", err)
	}

	// Search for the category with the matching name and get its ID
	var categoryID string
	for _, category := range categoriesList.Value {
		if category.DisplayName == categoryName {
			categoryID = category.ID
			break
		}
	}

	if categoryID == "" {
		return nil, fmt.Errorf(shared.ErrorMsgFailedGetByName, "device category", categoryName, "Category not found")
	}

	// Update the category using its ID
	return c.UpdateDeviceCategoryByID(categoryID, updateRequest)
}

// DeleteDeviceCategoryByID deletes a specific Device Category identified by its ID.
func (c *Client) DeleteDeviceCategoryByID(deviceCategoryId string) error {
	endpoint := fmt.Sprintf("%s/%s", uriBetaDeviceCategories, deviceCategoryId)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedDeleteByID, "device category", deviceCategoryId, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteDeviceCategoryByDisplayName deletes a specific Device Category identified by its name.
func (c *Client) DeleteDeviceCategoryByDisplayName(categoryName string) error {
	// Retrieve all device categories
	categoriesList, err := c.GetDeviceCategories()
	if err != nil {
		return fmt.Errorf(shared.ErrorMsgFailedGet, "device categories", err)
	}

	// Search for the category with the matching name and get its ID
	var categoryID string
	for _, category := range categoriesList.Value {
		if category.DisplayName == categoryName {
			categoryID = category.ID
			break
		}
	}

	if categoryID == "" {
		return fmt.Errorf(shared.ErrorMsgFailedGetByName, "device category", categoryName, "Category not found")
	}

	// Delete the category using its ID
	return c.DeleteDeviceCategoryByID(categoryID)
}

package intune

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// UnmarshalJSON is a custom unmarshaler for DynamicValue, allowing it to
// dynamically handle different data types in JSON. When unmarshaling JSON data,
// it attempts to decode the value into one of several basic data types:
// int, float64, bool, or string. If none of these types are compatible,
// the value is treated as a raw JSON message. This approach provides flexibility
// in handling various types of data that might be encountered in a graph JSON payload,
// ensuring that the DynamicValue struct can accommodate a wide range of
// JSON structures and types.
func (dv *DynamicValue) UnmarshalJSON(data []byte) error {
	// First, try to unmarshal into the simplest types (int, float, bool, string)
	var intValue int
	if err := json.Unmarshal(data, &intValue); err == nil {
		dv.Value = intValue
		return nil
	}

	var floatValue float64
	if err := json.Unmarshal(data, &floatValue); err == nil {
		dv.Value = floatValue
		return nil
	}

	var boolValue bool
	if err := json.Unmarshal(data, &boolValue); err == nil {
		dv.Value = boolValue
		return nil
	}

	var stringValue string
	if err := json.Unmarshal(data, &stringValue); err == nil {
		dv.Value = stringValue
		return nil
	}

	// If none of the simple types match, treat as a raw JSON message
	dv.Value = json.RawMessage(data)
	return nil
}

// GetDecryptedOmaSetting makes a request to Microsoft Graph API to retrieve the plain text value of an encrypted OMA setting.
// It constructs the endpoint URL using the provided base URL, profile ID, and secret reference value ID.
// The function returns the decrypted value of the OMA setting or an error if the retrieval is unsuccessful.
func (c *Client) GetDecryptedOmaSetting(baseURL, profileId, secretReferenceValueId string) (string, error) {
	endpoint := fmt.Sprintf("%s/%s/getOmaSettingPlainTextValue(secretReferenceValueId='%s')", baseURL, profileId, secretReferenceValueId)

	var decryptedValue struct {
		Value string `json:"value"`
	}

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &decryptedValue)
	if err != nil {
		return "", fmt.Errorf("failed to get decrypted OMA setting: %v", err)
	}

	// Check if the HTTP request was successful
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	return decryptedValue.Value, nil
}

package intune

import "encoding/json"

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

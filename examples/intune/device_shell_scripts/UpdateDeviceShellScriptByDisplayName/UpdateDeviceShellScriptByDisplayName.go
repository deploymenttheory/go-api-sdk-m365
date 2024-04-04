package main

import (
	"encoding/json"
	"fmt"
	"log"

	// Import http_client for logging

	"github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/msgraph/clientconfig.json"

	// Initialize the msgraph client with the HTTP client configuration
	client, err := intune.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create an update request for the Device Shell Script
	updateRequest := &intune.ResourceDeviceShellScript{
		ExecutionFrequency:          "PT15M",
		RetryCount:                  1,
		BlockExecutionNotifications: true,
		DisplayName:                 "intune - Updated macOS script By Display Name",
		Description:                 "Updated Description",
		ScriptContent:               "IyEvYmluL2Jhc2gKCiMjIwojCiMgICAgICAgICAgICBOYW1lOiAgVXBkYXRlIFNTSCBQdWJsaWMgS2V5LnNoCiMgICAgIERlc2NyaXB0aW9uOiAgQWRkcyBuZXcgU1NIIHB1YmxpYyBrZXkgdG8gc3BlY2lmaWVkIHVzZXIgYWNjb3VudCBmb3IKIyAgICAgICAgICAgICAgICAgICByZW1vdGUgYWNjZXNzLCByZXBsYWNpbmcgYW55IGV4aXN0aW5nIGtleXMuCiMgICAgICAgICBDcmVhdGVkOiAgMjAxNy0wMi0xNAojICAgTGFzdCBNb2RpZmllZDogIDIwMjAtMDctMDgKIyAgICAgICAgIFZlcnNpb246ICAxLjIuMwojCiMKIyBDb3B5cmlnaHQgMjAxNyBQYWxhbnRpciBUZWNobm9sb2dpZXMsIEluYy4KIwojIExpY2Vuc2VkIHVuZGVyIHRoZSBBcGFjaGUgTGljZW5zZSwgVmVyc2lvbiAyLjAgKHRoZSAiTGljZW5zZSIpOwojIHlvdSBtYXkgbm90IHVzZSB0aGlzIGZpbGUgZXhjZXB0IGluIGNvbXBsaWFuY2Ugd2l0aCB0aGUgTGljZW5zZS4KIyBZb3UgbWF5IG9idGFpbiBhIGNvcHkgb2YgdGhlIExpY2Vuc2UgYXQKIwojIGh0dHA6Ly93d3cuYXBhY2hlLm9yZy9saWNlbnNlcy9MSUNFTlNFLTIuMAojCiMgVW5sZXNzIHJlcXVpcmVkIGJ5IGFwcGxpY2FibGUgbGF3IG9yIGFncmVlZCB0byBpbiB3cml0aW5nLCBzb2Z0d2FyZQojIGRpc3RyaWJ1dGVkIHVuZGVyIHRoZSBMaWNlbnNlIGlzIGRpc3RyaWJ1dGVkIG9uIGFuICJBUyBJUyIgQkFTSVMsCiMgV0lUSE9VVCBXQVJSQU5USUVTIE9SIENPTkRJVElPTlMgT0YgQU5ZIEtJTkQsIGVpdGhlciBleHByZXNzIG9yIGltcGxpZWQuCiMgU2VlIHRoZSBMaWNlbnNlIGZvciB0aGUgc3BlY2lmaWMgbGFuZ3VhZ2UgZ292ZXJuaW5nIHBlcm1pc3Npb25zIGFuZAojIGxpbWl0YXRpb25zIHVuZGVyIHRoZSBMaWNlbnNlLgojCiMKIyMjCgoKCiMjIyMjIyMjIyMgdmFyaWFibGUtaW5nICMjIyMjIyMjIyMKCgoKIyBKYW1mIFBybyBzY3JpcHQgcGFyYW1ldGVyICJUYXJnZXQgVXNlciIKdGFyZ2V0VXNlcj0iJDQiCiMgSmFtZiBQcm8gc2NyaXB0IHBhcmFtZXRlciAiUHVibGljIFNTSCBLZXkgUGF0aCIKcHVibGljU1NIS2V5UGF0aD0iJDYiCnRhcmdldFVzZXJIb21lPSQoL3Vzci9iaW4vZHNjbCAuIC1yZWFkICIvVXNlcnMvJHRhcmdldFVzZXIiIE5GU0hvbWVEaXJlY3RvcnkgfCAvdXNyL2Jpbi9hd2sgJ3twcmludCAkTkZ9JykKdGFyZ2V0VXNlclNTSFNldHRpbmdzPSIkdGFyZ2V0VXNlckhvbWUvLnNzaCIKdGFyZ2V0VXNlclNTSEF1dGhvcml6ZWRLZXlzUGF0aD0iJHRhcmdldFVzZXJTU0hTZXR0aW5ncy9hdXRob3JpemVkX2tleXMiCgoKCiMjIyMjIyMjIyMgZnVuY3Rpb24taW5nICMjIyMjIyMjIyMKCgoKIyBFeGl0cyBpZiBhbnkgcmVxdWlyZWQgSmFtZiBQcm8gYXJndW1lbnRzIGFyZSB1bmRlZmluZWQuCmZ1bmN0aW9uIGNoZWNrX2phbWZfcHJvX2FyZ3VtZW50cyB7CiAgamFtZlByb0FyZ3VtZW50cz0oCiAgICAiJHRhcmdldFVzZXIiCiAgICAiJHB1YmxpY1NTSEtleVBhdGgiCiAgKQogIGZvciBhcmd1bWVudCBpbiAiJHtqYW1mUHJvQXJndW1lbnRzW0BdfSI7IGRvCiAgICBpZiBbWyAiJGFyZ3VtZW50IiA9ICIiIF1dOyB0aGVuCiAgICAgIGVjaG8gIuKdjCBFUlJPUjogVW5kZWZpbmVkIEphbWYgUHJvIGFyZ3VtZW50LCB1bmFibGUgdG8gcHJvY2VlZC4iCiAgICAgIGV4aXQgNzQKICAgIGZpCiAgZG9uZQp9CgoKIyBFeGl0cyBpZiBwdWJsaWMgU1NIIGtleSBkb2VzIG5vdCBleGlzdC4KZnVuY3Rpb24gcHVibGljX3NzaF9rZXlfY2hlY2sgewogIGlmIFsgISAtZSAiJHB1YmxpY1NTSEtleVBhdGgiIF07IHRoZW4KICAgIGVjaG8gIuKdjCBFUlJPUjogUHVibGljIFNTSCBLZXkgbm90IGZvdW5kIGF0IHNwZWNpZmllZCBwYXRoLCB1bmFibGUgdG8gcHJvY2VlZC4gUGxlYXNlIGNoZWNrIFB1YmxpYyBTU0ggS2V5IFBhdGggcGFyYW1ldGVyIGluIEphbWYgUHJvIHBvbGljeS4iCiAgICBleGl0IDc0CiAgZmkKfQoKCgojIyMjIyMjIyMjIG1haW4gcHJvY2VzcyAjIyMjIyMjIyMjCgoKCiMgRXhpdCBpZiBhbnkgcmVxdWlyZWQgSmFtZiBQcm8gYXJndW1lbnRzIGFyZSB1bmRlZmluZWQuCmNoZWNrX2phbWZfcHJvX2FyZ3VtZW50cwpwdWJsaWNfc3NoX2tleV9jaGVjawoKCgoKCiMgUmVtb3ZlIGV4aXN0aW5nIFNTSCBrZXlzIGlmIHByZXNlbnQuCmlmIFsgLWUgIiR0YXJnZXRVc2VyU1NIQXV0aG9yaXplZEtleXNQYXRoIiBdOyB0aGVuCiAgL2Jpbi9ybSAiJHRhcmdldFVzZXJTU0hBdXRob3JpemVkS2V5c1BhdGgiCmZpCgoKIyBBZGQgLnNzaC9hdXRob3JpemVkX2tleXMgYW5kIHBvcHVsYXRlIHdpdGggdXNlcidzIHB1YmxpYyBrZXkuCi9iaW4vbWtkaXIgLXAgIiR0YXJnZXRVc2VyU1NIU2V0dGluZ3MiCi91c3IvYmluL3RvdWNoICIkdGFyZ2V0VXNlclNTSEF1dGhvcml6ZWRLZXlzUGF0aCIKL2Jpbi9jYXQgIiRwdWJsaWNTU0hLZXlQYXRoIiA+PiAiJHRhcmdldFVzZXJTU0hBdXRob3JpemVkS2V5c1BhdGgiCi91c3Ivc2Jpbi9jaG93biAtUiAiJHRhcmdldFVzZXIiICIkdGFyZ2V0VXNlclNTSFNldHRpbmdzLyIKL2Jpbi9jaG1vZCA3MDAgIiR0YXJnZXRVc2VyU1NIU2V0dGluZ3MiCi9iaW4vY2htb2QgNjAwICIkdGFyZ2V0VXNlclNTSEF1dGhvcml6ZWRLZXlzUGF0aCIKCgoKZXhpdCAwCg==",
		RunAsAccount:                "user",
		FileName:                    "Updated File Name",
		RoleScopeTagIds:             []string{"0"},
	}

	// Replace "scriptName" with the actual ID of the Device Shell Script you want to update
	scriptName := "macOS-shell_script-created_with_intune_withJSON"

	// Update the Device Shell Script by its ID
	updatedShellScript, err := client.UpdateDeviceShellScriptByDisplayName(scriptName, updateRequest)
	if err != nil {
		log.Fatalf("Failed to update Device Shell Script by display name: %v", err)
	}

	// Pretty print the device management scripts
	jsonData, err := json.MarshalIndent(updatedShellScript, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal device management scripts: %v", err)
	}
	fmt.Println(string(jsonData))
}

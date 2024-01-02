package powershell_scripts

import (
	"fmt"
	"log"
	"strings"

	shared "github.com/deploymenttheory/go-api-sdk-m365/export/library"
	intuneSDK "github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
)

func Backup(client *intuneSDK.Client, outputPath, outputFormat string, excludeAssignments bool, prefix string, appendID bool) error {
	log.Println("Starting device management script backup...")

	// Retrieve all Device Management Scripts
	scripts, err := client.GetDeviceManagementScripts()
	if err != nil {
		log.Println("Error getting powershell scripts:", err)
		return err
	}

	log.Printf("Found %d powershell scripts\n", len(scripts.Value))

	// Process each script
	for _, script := range scripts.Value {
		log.Printf("Processing script: %s\n", script.DisplayName)

		// Filter based on prefix
		if prefix != "" && !strings.HasPrefix(script.DisplayName, prefix) {
			log.Printf("Skipping script '%s' due to prefix mismatch\n", script.DisplayName)
			continue
		}

		// Get detailed information for each script
		detailedScript, err := client.GetDeviceManagementScriptByID(script.ID)
		if err != nil {
			log.Println("Error getting script details:", err)
			continue
		}

		// Convert script details to a map
		scriptMap := shared.ConvertStructToMap(detailedScript)
		if scriptMap == nil {
			log.Println("Error converting script details to map")
			continue
		}

		// Construct filename
		filename, ok := scriptMap["displayName"].(string)
		if !ok {
			log.Println("Error: displayName is not a string or is nil")
			continue
		}
		filename = fmt.Sprintf("%s__%s", filename, script.ID)

		// Save Device Management Script
		log.Printf("Saving script to '%s' in format '%s'\n", filename, outputFormat)
		err = shared.SaveOutput(outputFormat, outputPath, filename, scriptMap)
		if err != nil {
			log.Println("Error saving script output:", err)
			continue
		}

		// Save the script content
		log.Printf("Saving script content for '%s'\n", scriptMap["displayName"].(string))
		err = shared.SaveScript(outputPath, filename, "ScriptContent", scriptMap["scriptContent"].(string))
		if err != nil {
			log.Println("Error saving script content:", err)
			continue
		}
	}

	log.Println("Device management script backup completed.")
	return nil
}

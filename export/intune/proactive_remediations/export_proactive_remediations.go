package proactive_remediations

import (
	"fmt"
	"log"
	"strings"

	shared "github.com/deploymenttheory/go-api-sdk-m365/export/library"
	intuneSDK "github.com/deploymenttheory/go-api-sdk-m365/sdk/m365/intune"
)

func Backup(client *intuneSDK.Client, outputPath, outputFormat string, excludeAssignments bool, prefix string, appendID bool) error {
	log.Println("Starting proactive remediation backup...")

	// Retrieve all Proactive Remediations
	remediations, err := client.GetProactiveRemediations()
	if err != nil {
		log.Println("Error getting proactive remediations:", err)
		return err
	}

	log.Printf("Found %d proactive remediations\n", len(remediations.Value))

	// Process each remediation
	for _, remediation := range remediations.Value {
		log.Printf("Processing remediation: %s\n", remediation.DisplayName)

		// Filter based on prefix
		if prefix != "" && !strings.HasPrefix(remediation.DisplayName, prefix) {
			log.Printf("Skipping remediation '%s' due to prefix mismatch\n", remediation.DisplayName)
			continue
		}

		var remediationDetails interface{}
		if !excludeAssignments {
			log.Printf("Getting details for remediation '%s'\n", remediation.DisplayName)
			var err error
			remediationDetails, err = client.GetProactiveRemediationByID(remediation.ID)
			if err != nil {
				log.Println("Error getting remediation details:", err)
				continue
			}
		}

		// Convert remediation details to a map
		remediationMap := shared.ConvertStructToMap(remediationDetails)
		if remediationMap == nil {
			log.Println("Error converting remediation details to map")
			continue
		}

		// Construct filename
		filename, ok := remediationMap["displayName"].(string)
		if !ok {
			log.Println("Error: displayName is not a string or is nil")
			continue
		}

		id, ok := remediationMap["id"].(string)
		if !ok {
			log.Println("Error: 'id' is either nil or not a string")
			continue
		}
		filename = fmt.Sprintf("%s__%s", filename, id)

		// Save Proactive Remediation
		log.Printf("Saving remediation to '%s' in format '%s'\n", filename, outputFormat)
		err = shared.SaveOutput(outputFormat, outputPath, filename, remediationMap)
		if err != nil {
			log.Println("Error saving output:", err)
			continue
		}

		// Save detection and remediation scripts
		log.Printf("Saving detection script for '%s'\n", remediationMap["displayName"].(string))
		err = shared.SaveScript(outputPath, filename, "DetectionScript", remediationMap["detectionScriptContent"].(string))
		if err != nil {
			log.Println("Error saving detection script:", err)
			continue
		}

		log.Printf("Saving remediation script for '%s'\n", remediationMap["displayName"].(string))
		err = shared.SaveScript(outputPath, filename, "RemediationScript", remediationMap["remediationScriptContent"].(string))
		if err != nil {
			log.Println("Error saving remediation script:", err)
			continue
		}
	}

	log.Println("Proactive remediation backup completed.")
	return nil
}

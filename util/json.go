package util

import (
	"encoding/json"
	"log"
)

// MergeJSONData merges JSON data into a struct
func MergeJSONData(target interface{}, data map[string]interface{}) {
	// Convert target to JSON
	jsonData, err := json.Marshal(target)
	if err != nil {
		log.Fatalf("Failed to marshal target: %v", err)
	}

	// Unmarshal into map
	var targetMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &targetMap); err != nil {
		log.Fatalf("Failed to unmarshal target: %v", err)
	}

	// Merge data
	for k, v := range data {
		targetMap[k] = v
	}

	// Convert back to struct
	jsonData, err = json.Marshal(targetMap)
	if err != nil {
		log.Fatalf("Failed to marshal merged data: %v", err)
	}

	if err := json.Unmarshal(jsonData, target); err != nil {
		log.Fatalf("Failed to unmarshal merged data: %v", err)
	}
}

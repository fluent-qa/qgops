package yamlutil

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

func ToStruct(yamlStr string, data any) {
	err := yaml.Unmarshal([]byte(yamlStr), &data)
	if err != nil {
		slog.Error("convert yaml str failed", data)
	}
}

func ToStructureString(any any) string {
	jsonBytes, err := yaml.Marshal(any)
	if err != nil {
		slog.Error("convert yaml str failed", any)
		return ""
	}
	return string(jsonBytes)
}

// NOT WORK
func ToStructFromFile(yamlFile string, data any) any {
	yamlStr, err := os.ReadFile(yamlFile)
	if err != nil {
		slog.Error("read yaml str failed")
	}
	err = yaml.Unmarshal(yamlStr, &data)
	if err != nil {
		slog.Error("convert yaml str failed")
	}
	return data
}

func WriteStructToFile(data any, filePath string) {
	yamlData, err := yaml.Marshal(&data)
	if err != nil {
		slog.Error("Error marshaling to YAML:", err)
		return
	}

	// Write the YAML data to a file
	err = os.WriteFile(filePath, yamlData, 0644)
	if err != nil {
		slog.Error("Error writing YAML Failed:", err)
		return
	}

	slog.Info("YAML file written successfully.")
}

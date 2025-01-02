package jsonutil

import (
	"encoding/json"
	"log"
	"log/slog"
	"os"
)

func ToStruct(jsonstr string, any any) {

	err := json.Unmarshal([]byte(jsonstr), &any)
	if err != nil {
		slog.Error("read json failed", err)
		log.Fatal("convert json str failed", any)
	}
}

func ToString(any any) string {
	jsonBytes, err := json.Marshal(any)
	if err != nil {
		log.Fatal("convert json str failed", any)
		return ""
	}
	return string(jsonBytes)
}

func ToStructFromFile(jsonFile string, data any) {
	jsonBytes, err := os.ReadFile(jsonFile)
	if err != nil {
		slog.Error("read json failed", err)
	}
	ToStruct(string(jsonBytes), &data)
}

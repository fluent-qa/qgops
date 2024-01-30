package utils

import (
	"path/filepath"
	"strings"
)

func GetFileExtension(filePath string) string {

	filePostFix := filepath.Ext(filePath)
	return strings.Replace(filePostFix, ".", "", 1)
}

package utils

import (
	"errors"
	"os"
)

func DirectoryExist(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New(`${path} is existing, but not a Directory`)
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WriteStringToFile(content, targetPath string) error {
	err := os.WriteFile(targetPath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadStringFromFile(targetPath string) (string, error) {
	file, err := os.ReadFile(targetPath)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

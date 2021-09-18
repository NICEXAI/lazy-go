package util

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

// GetCurrentPath get current folder path
func GetCurrentPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(dir, `\`, `/`), nil
}

// GetProjectPath get project path
func GetProjectPath(name string) (string, error) {
	currentPath, err := GetCurrentPath()
	if err != nil {
		return "", err
	}

	return path.Join(currentPath, name), nil
}

// MkdirIfNotExist makes directories if the input path is not exists
func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}

// RemoveFolderIfExist deletes the folder if it is exists
func RemoveFolderIfExist(dir string) error {
	if !IsFolderExist(dir) {
		return nil
	}
	fmt.Println(dir)
	return os.RemoveAll(dir)
}

// IsFolderExist determine if a folder already exists
func IsFolderExist(dir string) bool {
	s, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// GetGoPath get go path from env
func GetGoPath() (string, error) {
	cmd := exec.Command("go", "env", "GOPATH")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(string(out[:len(out)-1]), `\`, `/`), nil
}

// GetGoModulePath get go module path from local
func GetGoModulePath() (string, error) {
	goPath, err := GetGoPath()
	if err != nil {
		return "", err
	}
	return goPath + "/pkg/mod", nil
}

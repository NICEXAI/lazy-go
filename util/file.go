package util

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
	"strings"
)

// GetFileFullName get current file full name
func GetFileFullName(filePath string) string {
	pathArr := strings.Split(filePath, "/")
	return pathArr[len(pathArr)-1]
}

// GetFileName get current file name
func GetFileName(filePath string) string {
	fileName := GetFileFullName(filePath)
	fileArr := strings.Split(fileName, ".")
	if len(fileArr) < 2 {
		return fileArr[0]
	}
	return fileArr[len(fileArr)-2]
}

// GetFileExt get current file ext
func GetFileExt(filePath string) string {
	fileName := GetFileFullName(filePath)
	fileArr := strings.Split(fileName, ".")
	if len(fileArr) < 2 {
		return fileArr[0]
	}
	return fileArr[len(fileArr)-1]
}

// CreateIfNotExist creates a file if it is not exists
func CreateIfNotExist(file string) (*os.File, error) {
	_, err := os.Stat(file)
	if !os.IsNotExist(err) {
		return nil, fmt.Errorf("%s already exist", file)
	}

	return os.Create(file)
}

// RemoveIfExist deletes the specified file if it is exists
func RemoveIfExist(filename string) error {
	if !IsFileExist(filename) {
		return nil
	}
	fmt.Println(filename)
	return os.Remove(filename)
}

// IsFileExist returns true if the specified file is exists
func IsFileExist(file string) bool {
	info, err := os.Stat(file)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// CloneProjectFromRemote clone project from remote
func CloneProjectFromRemote(dir, remote string) error {
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:               "https://" + remote,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Progress:          os.Stdout,
	})
	return err
}

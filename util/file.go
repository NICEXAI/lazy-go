package util

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
)

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

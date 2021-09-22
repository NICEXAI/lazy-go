package config

import (
	"github.com/NICEXAI/go2struct"
	"github.com/NICEXAI/lazy-go/internal/errorx"
	"github.com/NICEXAI/lazy-go/util"
	"os"
	"path"
	"path/filepath"
)

const (
	configTmpl = "// Code generated by lazyGo. DO NOT EDIT.\n\npackage config\n\n"

	configPath = "./example/config"
	configName = "settings.yaml"

	configGoPath = "./example/internal/config"
	configGoName = "options.go"
)

// Sync Auto-sync configuration
func Sync() error {
	//check if the project is valid
	var (
		configAbsPath, configFileAbsPath, configGoAbsPath, configGoFileAbsPath string
		configContent, structContent                                           []byte
		file                                                                   *os.File
		err                                                                    error
	)

	configAbsPath, err = filepath.Abs(configPath)
	if err != nil {
		return err
	}

	if !util.IsFolderExist(configAbsPath) {
		return errorx.SDKConfigFolderNotExist
	}

	configFileAbsPath, err = filepath.Abs(path.Join(configPath, configName))
	if err != nil {
		return err
	}

	if !util.IsFileExist(configFileAbsPath) {
		return nil
	}

	configGoAbsPath, err = filepath.Abs(configGoPath)
	if err != nil {
		return err
	}

	if err = util.MkdirIfNotExist(configGoAbsPath); err != nil {
		return err
	}

	configGoFileAbsPath, err = filepath.Abs(path.Join(configGoPath, configGoName))
	if err != nil {
		return err
	}

	configContent, err = os.ReadFile(configFileAbsPath)
	if err != nil {
		return err
	}

	structContent, err = go2struct.YAML2Struct("options", configContent)
	if err != nil {
		return err
	}

	file, err = os.Create(configGoFileAbsPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(configTmpl + string(structContent)); err != nil {
		return err
	}

	return nil
}

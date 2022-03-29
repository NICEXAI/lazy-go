package project

import (
	"encoding/json"
	"github.com/NICEXAI/ghost"
	"github.com/NICEXAI/lazy-go/internal/errorx"
	"github.com/NICEXAI/lazy-go/util"
	"github.com/fatih/color"
	"os"
	"path"
)

const apiProjectRepo = "github.com/NICEXAI/lazy-scaffold-api"

// APIOptions api project options
type APIOptions struct {
	Name    string   `json:"name"`     // project name
	Dir     string   `json:"dir"`      // project dir
	ModPath string   `json:"mod_path"` // go module path
	PkgList []string `json:"pkg_list"` // pkg list
}

// InitAPIProject create a new API project
func InitAPIProject(options APIOptions) error {
	color.Blue("git clone %s --recursive", apiProjectRepo)

	//clone repo to local
	if err := util.CloneProjectFromRemote(options.Dir, apiProjectRepo); err != nil {
		return err
	}

	//parse template
	var params map[string]interface{}

	bData, err := json.Marshal(options)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(bData, &params); err != nil {
		return err
	}

	if err = ghost.ParseAll(options.Dir, options.Dir, params); err != nil {
		return err
	}

	return nil
}

// GetPkgFromAPIProject get pkg list from api project
func GetPkgFromAPIProject(options APIOptions) ([]string, error) {
	var pkgList []string

	if !util.IsFolderExist(options.Dir) {
		return nil, errorx.SDKProjectNotExist
	}

	pkgDir := path.Join(options.Dir, "pkg")
	if !util.IsFolderExist(path.Join(options.Dir, "pkg")) {
		return nil, errorx.SDKPkgNotExist
	}

	fileInfoList, err := os.ReadDir(pkgDir)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileInfoList {
		if !fileInfo.IsDir() {
			continue
		}
		pkgList = append(pkgList, fileInfo.Name())
	}

	return pkgList, nil
}

// InstallPkgToApiProject install package to api project
func InstallPkgToApiProject(options APIOptions, pkgList []string) error {
	pkgDir := path.Join(options.Dir, "pkg")

	fileInfoList, err := os.ReadDir(pkgDir)
	if err != nil {
		return err
	}

	for _, fileInfo := range fileInfoList {
		if !fileInfo.IsDir() {
			continue
		}
		if !util.Include(fileInfo.Name(), pkgList) {
			return util.RemoveFolderIfExist(path.Join(pkgDir, fileInfo.Name()))
		}
	}

	return nil
}

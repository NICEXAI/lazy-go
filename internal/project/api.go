package project

import (
	"encoding/json"
	"github.com/NICEXAI/lazy-go/util"
	"github.com/NICEXAI/lazy-template-engine"
	"github.com/fatih/color"
)

const apiProjectRepo = "github.com/NICEXAI/lazy-scaffold-api"

// APIOptions api project options
type APIOptions struct {
	Name    string `json:"name"`     // project name
	Dir     string `json:"dir"`      // project dir
	ModPath string `json:"mod_path"` // go module path
}

// NewAPIProject create a new API project
func NewAPIProject(options APIOptions) error {
	color.Blue("git clone %s --recursive", apiProjectRepo)

	//clone repo to local
	if err := util.CloneProjectFromRemote(options.Dir, apiProjectRepo); err != nil {
		return err
	}

	//parse template
	var params map[string]string

	bData, err := json.Marshal(options)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(bData, &params); err != nil {
		return err
	}

	if err = lazyTemplate.ParseAll(options.Dir, options.Dir, params); err != nil {
		return err
	}

	return nil
}

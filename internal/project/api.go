package project

import (
	"github.com/NICEXAI/lazy-go/util"
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

	return util.CloneProjectFromRemote(options.Dir, apiProjectRepo)
}

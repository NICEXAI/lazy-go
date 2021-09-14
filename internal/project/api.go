package project

import (
	"github.com/fatih/color"
	"lazy-go/util"
)

const apiProjectRepo = "github.com/NICEXAI/quick-comment"

// APIOptions api project options
type APIOptions struct {
	Name    string `json:"name"`     // project name
	Dir     string `json:"dir"`      // project dir
	ModPath string `json:"mod_path"` // go module path
}

// NewAPIProject create a new API project
func NewAPIProject(options APIOptions) error {
	color.Blue("git clone %s --recursive", apiProjectRepo)
	color.Blue(options.Dir)

	return util.CloneProjectFromRemote(options.Dir, apiProjectRepo)
}

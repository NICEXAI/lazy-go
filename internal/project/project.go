package project

import (
	"github.com/fatih/color"
	"lazy-go/util"
)

const (
	//defaultProjectRepo = "https://github.com/NICEXAI/quick-comment"
)

// Options new project options
type Options struct {
	Name string // project name
	Type string // project type, default:normal
}

// New crate a new project
func New(options Options) error {
	var (
		currentPath string
		err         error
	)

	currentPath, err = util.GetCurrentPath()
	if err != nil {
		return err
	}

	color.Green(currentPath)
	return nil
}

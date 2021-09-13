package project

import "github.com/fatih/color"

const apiProjectRepo = "https://github.com/NICEXAI/quick-comment"

// APIOptions api project options
type APIOptions struct {
	Name string `json:"name"` // project name
	Dir  string `json:"dir"`  // project dir
}

// NewAPIProject create a new API project
func NewAPIProject(options APIOptions) error {
	color.Green(options.Dir)
	return nil
}

package project

import (
	"lazy-go/internal/errorx"
	"lazy-go/util"
)

// Options new project options
type Options struct {
	Name     string // project name
	Type     string // project type: API、gRPC
	IsForced bool   // forced new project
}

// New crate a new project
func New(options Options) error {
	projectDir, err := util.GetProjectPath(options.Name)
	if err != nil {
		return err
	}
	if !options.IsForced && util.IsFolderExist(projectDir) {
		return errorx.SDKProjectAlreadyExist
	}

	if options.IsForced && util.IsFolderExist(projectDir) {
		if err = util.RemoveIfExist(projectDir); err != nil {
			return err
		}
	}

	modPath, err := util.GetGoModulePath()
	if err != nil {
		return err
	}

	switch options.Type {
	case "API":
		return NewAPIProject(APIOptions{
			Name:    options.Name,
			Dir:     projectDir,
			ModPath: modPath,
		})
	case "gRPC":
		return NewGRPCProject()
	}
	return nil
}

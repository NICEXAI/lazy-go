package cmd

import (
	"fmt"

	"github.com/NICEXAI/lazy-go/internal/errorx"
	"github.com/NICEXAI/lazy-go/internal/project"
	"github.com/NICEXAI/lazy-go/util"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "project management",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Println(err)
		}
	},
}

var isForced bool

var projectNew = &cobra.Command{
	Use:   "new",
	Short: "crate a new project",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			projectName string
			projectType string
		)

		if len(args) == 0 {
			name, err := util.GetValueFromInput("project name")
			if err != nil {
				color.Red("%v: %s", errorx.SDKProjectNameGetFailed, err.Error())
				return
			}
			projectName = name
		} else {
			projectName = args[0]
		}

		projectDir, err := util.GetProjectPath(projectName)
		if err != nil {
			return
		}

		if !isForced && util.IsFolderExist(projectDir) {
			color.Red("%v", errorx.SDKProjectAlreadyExist)
			return
		}

		if isForced && util.IsFolderExist(projectDir) {
			if err = util.RemoveIfExist(projectDir); err != nil {
				return
			}
		}

		projectType, err = util.GetValueFromSelect("project type", []string{"API", "gRPC"})
		if err != nil {
			color.Red("%v: %s", errorx.SDKProjectTypeGetFailed, err.Error())
			return
		}

		switch projectType {
		case "API":
			apiOptions := project.APIOptions{
				Name: projectName,
				Dir:  projectDir,
			}

			//init api project
			if err = project.InitAPIProject(apiOptions); err != nil {
				color.Red("%v: %s", errorx.SDKProjectInitFailed, err.Error())
				return
			}

			//get pkg list
			var pkgList []string

			pkgList, err = project.GetPkgFromAPIProject(apiOptions)
			if err != nil {
				color.Red("%v: %s", errorx.SDKProjectInitFailed, err.Error())
				return
			}

			//select pkg
			selectPkgList, err := util.GetValueFromMultiSelect("package install:", pkgList)
			if err != nil {
				color.Red("%v: %s", errorx.SDKProjectInitFailed, err.Error())
				return
			}

			if err = project.InstallPkgToApiProject(apiOptions, selectPkgList); err != nil {
				color.Red("%v: %s", errorx.SDKProjectInitFailed, err.Error())
				return
			}

			color.Green("API project created success")
		case "gRPC":
			project.NewGRPCProject()
		}
	},
}

func init() {
	projectNew.Flags().BoolVarP(&isForced, "force", "f", false, "forced new project")

	projectCmd.AddCommand(projectNew)
}

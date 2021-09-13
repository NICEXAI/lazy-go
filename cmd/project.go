package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"lazy-go/internal/errorx"
	"lazy-go/util"

	"github.com/spf13/cobra"

	"lazy-go/internal/project"
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

var projectNew = &cobra.Command{
	Use:   "new",
	Short: "crate a new project",
	Run: func(cmd *cobra.Command, args []string) {
		options := project.Options{}

		if len(args) == 0 {
			name, err := util.GetValueFromInput("project name")
			if err != nil {
				color.Red("%v: %s", errorx.SDKProjectNameGetFailed, err.Error())
				return
			}
			options.Name = name.String()
		}

		projectType, err := util.GetValueFromSelect("project type", []string{"API", "gRPC"})
		if err != nil {
			color.Red("%v: %s", errorx.SDKProjectTypeGetFailed, err.Error())
			return
		}

		options.Type = projectType.String()

		if err := project.New(options); err != nil {
			color.Red("%v: %s", errorx.SDKProjectInitFailed, err.Error())
			return
		}
	},
}

func init() {
	projectCmd.AddCommand(projectNew)
}

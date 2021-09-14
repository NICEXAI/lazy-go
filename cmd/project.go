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
		options := project.Options{}

		if len(args) == 0 {
			name, err := util.GetValueFromInput("project name")
			if err != nil {
				color.Red("%v: %s", errorx.SDKProjectNameGetFailed, err.Error())
				return
			}
			options.Name = name
		} else {
			options.Name = args[0]
		}

		projectType, err := util.GetValueFromSelect("project type", []string{"API", "gRPC"})
		if err != nil {
			color.Red("%v: %s", errorx.SDKProjectTypeGetFailed, err.Error())
			return
		}

		options.Type = projectType
		options.IsForced = isForced

		if err := project.New(options); err != nil {
			color.Red("%v: %s", errorx.SDKProjectInitFailed, err.Error())
			return
		}
	},
}

func init() {
	projectNew.Flags().BoolVarP(&isForced, "force", "f", false, "forced new project")

	projectCmd.AddCommand(projectNew)
}

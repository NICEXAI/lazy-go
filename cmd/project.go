package cmd

import (
	"fmt"
	"github.com/fatih/color"
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

var (
	projectType string
)

var projectNew = &cobra.Command{
	Use:   "new",
	Short: "crate a new project",
	Run: func(cmd *cobra.Command, args []string) {
		options := project.Options{
			Type: projectType,
		}

		if len(args) == 0 {
			name, err := util.GetStringFromInput("project name")
			if err != nil {
				color.Red("project name get failed: %s", err.Error())
				return
			}
			color.Blue(name)
		}

		if err := project.New(options); err != nil {
			color.Red("project init failed: %s", err.Error())
			return
		}
		fmt.Println(args, projectType)
	},
}

func init() {
	projectNew.Flags().StringVarP(&projectType, "type", "t", "normal", "project type")

	projectCmd.AddCommand(projectNew)
}

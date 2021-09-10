package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var project = &cobra.Command{
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
		fmt.Println("generate project", args)
	},
}

func init() {
	project.AddCommand(projectNew)
}

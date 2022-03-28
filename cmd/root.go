package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

const version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "lazy",
	Short: "lazy-go is an easy-to-use WEB framework.",
	Long:  "lazy-go is an easy-to-use WEB framework, it can quickly create projects and generate code.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Println(err)
		}
	},
	Version: fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH),
}

func init() {
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(configCmd)
}

// Execute entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

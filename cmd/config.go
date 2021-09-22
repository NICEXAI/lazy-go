package cmd

import (
	"fmt"
	"github.com/NICEXAI/lazy-go/internal/config"
	"github.com/NICEXAI/lazy-go/internal/errorx"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config management",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Println(err)
		}
	},
}

var configSync = &cobra.Command{
	Use:   "sync",
	Short: "config auto sync",
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.Sync(); err != nil {
			color.Red("%v: %s", errorx.SDKConfigSyncFailed, err.Error())
			return
		}
		color.Green("config auto_sync success")
	},
}

func init() {
	configCmd.AddCommand(configSync)
}

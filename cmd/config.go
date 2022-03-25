package cmd

import (
	"fmt"
	"github.com/NICEXAI/lazy-go/internal/config"
	"github.com/NICEXAI/lazy-go/internal/errorx"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	inputFilePath  string
	outputFilePath string
	moduleName     string
	fieldTagName   string
	watchMode      bool
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config management",
	Long:  "Use the command to convert arbitrary formats to Go Struct (including json, toml, yaml, etc.)",
	Run: func(cmd *cobra.Command, args []string) {
		if inputFilePath == "" || outputFilePath == "" {
			if err := cmd.Help(); err != nil {
				color.Red(err.Error())
			}
			return
		}

		if err := config.Convert(inputFilePath, outputFilePath, moduleName, fieldTagName); err != nil {
			color.Red(err.Error())
			return
		}

		color.Green("file convert success")

		if watchMode {
			fsTask, err := config.Watch(inputFilePath, outputFilePath, moduleName, fieldTagName)
			if err != nil {
				color.Red("%v: %s", errorx.SDKConfigGenerateFailed, err.Error())
				return
			}

			color.Blue("file listening...")
			fsTask.Wait()
		}
	},
	Version: fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH),
}

func init() {
	configCmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "source file path")
	configCmd.Flags().StringVarP(&outputFilePath, "output", "o", "", "target file path")
	configCmd.Flags().StringVarP(&moduleName, "module", "m", "", "module name of the target file, default: target file name")
	configCmd.Flags().StringVarP(&fieldTagName, "tag", "t", "json", "set filed tag name, default: json")
	configCmd.Flags().BoolVarP(&watchMode, "watch", "w", false, "listening file changes and auto convert content to Go Struct")
}

package util

import (
	"errors"
	"github.com/manifoldco/promptui"
)

// GetStringFromInput get string from input
func GetStringFromInput(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			if len(s) == 0 {
				return errors.New(" project name cannot be empty")
			}
			return nil
		},
	}

	return prompt.Run()
}

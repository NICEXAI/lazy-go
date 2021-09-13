package util

import (
	"errors"
	"github.com/manifoldco/promptui"
)

// Prompt prompt
type Prompt struct {
	originValue string // 原始值
}

func (p Prompt) String() string {
	return p.originValue
}

// GetValueFromInput get value from input
func GetValueFromInput(label string) (Prompt, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			if len(s) == 0 {
				return errors.New(" project name cannot be empty")
			}
			return nil
		},
	}

	origin, err := prompt.Run()
	if err != nil {
		return Prompt{}, err
	}

	return Prompt{originValue: origin}, nil
}

// GetValueFromSelect get value from select
func GetValueFromSelect(label string, options []string) (Prompt, error) {
	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	_, origin, err := prompt.Run()
	if err != nil {
		return Prompt{}, err
	}

	return Prompt{originValue: origin}, nil
}

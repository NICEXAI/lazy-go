package util

import (
	"github.com/AlecAivazis/survey/v2"
)

// GetValueFromInput get value from input
func GetValueFromInput(label string) (string, error) {
	origin := ""

	prompt := &survey.Input{
		Message: label + ": ",
	}
	if err := survey.AskOne(prompt, &origin); err != nil {
		return "", err
	}

	return origin, nil
}

// GetValueFromSelect get value from select
func GetValueFromSelect(label string, options []string) (string, error) {
	origin := ""

	prompt := &survey.Select{
		Message: label + ": ",
		Options: options,
	}

	if err := survey.AskOne(prompt, &origin); err != nil {
		return "", err
	}

	return origin, nil
}

// GetValueFromMultiSelect get value from multi select
func GetValueFromMultiSelect(label string, options []string) ([]string, error) {
	var origin []string

	prompt := &survey.Select{
		Message: label + ": ",
		Options: options,
	}

	if err := survey.AskOne(prompt, &origin); err != nil {
		return nil, err
	}

	return origin, nil
}

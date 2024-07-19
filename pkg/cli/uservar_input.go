package cli

import (
	"fmt"

	"github.com/atomicptr/tmplr/pkg/tmpl"
	"github.com/charmbracelet/huh"
)

func inputUserVar(userVar tmpl.UserVar) (string, error) {
	var value string

	prompt := userVar.Prompt

	if prompt == "" {
		prompt = fmt.Sprintf("Enter a value for var '%s'", userVar.Name)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(prompt).
				Prompt("> ").
				Value(&value),
		),
	)

	err := form.Run()
	if err != nil {
		return "", err
	}

	if value == "" {
		value = userVar.Default
	}

	return value, nil
}

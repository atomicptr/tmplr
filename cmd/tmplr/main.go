package main

import (
	"fmt"

	"github.com/atomicptr/tmplr/pkg/cli"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	err := cli.Run()
	if err != nil {
		style := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF0000"))
		fmt.Println(style.Render("[ERROR]"), err)
	}
}

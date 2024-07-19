package cli

import (
	"fmt"

	"github.com/atomicptr/tmplr/pkg/tmpl"
	"github.com/charmbracelet/huh"
)

func selectTemplate(templates []*tmpl.Template) (*tmpl.Template, error) {
	var selected *tmpl.Template

	var options []huh.Option[*tmpl.Template]

	for _, templ := range templates {
		name := templ.Name

		if name == "" {
			name = templ.TemplateName
		}

		if templ.Name != "" && templ.TemplateName != "" {
			name = fmt.Sprintf("%s - File: %s", templ.Name, templ.TemplateName)
		}

		options = append(options, huh.NewOption(name, templ))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[*tmpl.Template]().Options(options...).
				Title("Select Template").
				Description("This filename matches multiple templates, please select one").
				Value(&selected),
		),
	)

	err := form.Run()
	if err != nil {
		return nil, err
	}

	return selected, nil
}

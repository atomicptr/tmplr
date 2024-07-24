package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/atomicptr/tmplr/pkg/fs"
	"github.com/atomicptr/tmplr/pkg/meta"
	"github.com/atomicptr/tmplr/pkg/tmpl"
	"github.com/charmbracelet/huh"
	flag "github.com/spf13/pflag"
)

func Run() error {
	showVersion := false
	showTemplateDir := false

	// parse flags
	flag.BoolVarP(&showTemplateDir, "template-dir", "", false, "Print template directory location")
	flag.BoolVarP(&showVersion, "version", "", false, "Print tmplr version")
	flag.Parse()

	if showVersion {
		fmt.Println(meta.VersionString())
		return nil
	}

	if showTemplateDir {
		templateDir, err := fs.TemplateDir()
		if err != nil {
			return err
		}

		fmt.Println(templateDir)

		return nil
	}

	if flag.NArg() == 0 {
		flag.Usage()
		fmt.Println("\nPlease provide files to create.")
		os.Exit(1)
	}

	templateFiles, err := fs.ListTemplateFiles()
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, arg := range flag.Args() {
		p, err := filepath.Abs(arg)
		if err != nil {
			return err
		}

		matchingTemplates := tmpl.FindMatchingTemplates(p, templateFiles)

		if len(matchingTemplates) == 0 {
			f, err := fs.OpenFile(p)
			if err != nil {
				return err
			}
			defer (func() {
				err := f.Close()
				if err != nil {
					panic(err)
				}
			})()
			continue
		}

		var templates []*tmpl.Template

		for _, templateFile := range matchingTemplates {
			template, err := tmpl.LoadTemplate(p, templateFile)
			if err != nil {
				return err
			}

			templates = append(templates, template)
		}

		var selected *tmpl.Template = nil

		if len(templates) == 1 {
			selected = templates[0]
		}

		if selected == nil {
			tpl, err := selectTemplate(templates)

			if err == huh.ErrUserAborted {
				return nil
			}

			if err != nil {
				return err
			}

			// user picked quit
			if tpl == nil {
				return nil
			}

			selected = tpl
		}

		for _, userVar := range selected.UserVars {
			val, err := inputUserVar(userVar)
			if err != nil {
				return err
			}

			selected.Data[userVar.Name] = val
		}

		// add some vars that are always present
		selected.Data["_cwd"] = cwd
		selected.Data["_path"] = p
		selected.Data["_dirname"] = filepath.Base(filepath.Dir(p))
		selected.Data["_filename"] = filepath.Base(p)

		fmt.Println(selected.Data)

		data, err := selected.Render()
		if err != nil {
			return err
		}

		f, err := fs.OpenFile(p)
		if err != nil {
			return fmt.Errorf("could not create file %s: %w", p, err)
		}
		defer (func() {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		})()

		_, err = f.WriteString(data)
		if err != nil {
			return err
		}
	}

	return nil
}

package tmpl

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/adrg/frontmatter"
)

type UserVar struct {
	Name    string `yaml:"name"`
	Prompt  string `yaml:"prompt"`
	Default string `yaml:"default"`
}

type Template struct {
	Name         string    `yaml:"name"`
	UserVars     []UserVar `yaml:"vars"`
	TemplateName string
	Content      string
	Data         map[string]string
}

func (t *Template) Render() (string, error) {
	tmpl, err := template.New(t.TemplateName).Parse(t.Content)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, t.Data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func LoadTemplate(filePath string, templatePath string) (*Template, error) {
	f, err := os.Open(templatePath)
	if err != nil {
		return nil, err
	}
	defer (func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	})()
	return ParseTemplate(filepath.Base(filePath), filepath.Base(templatePath), f)
}

func ParseTemplate(fileName, templateName string, r io.Reader) (*Template, error) {
	var template Template
	content, err := frontmatter.Parse(r, &template)
	if err != nil {
		return nil, err
	}

	template.TemplateName = templateName

	if template.Data == nil {
		template.Data = make(map[string]string)
	}

	nameVars := ParseFilename(fileName, templateName)

	for key, value := range nameVars {
		template.Data[key] = value
	}

	template.Content = string(content)
	return &template, nil
}

func FindMatchingTemplates(filePath string, templateFiles []string) []string {
	var matching []string

	fname := filepath.Base(filePath)

	for _, templateFile := range templateFiles {
		tname := filepath.Base(templateFile)

		// name is equals? Match
		if fname == tname {
			matching = append(matching, templateFile)
			continue
		}

		// filename does not contain vars? Skip
		if !strings.Contains(tname, "[") {
			continue
		}

		if !MatchesFilename(fname, tname) {
			continue
		}

		matching = append(matching, templateFile)
	}

	return matching
}

func MatchesFilename(fileName, templateName string) bool {
	pattern := fmt.Sprintf("^%s$", regexp.MustCompile(`\[([^\]]+)\]`).ReplaceAllString(templateName, `([^/]+)`))

	matched, err := regexp.MatchString(pattern, fileName)
	if err != nil {
		return false
	}

	return matched
}

func ParseFilename(fileName, templateName string) map[string]string {
	if !MatchesFilename(fileName, templateName) {
		return nil
	}

	result := make(map[string]string)

	varNamesPattern := regexp.MustCompile(`\[([^\]]+)\]`)
	varNames := varNamesPattern.FindAllStringSubmatch(templateName, -1)

	valuePattern := fmt.Sprintf("^%s$", varNamesPattern.ReplaceAllString(templateName, `([^/]+)`))

	r := regexp.MustCompile(valuePattern)
	matches := r.FindStringSubmatch(fileName)

	if len(matches) > 0 {
		for i, match := range matches[1:] {
			result[varNames[i][1]] = match
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result
}

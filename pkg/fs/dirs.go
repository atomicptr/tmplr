package fs

import (
	"os"
	"path/filepath"

	"github.com/atomicptr/tplr/pkg/meta"
)

func TemplateDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	templateDir := filepath.Join(configDir, meta.ConfigDirName, "templates")

	err = Mkdir(templateDir)
	if err != nil {
		return "", err
	}

	return templateDir, nil
}

func ListTemplateFiles() ([]string, error) {
	templateDir, err := TemplateDir()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(templateDir)
	if err != nil {
		return nil, err
	}

	var files []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		files = append(files, filepath.Join(templateDir, entry.Name()))
	}

	return files, nil
}

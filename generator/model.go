package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/awgst/goes"
)

type Model struct {
}

func (m *Model) Make(name string, dir string, packages string, args ...string) error {
	filename := fmt.Sprintf("%v.%v", goes.SnakeCase(name), "go")

	os.Mkdir(dir, 0755)
	path := filepath.Join(dir, filename)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return fmt.Errorf("File model is already exists : %v", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create model file: %w", err)
	}
	defer f.Close()

	vars := tmplVars{
		Version:   "1",
		CamelName: goes.CamelCase(name),
	}

	var t = ""

	if len(args) > 0 {
		t = args[0]
	}

	tmpl := m.getTemplate(name, packages, t)
	if err := tmpl.Execute(f, vars); err != nil {
		return fmt.Errorf("Failed to execute tmpl: %w", err)
	}

	log.Printf("Created new file: %s\n", f.Name())
	return nil
}

func (m *Model) getTemplate(name string, packages string, tmpl string) *template.Template {
	var parsed = fmt.Sprintf(`package %v

type %v struct {
}
	`, packages, name)
	if tmpl != "" {
		parsed = tmpl
	}
	return template.Must(template.New("goes.model").Parse(parsed))
}

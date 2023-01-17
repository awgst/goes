package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/awgst/goes"
)

type Repository struct {
}

func (c *Repository) Make(name string, dir string, packages string) error {
	filename := fmt.Sprintf("%v.%v", goes.SnakeCase(name), "go")

	os.Mkdir(dir, 0755)
	path := filepath.Join(dir, filename)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return fmt.Errorf("File repository is already exists : %v", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create repository file: %w", err)
	}
	defer f.Close()

	vars := tmplVars{
		Version:   "1",
		CamelName: goes.CamelCase(name),
	}

	tmpl := c.getTemplate(name, packages)
	if err := tmpl.Execute(f, vars); err != nil {
		return fmt.Errorf("Failed to execute tmpl: %w", err)
	}

	log.Printf("Created new file: %s\n", f.Name())
	return nil
}

func (c *Repository) getTemplate(name string, packages string) *template.Template {
	structName := goes.InitialLowerCase(name)
	var parsed = fmt.Sprintf(`package %v

import (
	"gorm.io/gorm"
)

type %vInterface interface {
}

type %v struct {
	db *gorm.DB
}

func New%v() %vInterface {
	return &%v{}
}

	`, packages, name, structName, name, name, structName)
	return template.Must(template.New("goes.repository").Parse(parsed))
}

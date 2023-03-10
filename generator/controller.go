package generator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/awgst/goes"
)

type Controller struct {
}

func (c *Controller) Make(name string, dir string, packages string, args ...string) error {
	filename := fmt.Sprintf("%v.%v", goes.SnakeCase(name), "go")

	os.Mkdir(dir, 0755)
	path := filepath.Join(dir, filename)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return fmt.Errorf("File controller is already exists : %v", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create controller file: %w", err)
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

	tmpl := c.getTemplate(name, packages, t)
	if err := tmpl.Execute(f, vars); err != nil {
		return fmt.Errorf("Failed to execute tmpl: %w", err)
	}

	log.Printf("Created new file: %s\n", f.Name())
	return nil
}

func (c *Controller) getTemplate(name string, packages string, tmpl string) *template.Template {
	structName := goes.InitialLowerCase(name)
	var parsed = fmt.Sprintf(`package %v

import (
	"github.com/gin-gonic/gin"
)

type %vInterface interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
}

type %v struct {
}

func New%v() %vInterface {
	return &%v{}
}

// Get resources list
func (ctr *%v) Index(c *gin.Context) {

}

// Get resource by id
func (ctr *%v) Show(c *gin.Context) {
	id := c.Param("id")
}

// Create resource
func (ctr *%v) Create(c *gin.Context) {

}

// Update resource
func (ctr *%v) Update(c *gin.Context) {
	id := c.Param("id")
}

// Delete resource
func (ctr *%v) Delete(c *gin.Context) {
	id := c.Param("id")
}

	`, packages, name, structName, name, name, structName, structName, structName, structName, structName, structName)
	if tmpl != "" {
		parsed = tmpl
	}
	return template.Must(template.New("goes.controller").Parse(parsed))
}

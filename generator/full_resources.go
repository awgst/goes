package generator

import (
	"fmt"

	"github.com/awgst/goes"
)

type FullResource struct {
	model      Model
	repository Repository
	service    Service
	controller Controller
	request    Request
	response   Response
}

func (fr *FullResource) Make(name string, dir string, packages string) error {

	modelName := name + "Model"
	modelPackages := "model"
	modelDir := dir + "/model"
	var modelTemplate = fmt.Sprintf(`package %v

type %v struct {
}
	`, modelPackages, modelName)
	err := fr.model.Make(modelName, modelDir, modelPackages, modelTemplate)

	repositoryName := name + "Repository"
	repositoryPackages := "repository"
	repositoryDir := dir + "/repository"
	structName := goes.InitialLowerCase(repositoryName)
	var repositoryTemplate = fmt.Sprintf(`package %v

import (
	"gorm.io/gorm"
)

type %vInterface interface {
}

type %v struct {
	db *gorm.DB
}

func New%v(db *gorm.DB) %vInterface {
	return &%v{db:db}
}

	`, repositoryPackages, repositoryName, structName, repositoryName, repositoryName, structName)
	err = fr.repository.Make(repositoryName, repositoryDir, repositoryPackages, repositoryTemplate)
	serviceName := name + "Service"
	servicePackage := "service"
	serviceDir := dir + "/service"
	structName = goes.InitialLowerCase(serviceName)

	var serviceTemplate = fmt.Sprintf(`package %v

type %vInterface interface {
}

type %v struct {
	repository repository.%vInterface
}

func New%v(repository repository.%vInterface) %vInterface {
	return &%v{repository: repository}
}

	`, servicePackage, serviceName, structName, repositoryName, serviceName, repositoryName, serviceName, structName)

	err = fr.service.Make(serviceName, serviceDir, servicePackage, serviceTemplate)

	controllerName := name + "Controller"
	controllerPackage := "controller"
	controllerDir := dir + "/controller"
	structName = goes.InitialLowerCase(controllerName)

	var controllerTemplate = fmt.Sprintf(`package %v

import (
	"github.com/gin-gonic/gin"
)

type %vInterface interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type %v struct {
	service service.%vInterface
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
	
		`, controllerPackage, controllerName, structName, serviceName, controllerName, controllerName, structName, structName, structName, structName, structName, structName)

	err = fr.controller.Make(controllerName, controllerDir, controllerPackage, controllerTemplate)

	return err
}

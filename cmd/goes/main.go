package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/awgst/goes/generator"
)

var (
	flags   = flag.NewFlagSet("goes", flag.ExitOnError)
	help    = flags.Bool("help", false, "print help")
	version = flags.Bool("version", false, "print version")
)

var (
	defaultPath = "./"
	goesVersion = ""
)

const (
	commandList = `
goes COMMAND PACKAGE
Commands:
  create:model NAME [PACKAGE]		Create model
  create:controller NAME [PACKAGE]	Create controller
  create:repository NAME [PACKAGE]	Create repository
  create:resources NAME [PACKAGE]	Create full resources
  create:response NAME [PACKAGE]	Create response
  create:request NAME [PACKAGE]		Create request
  create:service NAME [PACKAGE]		Create service
Options:
  -help		show help
  -version	show version
	`
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])
	if *version {
		if buildInfo, ok := debug.ReadBuildInfo(); ok && buildInfo != nil && goesVersion == "" {
			goesVersion = buildInfo.Main.Version
		}
		fmt.Printf("goes version:%s\n", goesVersion)
		return
	}

	if *help {
		flags.Usage()
		return
	}

	args := flags.Args()

	if len(args) < 2 {
		flags.Usage()
		return
	}

	commandArgs := strings.Split(args[0], ":")
	command := commandArgs[0]
	var generated = ""
	if len(commandArgs) > 1 {
		generated = commandArgs[1]
	}
	var packages = generated
	if len(args) >= 3 {
		packages = args[2]
	}
	generate(command, generated, args[1], packages)
}

func usage() {
	fmt.Println(commandList)
}

func generate(command string, generated string, name string, packages string) {
	switch command {
	case "create":
		if !isAvailable(generated) {
			fmt.Println(generated + " is not available. Please available command using goes -help")
			return
		}
		err := getGenerator(generated, name, packages)
		if err != nil {
			fmt.Println("Error on creating "+generated+" : ", err)
		}
		return
	default:
		fmt.Println(command + " is not available. Please available command using goes -help")
		return
	}
}

func isAvailable(generated string) bool {
	available := []string{
		"model",
		"controller",
		"repository",
		"resources",
		"response",
		"request",
		"service",
	}

	for _, a := range available {
		if generated == a {
			return true
		}
	}

	return false
}

func getGenerator(generated string, name string, packages string) error {
	var gen generator.GeneratorInterface
	switch generated {
	case "model":
		var modelGenerator generator.Model
		gen = &modelGenerator
	case "controller":
		var controllerGenerator generator.Controller
		gen = &controllerGenerator
	case "repository":
		var repositoryGenerator generator.Repository
		gen = &repositoryGenerator
	}

	dir := defaultPath + packages

	err := gen.Make(name, dir, packages)
	return err
}

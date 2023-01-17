package generator

type GeneratorInterface interface {
	Make(name string, dir string, packages string, args ...string) error
}

type tmplVars struct {
	Version   string
	CamelName string
}

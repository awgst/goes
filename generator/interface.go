package generator

type GeneratorInterface interface {
	Make(name string, dir string, packages string) error
}

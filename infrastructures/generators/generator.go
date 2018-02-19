package generators

// Generator is the interface used to generate passwords
type Generator interface {
	Generate(plain string) []string
}

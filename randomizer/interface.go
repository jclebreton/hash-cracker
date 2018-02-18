package randomizer

// Randomizeris the interface used to randomize plan passwords
type Randomizer interface {
	Randomize(plain string) []string
}

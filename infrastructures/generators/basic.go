package generators

import (
	"fmt"
	"strings"
)

// Basic is a basic implementation of a password generator
type Basic struct {
}

func (n *Basic) suffixWithNumbers(plain string, min, max int) []string {
	var result []string
	current := plain
	for i := min; i <= max; i++ {
		current = current + fmt.Sprintf("%d", i)
		result = append(result, current)
	}
	return result
}

func (n *Basic) title(plain string) string {
	return strings.Title(plain)
}

// Generate returns generated plain passwords
func (n *Basic) Generate(plain string) []string {
	var result []string
	result = append(result, plain)
	result = append(result, n.suffixWithNumbers(plain, 1, 9)...)

	for _, v := range result {
		result = append(result, n.title(v))
	}

	return result
}

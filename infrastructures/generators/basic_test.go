package generators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addSuffixNumbers_success(t *testing.T) {
	expected := []string{
		"qwerty1",
		"qwerty12",
		"qwerty123",
		"qwerty1234",
		"qwerty12345",
		"qwerty123456",
		"qwerty1234567",
		"qwerty12345678",
		"qwerty123456789",
	}
	numbers := &Basic{}
	result := numbers.suffixWithNumbers("qwerty", 1, 9)
	for k := range expected {
		assert.Equal(t, expected[k], result[k])
	}
}

func Test_title_success(t *testing.T) {
	numbers := &Basic{}
	assert.Equal(t, "Qwerty", numbers.title("qwerty"))
}

func TestRandomize_success(t *testing.T) {
	numbers := &Basic{}
	result := numbers.Generate("qwerty")
	assert.Equal(t, 20, len(result))
	assert.Equal(t, "qwerty", result[0])
	assert.Equal(t, "qwerty1", result[1])
	assert.Equal(t, "qwerty12", result[2])
	assert.Equal(t, "Qwerty", result[10])
	assert.Equal(t, "Qwerty1", result[11])
	assert.Equal(t, "Qwerty12", result[12])
}

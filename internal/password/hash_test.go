package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator_Generate_Compare(t *testing.T) {
	const password string = "secret"
	g := generator{}
	hash, err := g.Generate(password)
	assert.Nil(t, err)
	assert.Nil(t, g.Compare(hash, password))
}

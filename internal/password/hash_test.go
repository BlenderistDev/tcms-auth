package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate_Compare(t *testing.T) {
	const password string = "secret"
	hash, err := Generate(password)
	assert.Nil(t, err)
	assert.Nil(t, Compare(hash, password))
}

package session

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSessionToken(t *testing.T) {
	s := GenerateSessionToken()
	assert.Len(t, s, length)
}

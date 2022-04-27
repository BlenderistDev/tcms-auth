package password

import "golang.org/x/crypto/bcrypt"

// Generator generates and compare passwords
type Generator interface {
	Generate(s string) (string, error)
	Compare(hash string, s string) error
}

type generator struct{}

// NewGenerator creates new password generator instance
func NewGenerator() Generator {
	return &generator{}
}

//Generate a salted hash for the input string
func (g *generator) Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

//Compare string to generated hash
func (g *generator) Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

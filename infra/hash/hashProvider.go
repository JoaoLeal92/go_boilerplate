package hash

import "golang.org/x/crypto/bcrypt"

// Provider hash provider struct
type Provider struct{}

// NewProvider instantiates hash provider
func NewProvider() *Provider {
	return &Provider{}
}

// GenerateHash generates hashed password
func (p *Provider) GenerateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashedPassword), err
}

// CompareHashAndPassword compares user password with provided password
func (p *Provider) CompareHashAndPassword(userPassword string, inputPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(inputPassword))
}

package hash

import (
	"errors"
)

// FakeHashProvider hash provider struct
type FakeHashProvider struct{}

// NewFakeProvider instantiates hash provider
func NewFakeProvider() *FakeHashProvider {
	return &FakeHashProvider{}
}

// GenerateHash generates hashed password
func (p *FakeHashProvider) GenerateHash(password string) (string, error) {
	if password == "" {
		return "", errors.New("Password required")
	}
	return password, nil
}

// CompareHashAndPassword compares user password with provided password
func (p *FakeHashProvider) CompareHashAndPassword(userPassword string, inputPassword string) error {
	if userPassword == inputPassword {
		return nil
	}
	return errors.New("Passwords don't match")
}

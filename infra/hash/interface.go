package hash

// HashProvider interface for hash provider
type HashProvider interface {
	GenerateHash(password string) (string, error)
	CompareHashAndPassword(userPassword string, inputPassword string) error
}

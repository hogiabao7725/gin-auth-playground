package domain

type PasswordHasher interface {
	Hash(plain string) (string, error)
	Compare(hash, plain string) error
}

type IdentifierGenerator interface {
	Generate() string
}

package domain

import "strings"

type Name struct {
	value string
}

func NewName(raw string) (Name, error) {
	normalizedName, err := validateAndNormalizeName(raw)
	if err != nil {
		return Name{}, err
	}
	return Name{value: normalizedName}, nil
}

func ReconstituteName(raw string) Name {
	return Name{value: raw}
}

func (n Name) String() string {
	return n.value
}

func (n Name) Equal(other Name) bool {
	return n == other
}

func validateAndNormalizeName(raw string) (string, error) {
	normalized := strings.Join(strings.Fields(raw), " ")
	if normalized == "" {
		return "", ErrEmptyName
	}
	return normalized, nil
}

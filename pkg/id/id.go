package id

import "github.com/google/uuid"

// ID is an unique identifier
type ID string

// New generates a new id
func New() ID {
	return ID(uuid.Must(uuid.NewRandom()).String())
}

// Validate returns error if the id is not valid
func (id ID) Validate() error {
	_, err := uuid.Parse(string(id))
	return err
}

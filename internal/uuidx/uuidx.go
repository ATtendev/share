package uuidx

import "github.com/google/uuid"

func MustParsePointer(in string) *uuid.UUID {
	id := uuid.MustParse(in)
	return &id
}

func MustParse(in string) uuid.UUID {
	return uuid.MustParse(in)
}

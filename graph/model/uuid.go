package model

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

// MarshalUUID allows uuid to be marshalled by graphql
func MarshalUUID(id uuid.UUID) graphql.Marshaler {
	return graphql.MarshalString(id.String())
}

// UnmarshalUUID allows uuid to be unmarshalled by graphql
func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	idAsString, ok := v.(string)
	if !ok {
		return uuid.Nil, errors.New("id should be a valid UUID")
	}
	return uuid.Parse(idAsString)
}

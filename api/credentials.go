//go:generate protoc --gogofaster_out=plugins=grpc:. --micro_out=. --proto_path=$GOPATH/src:$GOPATH/pkg/mod:. credentials.proto

// Package credentials defines the protocol buffers API for the credentials
// service.
package credentials

import (
	"io"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalGQL marshals a CredentialType into a GraphQL string.
func (ct CredentialType) MarshalGQL(w io.Writer) {
	graphql.MarshalString(CredentialType_name[int32(ct)]).MarshalGQL(w)
}

// UnmarshalGQL unmarshals a string into a CredentialType.
func (ct *CredentialType) UnmarshalGQL(v interface{}) error {
	s, err := graphql.UnmarshalString(v)
	*ct = CredentialType(CredentialType_value[s])

	return err
}

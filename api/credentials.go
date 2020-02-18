//go:generate protoc --gogofaster_out=plugins=grpc:. --micro_out=. --proto_path=$GOPATH/src:$GOPATH/pkg/mod:. credentials.proto

package credentials

import (
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func (ct CredentialType) MarshalGQL(w io.Writer) {
	graphql.MarshalString(CredentialType_name[int32(ct)]).MarshalGQL(w)
}

func (ct *CredentialType) UnmarshalGQL(v interface{}) error {
	s, err := graphql.UnmarshalString(v)
	*ct = CredentialType(CredentialType_value[s])
	return err
}

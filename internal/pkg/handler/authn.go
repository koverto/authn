package handler

import (
	"context"
	"fmt"

	authn "github.com/koverto/authn/api"

	"github.com/koverto/mongo"
	mmongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type Authn struct {
	client mongo.Client
}

func New(conf *Config) (*Authn, error) {
	client, err := mongo.NewClient(conf.MongoUrl, conf.Name)
	if err != nil {
		return nil, err
	}

	var index mmongo.IndexModel
	index.Keys = bsonx.Doc{{Key: "userid", Value: bsonx.Int32(1)}}

	client.DefineIndexes(mongo.NewIndexSet("credentials", index))

	return &Authn{client}, nil
}

func (a *Authn) Create(ctx context.Context, in *authn.Credential, out *authn.CredentialResponse) error {
	return fmt.Errorf("not yet implemented") // TODO
}

func (a *Authn) Validate(ctx context.Context, in *authn.Credential, out *authn.CredentialResponse) error {
	return fmt.Errorf("not yet implemented") // TODO
}

func (a *Authn) Update(ctx context.Context, in *authn.CredentialUpdate, out *authn.CredentialResponse) error {
	return fmt.Errorf("not yet implemented") // TODO
}

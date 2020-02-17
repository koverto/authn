package handler

import (
	"context"
	"fmt"

	authn "github.com/koverto/authn/api"
	"github.com/koverto/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/koverto/mongo"
	"go.mongodb.org/mongo-driver/bson"
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
	in.Id = uuid.New()
	out.Success = false

	switch in.GetCredentialType() {
	case authn.CredentialType_PASSWORD:
		password, err := bcrypt.GenerateFromPassword(in.Credential, bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		in.Credential = password
	default:
		return fmt.Errorf("invalid credential type")
	}

	ins, err := bson.Marshal(in)
	if err != nil {
		return err
	}

	collection := a.client.Collection("credentials")
	_, err = collection.InsertOne(ctx, ins)
	if err == nil {
		out.Success = true
	}

	return err
}

func (a *Authn) Validate(ctx context.Context, in *authn.Credential, out *authn.CredentialResponse) error {
	out.Success = false

	filter := bson.D{{Key: "userid", Value: in.UserID}, {Key: "credentialtype", Value: in.CredentialType}}
	collection := a.client.Collection(("credentials"))

	result := &authn.Credential{}
	if err := collection.FindOne(ctx, filter).Decode(result); err != nil {
		return err
	}

	switch in.CredentialType {
	case authn.CredentialType_PASSWORD:
		if err := bcrypt.CompareHashAndPassword(result.GetCredential(), in.GetCredential()); err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid credential type")
	}

	out.Success = true
	return nil
}

func (a *Authn) Update(ctx context.Context, in *authn.CredentialUpdate, out *authn.CredentialResponse) error {
	return fmt.Errorf("not yet implemented") // TODO
}

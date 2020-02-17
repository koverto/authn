package handler

import (
	"context"

	authn "github.com/koverto/authn/api"

	"github.com/koverto/errors"
	"github.com/koverto/mongo"
	"github.com/koverto/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mmongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"golang.org/x/crypto/bcrypt"
)

const AUTHN_COLLECTION_CREDENTIALS = "credentials"

type Authn struct {
	*Config
	client mongo.Client
}

func New(conf *Config) (*Authn, error) {
	client, err := mongo.NewClient(conf.MongoUrl, conf.Name)
	if err != nil {
		return nil, err
	}

	var index mmongo.IndexModel
	index.Keys = bsonx.Doc{{Key: "userid", Value: bsonx.Int32(1)}}

	client.DefineIndexes(mongo.NewIndexSet(AUTHN_COLLECTION_CREDENTIALS, index))

	return &Authn{conf, client}, nil
}

func (a *Authn) Create(ctx context.Context, in *authn.Credential, out *authn.CredentialResponse) error {
	in.Id = uuid.New()
	out.Success = false

	switch ct := in.GetCredentialType(); ct {
	case authn.CredentialType_PASSWORD:
		password, err := bcrypt.GenerateFromPassword(in.Credential, bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		in.Credential = password
	default:
		return a.InvalidCredentialType(ct)
	}

	ins, err := bson.Marshal(in)
	if err != nil {
		return err
	}

	collection := a.client.Collection(AUTHN_COLLECTION_CREDENTIALS)
	_, err = collection.InsertOne(ctx, ins)
	if err == nil {
		out.Success = true
	}

	return err
}

func (a *Authn) Validate(ctx context.Context, in *authn.Credential, out *authn.CredentialResponse) error {
	out.Success = false

	filter := bson.D{{Key: "userid", Value: in.UserID}, {Key: "credentialtype", Value: in.CredentialType}}
	collection := a.client.Collection((AUTHN_COLLECTION_CREDENTIALS))

	result := &authn.Credential{}
	if err := collection.FindOne(ctx, filter).Decode(result); err != nil {
		if err == mmongo.ErrNoDocuments {
			return a.InvalidCredential()
		}

		return err
	}

	switch ct := in.GetCredentialType(); ct {
	case authn.CredentialType_PASSWORD:
		if err := bcrypt.CompareHashAndPassword(result.GetCredential(), in.GetCredential()); err != nil {
			return a.InvalidCredential()
		}

		out.Success = true
	default:
		return a.InvalidCredentialType(ct)
	}

	return nil
}

func (a *Authn) Update(ctx context.Context, in *authn.CredentialUpdate, out *authn.CredentialResponse) error {
	return errors.NotImplemented(a.ID())
}

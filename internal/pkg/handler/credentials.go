package handler

import (
	"context"

	credentials "github.com/koverto/credentials/api"

	"github.com/koverto/errors"
	"github.com/koverto/mongo"
	"github.com/koverto/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mmongo "go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const CREDENTIALS_COLLECTION = "credentials"

type Credentials struct {
	*Config
	client mongo.Client
}

func New(conf *Config) (*Credentials, error) {
	client, err := mongo.NewClient(conf.MongoUrl, conf.Name)
	if err != nil {
		return nil, err
	}

	var uidIndex mmongo.IndexModel
	uidIndex.Keys = bson.M{"userid": 1}

	var credIndex mmongo.IndexModel
	credIndex.Keys = bson.M{"credential": "hashed"}

	client.DefineIndexes(mongo.NewIndexSet(CREDENTIALS_COLLECTION, uidIndex, credIndex))
	if err := client.Connect(); err != nil {
		return nil, err
	}

	return &Credentials{conf, client}, nil
}

func (a *Credentials) Create(ctx context.Context, in *credentials.Credential, out *credentials.CredentialResponse) error {
	in.Id = uuid.New()
	out.Success = false

	switch ct := in.GetCredentialType(); ct {
	case credentials.CredentialType_PASSWORD:
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

	collection := a.client.Collection(CREDENTIALS_COLLECTION)
	_, err = collection.InsertOne(ctx, ins)
	if err == nil {
		out.Success = true
	}

	return err
}

func (a *Credentials) Validate(ctx context.Context, in *credentials.Credential, out *credentials.CredentialResponse) error {
	out.Success = false

	filter := bson.M{
		"userid":         in.UserID,
		"credentialtype": in.CredentialType,
	}

	collection := a.client.Collection((CREDENTIALS_COLLECTION))
	result := &credentials.Credential{}

	if err := collection.FindOne(ctx, filter).Decode(result); err != nil {
		if err == mmongo.ErrNoDocuments {
			return a.InvalidCredential()
		}
		return err
	}

	switch ct := in.GetCredentialType(); ct {
	case credentials.CredentialType_PASSWORD:
		if err := bcrypt.CompareHashAndPassword(result.GetCredential(), in.GetCredential()); err != nil {
			return a.InvalidCredential()
		}
	default:
		return a.InvalidCredentialType(ct)
	}

	out.Success = true
	return nil
}

func (a *Credentials) Update(ctx context.Context, in *credentials.CredentialUpdate, out *credentials.CredentialResponse) error {
	return errors.NotImplemented(a.ID())
}

// Package handler defines the gRPC endpoint handlers for the Credentials service.
package handler

import (
	"context"

	credentials "github.com/koverto/credentials/api"

	"github.com/koverto/errors"
	"github.com/koverto/micro/v2"
	"github.com/koverto/mongo"
	"github.com/koverto/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mmongo "go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const credentialsCollection = "credentials"

// Credentials defines the Credentials service.
type Credentials struct {
	*Config
	*micro.Service
	client mongo.Client
}

// Config contains the configuration for an instance of the Credentials service
// handlers.
type Config struct {
	MongoURL string `json:"mongourl"`
}

// New creates a new instance of the Credentials service handlers.
func New(conf *Config, service *micro.Service) (*Credentials, error) {
	client, err := mongo.NewClient(conf.MongoURL, "credentials")
	if err != nil {
		return nil, err
	}

	var uidIndex mmongo.IndexModel
	uidIndex.Keys = bson.M{"userid": 1}

	var credIndex mmongo.IndexModel
	credIndex.Keys = bson.M{"credential": "hashed"}

	client.DefineIndexes(mongo.NewIndexSet(credentialsCollection, uidIndex, credIndex))

	if err := client.Connect(); err != nil {
		return nil, err
	}

	return &Credentials{conf, service, client}, nil
}

// Create inserts a new Credential object into the database.
func (a *Credentials) Create(
	ctx context.Context,
	in *credentials.Credential,
	out *credentials.CredentialResponse,
) error {
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

	collection := a.client.Collection(credentialsCollection)
	_, err = collection.InsertOne(ctx, ins)

	if err == nil {
		out.Success = true
	}

	return err
}

// Validate validates the given Credential object against the database-stored value.
func (a *Credentials) Validate(
	ctx context.Context,
	in *credentials.Credential,
	out *credentials.CredentialResponse,
) error {
	out.Success = false

	switch ct := in.GetCredentialType(); ct {
	case credentials.CredentialType_PASSWORD:
		filter := bson.M{
			"userid":         in.UserID,
			"credentialtype": in.CredentialType,
		}

		collection := a.client.Collection((credentialsCollection))
		result := &credentials.Credential{}

		if err := collection.FindOne(ctx, filter).Decode(result); err != nil {
			if err == mmongo.ErrNoDocuments {
				return a.InvalidCredential()
			}

			return err
		}

		if err := bcrypt.CompareHashAndPassword(result.GetCredential(), in.GetCredential()); err != nil {
			return a.InvalidCredential()
		}
	default:
		return a.InvalidCredentialType(ct)
	}

	out.Success = true

	return nil
}

// Update updates a set of credentials stored in the database.
func (a *Credentials) Update(
	ctx context.Context,
	in *credentials.CredentialUpdate,
	out *credentials.CredentialResponse,
) error {
	return errors.NotImplemented(a.Name)
}

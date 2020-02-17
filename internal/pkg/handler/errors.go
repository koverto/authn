package handler

import (
	credentials "github.com/koverto/credentials/api"

	"github.com/micro/go-micro/v2/errors"
)

func (a *Credentials) InvalidCredential() error {
	return errors.BadRequest(a.ID(), "invalid credentials")
}

func (a *Credentials) InvalidCredentialType(credentialType credentials.CredentialType) error {
	return errors.BadRequest(a.ID(), "invalid credential type: %s", credentialType)
}
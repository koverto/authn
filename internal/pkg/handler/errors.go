package handler

import (
	authn "github.com/koverto/authn/api"

	"github.com/micro/go-micro/v2/errors"
)

func (a *Authn) InvalidCredential() error {
	return errors.BadRequest(a.ID(), "invalid credentials")
}

func (a *Authn) InvalidCredentialType(credentialType authn.CredentialType) error {
	return errors.BadRequest(a.ID(), "invalid credential type: %s", credentialType)
}

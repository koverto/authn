package handler

import (
	credentials "github.com/koverto/credentials/api"

	"github.com/micro/go-micro/v2/errors"
)

// InvalidCredential returns a 400 Bad Request error if the given credentials could not be
// validated.
func (a *Credentials) InvalidCredential() error {
	return errors.BadRequest(a.Name, "invalid credentials")
}

// InvalidCredentialType returns a 400 Bad Request error if the given credentials contained an
// invalid CredentialType value.
func (a *Credentials) InvalidCredentialType(credentialType credentials.CredentialType) error {
	return errors.BadRequest(a.Name, "invalid credential type: %s", credentialType)
}

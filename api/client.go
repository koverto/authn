package credentials

import (
	"github.com/micro/go-micro/v2/client"
)

// Name is the identifying name of the Credentials service.
const Name = "com.koverto.svc.credentials"

// Client defines a client for the Credentials service.
type Client struct {
	CredentialsService
}

// NewClient creates a new client for the Credentials service.
func NewClient(client client.Client) *Client {
	return &Client{NewCredentialsService(Name, client)}
}

// Name returns the name of the Credentials service.
func (c *Client) Name() string {
	return Name
}

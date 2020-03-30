package credentials

import (
	"github.com/micro/go-micro/v2/client"
)

const Name = "com.koverto.svc.credentials"

type Client struct {
	CredentialsService
}

func NewClient(client client.Client) *Client {
	return &Client{NewCredentialsService(Name, client)}
}

func (c *Client) Name() string {
	return Name
}

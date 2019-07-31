package profitbricks

import (
	"github.com/hashicorp/terraform/httpclient"
	"github.com/profitbricks/profitbricks-sdk-go"
)

type Config struct {
	Username string
	Password string
	Endpoint string
	Retries  int
	Token    string

	terraformVersion string
}

// Client() returns a new client for accessing ProfitBricks.
func (c *Config) Client() (*profitbricks.Client, error) {
	var client *profitbricks.Client
	if c.Token != "" {
		client = profitbricks.NewClientbyToken(c.Token)
	} else {
		client = profitbricks.NewClient(c.Username, c.Password)
	}
	tfUserAgent := httpclient.UserAgent(c.terraformVersion)
	client.SetUserAgent(tfUserAgent.String())
	client.SetDepth(5)

	if len(c.Endpoint) > 0 {
		client.SetURL(c.Endpoint)
	}
	return client, nil
}

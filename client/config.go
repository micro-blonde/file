package client

type config struct {
	Enabled *bool

	enabled bool
}

func (c *config) Initialize() {
	c.enabled = c.Enabled == nil || *c.Enabled
}

package zenrows

const zenRowsAPIURLv1 = "https://api.zenrows.com/v1/"

// ClientConfig Configuration with the key and base API URL
type ClientConfig struct {
	key string

	BaseURL string
}

// DefaultConfig Generate default configuration -- currently only option but extensive for the future
func DefaultConfig() ClientConfig {
	return ClientConfig{
		BaseURL: zenRowsAPIURLv1,
	}
}

// ConfigCredentials Adds the apikey to the configuration -- in case they change the format of the credentials it will be easier to implement here
func (c *ClientConfig) ConfigCredentials(key string) {
	c.key = key
}

package zenrows

const zenRowsAPIURLv1 = "https://api.zenrows.com/v1/"

type ClientConfig struct {
	key string

	BaseURL string
}

func DefaultConfig() ClientConfig {
	return ClientConfig{
		BaseURL: zenRowsAPIURLv1,
	}
}

func (c *ClientConfig) ConfigCredentials(key string) {
	c.key = key
}

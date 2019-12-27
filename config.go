package dmt

type AuthConfig struct {
	Type   string            `json:"type" yaml:"type"`
	Users  map[string]string `json:"users" yaml:"users"`
	Secret string            `json:"secret" yaml:"secret"`
}

type RESTConfig struct {
	Port int          `json:"port" yaml:"port"`
	Host string       `json:"host" yaml:"host"`
	Auth []AuthConfig `json:"auth" yaml:"auth"`
}

type BrokerConfig struct {
	Protocol string `json:"protocol" yaml:"protocol"`
	Endpoint string `json:"endpoint" yaml:"endpoint"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	ClientID string `json:"client_id" yaml:"clientID"`
	Token    string `json:"token" yaml:"token"`
}

type ChannelConfig struct {
	URL    string                 `json:"url" yaml:"url"`
	Name   string                 `json:"name" yaml:"name"`
	Config map[string]interface{} `json:"config" yaml:"config"`
}

type Config struct {
	REST     RESTConfig      `json:"rest" yaml:"rest"`
	Broker   BrokerConfig    `json:"broker" yaml:"broker"`
	Channels []ChannelConfig `json:"channels" yaml:"channels"`
}

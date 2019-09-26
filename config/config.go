package config

type ServerConfig struct {
	Host      string `json:"host"`
	PortURL   string `json:"portUrl"`
	PortValue string `json:"portValue"`
}

type DatabaseConfig struct {
	DriverName			string `json:"driverName"`
	URL					string `json:"url"`
	ConnectionString	string `json:"connectionString"`
	MaxOpenConnections	int    `json:"maxOpenConnections"`
}

type SessionConfig struct {
	Name            string `json:"name"`
	Length          int    `json:"length"`
	LifetimeSeconds int    `json:"lifetime"`
	Path            string `json:"path"`
	HTTPOnly        bool   `json:"httpOnly"`
}

type AuthClient struct {
	URL     string `json:"url"`
	Address string `json:"address"`
}
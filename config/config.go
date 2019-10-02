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
	Name            string		`json:"name"`
	Value           string		`json:"length"`
	LifetimeSeconds int    		`json:"lifetime"`
	Secure			bool		`json:"secure"`
	Path 			string		`json:"path"`
	HTTPOnly        bool		`json:"httpOnly"`
}

type AuthConfig struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email 	 string `json:"email"`
}
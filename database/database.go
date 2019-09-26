package db

import (
	_ "github.com/lib/pq"
)

func main(c config.DatabaseConfig) {
	ConnectToDatabase(c.DriverName, c.URL)
}
package dbPsql

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectToDatabase(DriverName string, URL string) (db DB, err error) {
	newDb, err := sql.Open(DriverName, URL)
	if err != nil {
		panic(err)
		return
	}

	err = newDb.Ping()
	if err != nil {
		panic(err)
		return
	}

	db = DB{newDb}
	return
}

func CreateDB() {
	InitDatabase()
}
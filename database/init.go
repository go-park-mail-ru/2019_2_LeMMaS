// схема бд сервиса

package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

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

func (db *DB) InitDatabase() () { //TODO: сделать создание таблиц под каждую сущность
	var _, err = db.Exec(` 
CREATE TABLE User (
id 		int NOT NULL primary key,
login   char(15) NOT NULL,
passwordHash   varchar(255) NOT NULL,
avatarAddress  varchar(255),
role  char(15) NOT NULL
);

`)
if err != nil {
	panic(err)
	return
}
return
}
// формирование бд сервиса

package dbPsql

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func InitDatabase() () { //TODO: сделать создание таблиц под каждую сущность
	var _, err = DB.Exec(` 
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
package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

const (
	Host     = "database"
	Port     = 5432
	User     = "postgres"
	Password = 753159
	DBname   = "lojaDB"
)

func Connection() {
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable", Host, Port, User, Password, DBname)

	db, err = sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err.Error())
	}
}

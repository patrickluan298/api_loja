package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Printf("Erro ao conectar ao banco de dados: %v", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		log.Printf("Erro ao fazer ping no banco de dados: %v", err)
	}
}

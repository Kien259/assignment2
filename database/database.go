package database

import (
	"log"
	"os"

	"entgo.io/ent/dialect"
)

func ConnectDb() {
	dsn := "postgres://user_demo:root@localhost:5432/postgres?sslmode=disable"
	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatal("Failed opening connection to postgres: ", err)
		os.Exit(1)
	}
	log.Println("Connected to database")
	DBConn = client
}

package database

import (
	"log"
	"os"

	"entgo.io/ent/dialect"
	_ "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

var DBConn *ent.Client



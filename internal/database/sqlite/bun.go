package sqlite

import (
	"database/sql"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/users"
)

var db *bun.DB

func BunConnect() {
	//Reference: // https://bun.uptrace.dev/guide/golang-orm.html#using-bun-with-existing-code
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:tuhuebot.db?cache=shared")
	if err != nil {
		log.Fatal(err)
	}
	// Create a Bun database client.
	db = bun.NewDB(sqldb, sqlitedialect.New())
	Migrate()
}

func DB() *bun.DB {
	return db
}

func Migrate() {
	users.CreateTable(db)
}

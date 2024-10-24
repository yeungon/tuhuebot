package pg

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/yeungon/tuhuebot/internal/config"
)

var dbpg *bun.DB

func Connect() {
	var pg_url = config.Get().PG
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(pg_url)))
	dbpg = bun.NewDB(sqldb, pgdialect.New())
}

func PG() *bun.DB {
	return dbpg
}

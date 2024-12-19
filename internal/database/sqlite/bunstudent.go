package sqlite

import (
	"database/sql"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/yeungon/tuhuebot/internal/database/sqlite/students"
)

var dbstudent *bun.DB

func BunStudentConnect() {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:students.db?cache=shared")
	if err != nil {
		log.Fatal(err)
	}
	// Create a Bun database client.
	dbstudent = bun.NewDB(sqldb, sqlitedialect.New())
	MigrateStudent()
}

func DBSTUDENT() *bun.DB {
	return dbstudent
}

func MigrateStudent() {
	students.CreateTableStudents(dbstudent)
}

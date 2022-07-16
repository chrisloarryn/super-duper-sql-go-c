package storage

import (
	"GO_MINE/crud_psql/pkg/product"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
)

var (
	db                  *sql.DB
	once                sync.Once
	pgDataSourceName    = "postgres://postgres:postgres@localhost:65432/postgres?sslmode=disable"
	mysqlDataSourceName = "root:mysql@tcp(localhost:3308)/mysqldb?parseTime=true"
)

type Driver string

const (
	MySQL    Driver = "mysql"
	Postgres Driver = "postgres"
)

// New create a new instance of db
func New(d Driver) *sql.DB {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
	return db
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", pgDataSourceName)
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

func newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", mysqlDataSourceName)
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("conectado a mysql")
	})
}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}

	return null
}

func DAOProduct(d Driver) (product.Storage, error) {
	switch d {
	case MySQL:
		return newMySQLProduct(db), nil
	case Postgres:
		return newPsqlProduct(db), nil
	}
	return nil, fmt.Errorf("driver %s not supported", d)
}

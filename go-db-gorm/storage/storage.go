package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"sync"
)

var (
	db                  *gorm.DB
	once                sync.Once
	pgDataSourceName    = "postgres://postgres:postgres@localhost:65432/postgres?sslmode=disable"
	mysqlDataSourceName = "root:mysql@tcp(localhost:3308)/mysqldbd?parseTime=true"
)

type Driver string

const (
	MySQL    Driver = "mysql"
	Postgres Driver = "postgres"
)

// New create a new instance of db
func New(d Driver) *gorm.DB {
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
		db, err = gorm.Open("postgres", pgDataSourceName)
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

func newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("mysql", mysqlDataSourceName)
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("conectado a mysql")
	})
}

// DB return a unique instance of db
func DB() *gorm.DB {
	return db
}

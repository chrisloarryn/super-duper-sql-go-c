package storage

import (
	"GO_MINE/crud_psql/pkg/invoiceheader"
	"database/sql"
	"fmt"
	"log"
)

const (
	mysqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers (
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP
	)`
	mysqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES(?)`
)

// mysqlInvoiceHeaders is model used for work with mysql - invoice_headers
type mysqlInvoiceHeaders struct {
	db *sql.DB
}

// NewMySQLInvoiceHeader return a new pointer to PsqlInvoiceHeader
func NewMySQLInvoiceHeader(db *sql.DB) *mysqlInvoiceHeaders {
	return &mysqlInvoiceHeaders{db}
}

func (p *mysqlInvoiceHeaders) Migrate() error {
	stmt, err := p.db.Prepare(mysqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("::::::::::::::::::")
	fmt.Println("Migration invoice_headers complete")
	fmt.Println("::::::::::::::::::")

	return nil
}

// CreateTx implement the interface invoiceHeader.Storage
func (p *mysqlInvoiceHeaders) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(mysqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	result, err := stmt.Exec(m.Client)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = uint(id)

	return nil // QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
}

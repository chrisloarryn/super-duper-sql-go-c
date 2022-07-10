package storage

import (
	"GO_MINE/crud_psql/pkg/invoiceheader"
	"database/sql"
	"fmt"
	"log"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers (
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
	)`
	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES($1) RETURNING id, created_at`
)

// psqlInvoiceHeaders is model used for work with postgresql - invoice_headers
type psqlInvoiceHeaders struct {
	db *sql.DB
}

// NewPsqlInvoiceHeader return a new pointer to PsqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *psqlInvoiceHeaders {
	return &psqlInvoiceHeaders{db}
}

func (p *psqlInvoiceHeaders) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
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
func (p *psqlInvoiceHeaders) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
}

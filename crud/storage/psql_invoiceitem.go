package storage

import (
	"GO_MINE/crud_psql/pkg/invoiceitem"
	"database/sql"
	"fmt"
	"log"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items (
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
	psqlCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES($1, $2) RETURNING id, created_at`
)

// psqlInvoiceHeaders is model used for work with postgresql - invoice_items
type psqlInvoiceItems struct {
	db *sql.DB
}

// NewPsqlInvoiceItem return a new pointer to PsqlInvoiceItem
func NewPsqlInvoiceItem(db *sql.DB) *psqlInvoiceItems {
	return &psqlInvoiceItems{db}
}

func (p *psqlInvoiceItems) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
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
	fmt.Println("Migration invoice_items complete")
	fmt.Println("::::::::::::::::::")

	return nil
}

// CreateTx implement the interface invoiceitem.Storage
func (p *psqlInvoiceItems) CreateTx(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	for _, item := range ms {
		err = stmt.QueryRow(headerID, item.ProductID).Scan(
			&item.ID,
			&item.CreatedAt,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
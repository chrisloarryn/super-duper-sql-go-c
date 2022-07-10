package storage

import (
	"GO_MINE/crud_psql/pkg/invoiceitem"
	"database/sql"
	"fmt"
	"log"
)

const (
	mysqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items (
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
	mysqlCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES(?, ?)`
)

// mysqlInvoiceItems is model used for work with mysql - invoice_items
type mysqlInvoiceItems struct {
	db *sql.DB
}

// NewMySQLInvoiceItem return a new pointer to MySQLInvoiceItem
func NewMySQLInvoiceItem(db *sql.DB) *mysqlInvoiceItems {
	return &mysqlInvoiceItems{db}
}

func (p *mysqlInvoiceItems) Migrate() error {
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
func (p *mysqlInvoiceItems) CreateTx(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(mysqlCreateInvoiceItem)
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
		result, err := stmt.Exec(headerID, item.ProductID)
		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		item.ID = uint(id)
	}

	fmt.Printf("se creo el invoice_item correctamente con ID: %d\n", ms[0].ID)

	return nil
}

package storage

import (
	"GO_MINE/crud_psql/pkg/invoice"
	"GO_MINE/crud_psql/pkg/invoiceheader"
	"GO_MINE/crud_psql/pkg/invoiceitem"
	"database/sql"
	"fmt"
	"log"
)

// PsqlInvoice is used for work with postgresql - invoice
type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItem   invoiceitem.Storage
}

// NewPsqlInvoice return a new pointer to PsqlInvoice
func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItem:   i,
	}
}

// Create implement the interface invoice.Storage
func (p *PsqlInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			log.Fatal(err)
		}
	}(tx)

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	fmt.Printf("Invoice created with id: %d\n", m.Header.ID)

	if err := p.storageItem.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	fmt.Printf("Invoice items created with id: %d and length of %d\n", m.Header.ID, len(m.Items))

	return tx.Commit()
}

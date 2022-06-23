package storage

import (
	"GO_MINE/crud_psql/pkg/product"
	"database/sql"
	"fmt"
	"log"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products (
		id SERIAL NOT NULL,
		name VARCHAR(100) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`
	psqlCreateProduct  = `INSERT INTO products (name, observations, price, created_at) VALUES ($1, $2, $3, $4)`
	psqlGetAllProduct  = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	psqlGetByIDProduct = psqlGetAllProduct + " WHERE id = $1"
	psqlUpdateProduct  = "UPDATE products SET name = $1, observations = $2, price = $3, updated_at = $4 WHERE id = $5"
	psqlDeleteProduct  = "DELETE FROM products WHERE id = $1"
)

// psqlProduct is model used for work with postgresql - products
type psqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct return a new pointer to PsqlProduct
func NewPsqlProduct(db *sql.DB) *psqlProduct {
	return &psqlProduct{db}
}

func (p *psqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
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
	fmt.Println("Migration product complete")
	fmt.Println("::::::::::::::::::")

	return nil
}

func (p *psqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)
	if err != nil {
		return err
	}

	fmt.Println(":::::::::::::::::::::::")
	fmt.Println("Create product complete")
	fmt.Println(":::::::::::::::::::::::")

	return nil
}

func (p *psqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, *m)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("::::::::::::::::::")
	fmt.Println("Get all product complete")
	fmt.Println("::::::::::::::::::")

	return ms, nil
}

func (p *psqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetByIDProduct)
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	return scanRowProduct(stmt.QueryRow(id))
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}

	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull)
	if err != nil {
		return nil, err
	}

	m.Observations = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}

// Update implement the interface product.Storage
func (p *psqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("id: %d not found", m.ID)
	}

	fmt.Println("::::::::::::::::::") // Print message to console
	fmt.Println("Update product complete")
	fmt.Println("::::::::::::::::::")

	return nil
}

// Delete implement the interface product.Storage
func (p *psqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}

	fmt.Println("::::::::::::::::::") // Print message to console
	fmt.Println("Deleting product id: ", id)
	fmt.Println("::::::::::::::::::")

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(stmt)

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("id: %d not found", id)
	}

	fmt.Println("::::::::::::::::::") // Print message to console
	fmt.Println("Delete product complete")
	fmt.Println("::::::::::::::::::")

	return nil
}

package invoiceheader

import (
	"database/sql"
	"time"
)

// Model for invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Models slice of models
type Models []Model

type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, *Model) error
	//Create(*Model) error
	//Update(*Model) error
	//GetAll() (Models, error)
	//GetByID(uint) (*Model, error)
	//Delete(uint) error
}

type Service struct {
	storage Storage
}

// NewService returns a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used to create the table in the database
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}

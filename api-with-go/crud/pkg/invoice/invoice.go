package invoice

import (
	"GO_MINE/crud_psql/pkg/invoiceheader"
	"GO_MINE/crud_psql/pkg/invoiceitem"
)

type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

type Storage interface {
	Create(*Model) error
	//GetAll () ([]*Model, error)
}

// Service of invoice
type Service struct {
	storage Storage
}

// NewService returns a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Create is used to create an invoice
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}

package product

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Model for product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

var (
	ErrIDNotFound = errors.New("product does no contain id")
)

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s",
		m.ID, m.Name, m.Observations, m.Price,
		m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

func (m Models) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%02s | %-20s | %-20s | %5s | %10s | %10s\n",
		"id", "name", "observations", "price", "created_at", "updated_at"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}

// Models slice of models
type Models []Model

type Storage interface {
	Migrate() error
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
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

// Create // Create is used to create a product
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

// GetAll is used to get all products
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

// GetByID is used to get one product
func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}

// Update is used to update a product
func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIDNotFound
	}
	m.UpdatedAt = time.Now()
	return s.storage.Update(m)
}

// Delete is used to delete a product
func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}

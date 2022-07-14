package handler

import "crud_psql_7/model"

// Storage .
type Storage interface {
	Create(p *model.Person) error
	Update(ID int, p *model.Person) error
	Delete(ID int) error
	GetByID(ID int) (*model.Person, error)
	GetAll() (model.People, error)
}

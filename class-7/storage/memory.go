package storage

import (
	"crud_psql_7/model"
	"fmt"
)

type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

// NewMemory returns a new instance of db
func NewMemory() Memory {
	people := make(map[int]model.Person)

	return Memory{currentID: 0, Persons: people}
}

func (m *Memory) Create(p *model.Person) error {
	if p == nil {
		return model.ErrPersonCanNotBeNil
	}

	m.currentID++
	m.Persons[m.currentID] = *p

	return nil
}

// Update memory method
func (m *Memory) Update(ID int, p *model.Person) error {
	if p == nil {
		return model.ErrPersonCanNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	m.Persons[ID] = *p

	return nil
}

// Delete deletes a person from memory
func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	delete(m.Persons, ID)

	return nil
}

// GetByID retrieves a person by id from memory
func (m *Memory) GetByID(ID int) (*model.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return nil, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	return &person, nil
}

// GetAll returns all the entries that are in memory
func (m *Memory) GetAll() (model.People, error) {
	var people model.People
	for _, p := range m.Persons {
		people = append(people, &p)
	}

	return people, nil
}

package storage

import (
	"testing"

	"class5/api/model"
)

func TestCreate(t *testing.T) {
	m := NewMemory()
	table := []struct {
		name      string
		person    *model.Person
		wantError error
	}{
		{"Empty person", nil, model.ErrPersonCanNotBeNil},
		{"Cristobal", &model.Person{Name: "Cristobal"}, nil},
		{"Valentin", &model.Person{Name: "Valentin"}, nil},
		{"Aracely", &model.Person{Name: "Aracely"}, nil},
	}
	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			if gotError := m.Create(v.person); gotError != v.wantError {
				t.Errorf("want %v, got %v", v.wantError, gotError)
			}
		})
	}

	wantCount := len(table) - 1
	if m.currentID != wantCount {
		t.Errorf("currentID should be %d, and is %d", wantCount, m.currentID)
	}
}

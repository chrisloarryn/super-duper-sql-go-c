package handler_test

import (
	repo "class4"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

//go test -v -run=TestGet

func TestGet(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	repo.Get(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, w.Code)
	}

	if ctype := w.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("want %s, got %s", "application/json", ctype)
	}

	//t.Logf("%s", w.Body.String())
	p := new(Person)
	err := json.NewDecoder(w.Body).Decode(&p)
	if err != nil {
		t.Errorf("want nil, got %v", err)
		t.Errorf("json error: %v", err)
	}

	want := &Person{
		Name: "Jhoana",
		Age:  31,
	}

	if !reflect.DeepEqual(want, p) {
		t.Errorf("want %v, got %v", want, p)
	}
}

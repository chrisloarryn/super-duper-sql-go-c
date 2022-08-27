package handler_test

import (
	repo "class4"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type PersonWithEcho struct {
	Name string
	Age  int
}

//go test -v -run=TestGet

func TestGetWithEcho(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	e := echo.New()

	ctx := e.NewContext(r, w)

	err := repo.GetWithEcho(ctx)
	if err != nil {
		t.Errorf("want nil, got %v", err)
		t.Errorf("error: %v", err)
	}

	wantHTTP := struct {
		sc int
		h  string
	}{
		sc: http.StatusOK,
		h:  "application/json; charset=UTF-8",
	}

	if w.Code != wantHTTP.sc {
		t.Errorf("want %d, got %d", http.StatusOK, w.Code)
	}

	if ctype := w.Header().Get("Content-Type"); ctype != wantHTTP.h {
		t.Errorf("want %s, got %s", "application/json", ctype)
	}

	//t.Logf("%s", w.Body.String())
	p := new(PersonWithEcho)

	err = json.NewDecoder(w.Body).Decode(&p)
	if err != nil {
		t.Errorf("want nil, got %v", err)
		t.Errorf("json error: %v", err)
	}

	want := &PersonWithEcho{
		Name: "Jhoana",
		Age:  31,
	}

	if !reflect.DeepEqual(want, p) {
		t.Errorf("want %v, got %v", want, p)
	}
}

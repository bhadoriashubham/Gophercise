package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestHello(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hello)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := strings.Contains(rr.Body.String(), "<h1>Hello!</h1>")
	assert.Equal(t, true, expected, "they should be equal")
}
func TestFuncThatPanics(t *testing.T) {
	assert.Panic(t, funcThatPanics, "Oh no!")
}
func TestPanicDemo(t *testing.T) {
	req, err := http.NewRequest("GET", "/panic", nil)
	if err != nil {
		t.Fatalf("not able to request %v", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	panicDemo(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("not expected error in panic %v", res.StatusCode)
	}
}

func TestPanicAfterDemo(t *testing.T) {
	req, err := http.NewRequest("GET", "/panic-after", nil)
	if err != nil {
		t.Fatalf("not able to request %v", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	panicAfterDemo(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("not expected error in panic %v", res.StatusCode)
	}
}
func Test_recoverMw(t *testing.T) {
	handler := http.HandlerFunc(panicDemo)
	executeRequest("Get", "/panic", recoverMw(handler))
}

func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	rr.Result()
	handler.ServeHTTP(rr, req)
	return rr, err
}

func TestSourceCodeHandlerNeg(t *testing.T) {
	req, err := http.NewRequest("GET", "/debug", nil)
	if err != nil {
		t.Fatalf("not able to request %v", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	sourceCodeHandler(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("not expected error in panic %v", res.StatusCode)
	}
}
func TestSourceCodeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/debug?line=24&path=/home/gslab/hey/src/recover/main.go", nil)
	if err != nil {
		t.Fatalf("not able to request %v", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	sourceCodeHandler(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("not expected error in panic %v", res.StatusCode)
	}
}
func TestSourceCodeHandlerCopyFailed(t *testing.T) {

	req, err := http.NewRequest("GET", "/debug?line=24&path=/home/gslab/hey/src/recover/main.go", nil)
	if err != nil {
		t.Fatalf("not able to request %v", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	sourceCodeHandler(rec, req)
	sourceCodeHandler(rec, req)
	res := rec.Result()
	if res.StatusCode == http.StatusInternalServerError {
		t.Errorf("not expected error in panic %v", res.StatusCode)
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "t1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			main()
		})
	}
}

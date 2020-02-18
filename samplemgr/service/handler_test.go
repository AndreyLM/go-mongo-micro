package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/andreylm/go-mongo-micro/sqmplemgr/db"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestCreateSweat(t *testing.T) {
	var jsonStr = []byte(`{ "glucose": 1.12, "sodium": 0.98, "chloride": 0.003 }`)
	req, err := http.NewRequest("POST", "/entry", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createSweatHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wring status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetSweatSamples(t *testing.T) {
	req, err := http.NewRequest("GET", "/entry", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getSweatSamplesHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Invalid status code: got %v want %v", status, http.StatusOK)
	}

	var sweats []db.Sweat
	err = json.Unmarshal(rr.Body.Bytes(), &sweats)
	if err != nil {
		t.Fatal(err)
	}
	if len(sweats) < 1 {
		t.Errorf("No sweat samples. Expected at least 1")
	}
}

func TestGetSweat(t *testing.T) {
	t.Log("Testign GET /sweat/{id}")

	testsTable := map[string]struct {
		URL    string
		Params map[string]string
		Code   int
	}{
		"ValidRequest": {
			URL:    "/sweat",
			Params: map[string]string{"id": "5e345232cf662d32b2a3bb2c"},
			Code:   http.StatusOK,
		},
		"InvalidRequest": {
			URL:    "/sweat",
			Params: map[string]string{"id": "someinvalidid"},
			Code:   http.StatusNotFound,
		},
	}

	for i, v := range testsTable {
		t.Logf("\tTest '%s'. Checking URL '%s' with code %d", i, v.URL, v.Code)
		{
			req, err := http.NewRequest("GET", v.URL, nil)
			if err != nil {
				t.Fatalf("\t%s\t Cannot create request: %v", failed, err)
			} else {
				t.Logf("\t%s\t Request created", succeed)
			}
			req = mux.SetURLVars(req, v.Params)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getSweatByIDHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != v.Code {
				t.Fatalf("\t%s\t Invalid response status: want %v, get %v", failed, v.Code, status)
			} else {
				t.Logf("\t%s\t Request successful", succeed)
			}
		}
	}

}

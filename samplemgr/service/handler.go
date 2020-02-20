package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andreylm/go-mongo-micro/sqmplemgr/db"
	"github.com/gorilla/mux"
)

// PingResponse - ping response
type PingResponse struct {
	Message string `json:"message,omitempty"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	response := PingResponse{Message: "pong"}
	status := http.StatusOK
	respBytes, err := json.Marshal(response)
	if err != nil {
		panic(err)
		// status = http.StatusInternalServerError
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(respBytes)
}

func createSweatHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var s db.Sweat

	if err := decoder.Decode(&s); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r = WithUserContext(r)
	if err := s.Create(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func getSweatSamplesHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	sweats, err := db.ListAllSweat()
	if err != nil {
		fmt.Println("Error fetching data", err)
		status = http.StatusInternalServerError
	}
	respBytes, err := json.Marshal(sweats)
	if err != nil {
		fmt.Println("Error marhaling data", err)
		status = http.StatusInternalServerError
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(respBytes)

}

func getSweatByIDHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK

	vars := mux.Vars(r)
	log.Println(r.URL.String())
	ID, ok := vars["id"]
	if !ok {
		fmt.Println("Error getting id")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	sweat, err := db.GetByID(ID)
	if err != nil {
		fmt.Println("Error fetching data", err)
		status = http.StatusNotFound
	}

	respBytes, err := json.Marshal(sweat)
	if err != nil {
		fmt.Println("Error marhaling data", err)
		status = http.StatusInternalServerError
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(respBytes)
}

// WithUserContext ...
func WithUserContext(r *http.Request) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "UserID", r.Header.Get("UserID"))

	return r.WithContext(ctx)
}

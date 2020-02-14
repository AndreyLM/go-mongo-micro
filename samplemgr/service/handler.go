package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andreylm/go-mongo-micro/sqmplemgr/db"
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
	status := http.StatusOK

	var s db.Sweat

	if err := decoder.Decode(&s); err != nil {
		status = http.StatusInternalServerError
		panic(err)
	}

	fmt.Println(s)
	if err := s.Create(); err != nil {
		status = http.StatusInternalServerError
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
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
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(respBytes)

}

func getSweatByIDHandler(w http.ResponseWriter, r *http.Request) {

}

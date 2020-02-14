package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept"
	appName       = "sweatmgr"
)

// InitRouter - router initialization
func InitRouter() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	// Version 1 API management
	v1 := fmt.Sprintf("application/vnd.%s.v1", appName)

	r.HandleFunc("/sweat", createSweatHandler).Methods(http.MethodPost).Headers(versionHeader, v1)
	r.HandleFunc("/sweat_samples", getSweatSamplesHandler).Methods(http.MethodGet).Headers(versionHeader, v1)
	r.HandleFunc("/sweat/{id}", getSweatByIDHandler).Methods(http.MethodGet).Headers(versionHeader, v1)

	// Version 2 API management
	v2 := fmt.Sprintf("application/vnd.%s.v2", appName)

	r.HandleFunc("/sweat", createSweatHandler).Methods(http.MethodPost).Headers(versionHeader, v2)
	r.HandleFunc("/sweat_samples", getSweatSamplesHandler).Methods(http.MethodGet).Headers(versionHeader, v2)
	r.HandleFunc("/sweat/{id}", getSweatByIDHandler).Methods(http.MethodGet).Headers(versionHeader, v2)

	return
}

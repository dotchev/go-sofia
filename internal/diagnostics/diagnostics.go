package diagnostics

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewDiagnostics() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", healthz)
	router.HandleFunc("/ready", ready)
	return router
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Healthy")
}

func ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ready")
}

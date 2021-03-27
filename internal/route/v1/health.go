package v1

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

//Health the health endpoint handler
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	hostname, _ := os.Hostname()
	w.Write([]byte(fmt.Sprintf("Kubeshell is health!\n\nI am %s", hostname)))
}

func healthRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", Health)
	return r
}

package v1

import (
	"net/http"

	"github.com/go-chi/chi"
)

//K8SMainRouter K8S main router
func K8SMainRouter() http.Handler {
	r := chi.NewRouter()
	r.Mount("/health", healthRouter())
	return r
}

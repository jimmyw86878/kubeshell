package gateway

import (
	"context"
	"fmt"

	//"gatekeeper/internal/constants"

	v1 "gatekeeper/internal/route/v1"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
)

//StartServeK8S for demonstration of K8S
func StartServeK8S(c chan os.Signal) {
	r := chi.NewRouter()
	r.Mount("/kubeshell/v1", v1.K8SMainRouter())
	srv := http.Server{Addr: ":5666", Handler: r}
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("Shutting down...")
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()
			srv.Shutdown(ctx)
		}
	}()
	fmt.Println(fmt.Sprintf("Start kubeshell http service at port: %d", 5666))
	srv.ListenAndServe()
}

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
	port := loadStrEnv("port", "5666")
	srv := http.Server{Addr: fmt.Sprintf(":%s", port), Handler: r}
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("Shutting down...")
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()
			srv.Shutdown(ctx)
		}
	}()
	fmt.Println(fmt.Sprintf("Start kubeshell http service at port: %s", port))
	srv.ListenAndServe()
}

func loadStrEnv(envName string, defaultValue string) string {
	res := defaultValue
	if val, exists := os.LookupEnv(envName); exists {
		res = val
	}
	return res
}

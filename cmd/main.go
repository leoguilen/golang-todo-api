package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/leoguilen/simple-go-api/pkg/api/routes"
)

var (
	app_port string = os.Getenv("PORT")
)

func main() {
	r := mux.NewRouter()

	if os.Getenv("USE_SWAGGER") == "true" {
		r.Handle("/swagger.yaml", http.FileServer(http.Dir(os.Getenv("SWAGGER_DIR"))))

		// documentation for developers
		opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml", Path: "swagger"}
		swagger := middleware.SwaggerUI(opts, nil)
		r.Handle("/swagger", swagger)

		// documentation for share
		opts1 := middleware.RedocOpts{SpecURL: "/swagger.yaml", Path: "docs"}
		docs := middleware.Redoc(opts1, nil)
		r.Handle("/docs", docs)
	}

	sr := r.PathPrefix("/api").Subrouter()
	routes.MapHealthcheckRoutes(sr)
	routes.MapTodoRoutes(sr)

	fmt.Printf("Server started at http://localhost:%v", app_port)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%v", app_port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %v", err)
	}
}

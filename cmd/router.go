package cmd

import (
	"fmt"
	"net/http"

	"github.com/aminasadiam/Shop/config"
	"github.com/aminasadiam/Shop/handler"
)

func Router(config *config.Config) error {
	mux := http.NewServeMux()

	initRoutes(mux)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler: mux,
	}

	fmt.Printf("Server starting on %s\n", server.Addr)
	return server.ListenAndServe()
}

func initRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handler.IndexHandler)

	// user authentication
	mux.HandleFunc("/login", handler.LoginHandler)
	mux.HandleFunc("/register", handler.RegisterHandler)

	// product
	mux.HandleFunc("/products", handler.ProductListHandler)
	mux.HandleFunc("/categories", handler.ProductCategoriesHandler)
}

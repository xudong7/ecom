package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xudong7/ecom/service/products"
	"github.com/xudong7/ecom/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// user routes
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	// products routes
	productsStore := products.NewStore(s.db)
	productHandler := products.NewHandler(productsStore)
	productHandler.RegisterRoutes(subRouter)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}

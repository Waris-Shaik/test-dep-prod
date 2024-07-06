package api

import (
	"database/sql"
	"log"
	"net/http"
	"test-dep-prod/services/user"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// userStore
	userStore := user.NewStore(s.db)

	// userHandler
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	log.Printf("server is listening on PORT%v ⚡⚡⚡\n", s.addr)
	return http.ListenAndServe(s.addr, router)
}

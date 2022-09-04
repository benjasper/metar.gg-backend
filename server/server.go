package server

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"metar.gg/ent"
	"metar.gg/graph"
	"metar.gg/logging"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(db *ent.Client, logger *logging.Logger) error {
	if err := db.Schema.Create(
		context.Background(),
	); err != nil {
		logger.Fatal(err)
	}

	port := "80"

	// Configure the server and start listening on :80.
	srv := handler.NewDefaultServer(graph.NewSchema(db))
	http.Handle("/graphql", srv)
	logger.Info(fmt.Sprintf("Server started and listening on port %s", port))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return err
	}

	return nil
}

package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"log"
	"metar.gg/ent"
	"metar.gg/graph"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(db *ent.Client) error {

	if err := db.Schema.Create(
		context.Background(),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :80.
	srv := handler.NewDefaultServer(graph.NewSchema(db))
	http.Handle("/graphql", srv)
	log.Println("listening on :80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("http server terminated", err)
	}

	return nil
}

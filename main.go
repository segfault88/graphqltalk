package main

import (
	"context"
	_ "embed"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/zerolog/log"
	"github.com/segfault88/graphqltalk/store"
)

func main() {
	storage := store.New()
	startHTTP(storage)
}

//go:embed schema.graphql
var schema string

func startHTTP(storage *store.Service) {
	gqlSchema := graphql.MustParseSchema(schema, &RootResolver{}, graphql.UseFieldResolvers())

	r := mux.NewRouter()
	r.Handle("/query", &relay.Handler{Schema: gqlSchema})

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info().Msg("starting server")
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Error().Err(err).Msg("http server error, canceling context")
		return
	}
}

type RootResolver struct {
	storage *store.Service
}

func (r *RootResolver) Movio(ctx context.Context) *RootResolver {
	return r
}

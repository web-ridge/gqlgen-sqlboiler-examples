// THIS IS NOT A REAL EXAMPLE
// Please look at social-network for real server

package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-38-plugin-panics/graphql_models"
)

func main() {

	port := os.Getenv("PORT")

	c := graphql_models.Config{
		Resolvers: &Resolver{
			db: nil,
		},
	}
	srv := handler.New(graphql_models.NewExecutableSchema(c))
	r := chi.NewRouter()
	r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	r.Handle("/graphql", srv)

	log.Info().Str("host", "localhost").Str("port", port).Msg("Connect GraphQL playground")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal().Err(err).Str("Could not listen to port", port).Send()
	}
}

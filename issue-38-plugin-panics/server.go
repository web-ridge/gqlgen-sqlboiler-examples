// THIS IS NOT A REAL EXAMPLE
// Please look at social-network for real server

package main

import (
	"context"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/issue-12-string-ids/graphql_models"
)

func main() {

	port := os.Getenv("PORT")

	c := graphql_models.Config{
		Resolvers: &Resolver{
			db: nil,
		},
	}

	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		// THIS IS ONLY TO CHECK IF PROGRAM RUNS
		return next(ctx)
	}

	srv := handler.New(graphql_models.NewExecutableSchema(c))

	// srv.Use(apollotracing.Tracer{})

	r := chi.NewRouter()
	r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	r.Handle("/graphql", srv)

	log.Info().Str("host", "localhost").Str("port", port).Msg("Connect GraphQL playground")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal().Err(err).Str("Could not listen to port", port).Send()
	}
}

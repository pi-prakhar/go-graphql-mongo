package api

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pi-prakhar/go-graphql-mongo/internal/graph"
)

func gqlPlaygroundHandler() http.HandlerFunc {
	return playground.Handler("GraphQL playground", "/query")
}

func queryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})).ServeHTTP(w, r)
	}
}

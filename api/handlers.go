package api

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pi-prakhar/go-graphql-mongo/internal/graph"
	model "github.com/pi-prakhar/go-graphql-mongo/internal/models"
	"github.com/pi-prakhar/go-graphql-mongo/pkg/logger"
)

func gqlPlaygroundHandler() http.HandlerFunc {
	return playground.Handler("GraphQL playground", "/api/query")
}

func queryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})).ServeHTTP(w, r)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	res := model.SuccessResponse[string]{
		StatusCode: http.StatusOK,
		Message:    "Success!",
		Data:       "hello world",
	}
	res.WriteJSON(w, http.StatusOK)
	logger.Log.Info("Successfully called test handler")
}

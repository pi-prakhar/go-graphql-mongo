package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pi-prakhar/go-graphql-mongo/internal/config/db"
	"github.com/pi-prakhar/go-graphql-mongo/internal/graph"
	"github.com/pi-prakhar/go-graphql-mongo/pkg/logger"
	"github.com/pi-prakhar/utils/loader"
)

func init() {
	logger.InitLogger()
	logger.Log.Info("GO-PHONE-OTP-SERVICE Logger Started")

	err := loader.LoadEnv()

	if err != nil {
		logger.Log.Error("Failed to Load ENV", err)
	}

	db.Connect()
}

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"graphql-gen/graph"
	"graphql-gen/graph/config"
	"graphql-gen/graph/generated"
	"graphql-gen/graph/repository/store"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	conf := config.Get()
	repo := store.New(conf.MongoDBEndpoint, conf.MongoDBName, conf.MongoDBTableName)
	//repo := store.NewClient(client.New(), conf.HOST, conf.PATH)

	resolver := &graph.Resolver{Repo: repo}
	gqlConfig := generated.Config{Resolvers: resolver}
	schema := generated.NewExecutableSchema(gqlConfig)
	srv := handler.NewDefaultServer(schema)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

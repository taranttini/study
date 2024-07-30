package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/graph"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/database"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=uDB password=pDB! dbname=desafio21 sslmode=disable")
	if err != nil {
		log.Panicf("failed to open database: %v", err)
	}
	defer db.Close()

	orderDb := database.NewOrder(db)
	itemDb := database.NewItem(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		OrderDB: orderDb,
		ItemDB:  itemDb,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

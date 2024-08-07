package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/database"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/infra/graph"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/infra/pb"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/service"
	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

func main() {
	listenerHttp, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	listenerGrpc, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// db --start
	db, err := sql.Open("postgres", "host=localhost port=5432 user=uDB password=pDB! dbname=desafio21 sslmode=disable")
	if err != nil {
		log.Panicf("failed to open database: %v", err)
	}
	defer db.Close()

	orderDb := database.NewOrder(db)
	itemDb := database.NewItem(db)
	// db --end

	g := new(errgroup.Group)
	g.Go(func() error { return StartGrpcServer(orderDb, itemDb, listenerGrpc) })
	g.Go(func() error { return StartGraphQl(orderDb, itemDb, listenerHttp) })

	fmt.Printf("running server: %v \n", listenerGrpc.Addr())
	fmt.Printf("running server: %v \n", listenerHttp.Addr())

	g.Wait()
}

func StartGrpcServer(orderDb *database.Order, itemDb *database.Item, l net.Listener) error {
	orderService := service.NewOrderService(*orderDb, *itemDb)

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	return grpcServer.Serve(l)
}

func StartGraphQl(orderDb *database.Order, itemDb *database.Item, l net.Listener) error {

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		OrderDB: orderDb,
		ItemDB:  itemDb,
	}}))

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	httpServer := &http.Server{Handler: mux}
	return httpServer.Serve(l)
}

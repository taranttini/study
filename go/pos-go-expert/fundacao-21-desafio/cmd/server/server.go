package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/database"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/entity"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/infra/graph"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/infra/pb"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/service"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/usecase"
	"golang.org/x/sync/errgroup"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

func main() {
	listenerHttp, err := net.Listen("tcp", "app:8080")
	if err != nil {
		log.Fatal(err)
	}
	listenerGrpc, err := net.Listen("tcp", "app:50051")
	if err != nil {
		log.Fatal(err)
	}

	// db --start
	db, err := sql.Open("postgres", "host=postgres port=5432 user=uDB password=pDB! dbname=desafio21 sslmode=disable")
	if err != nil {
		log.Panicf("failed to open database: %v", err)
	}
	defer db.Close()

	orderDb := database.NewOrderRepository(db)
	itemDb := database.NewItemRepository(db)
	// db --end

	g := new(errgroup.Group)
	g.Go(func() error { return StartGrpcServer(orderDb, itemDb, listenerGrpc) })
	g.Go(func() error { return StartGraphQl(orderDb, itemDb, listenerHttp) })

	fmt.Printf("running server: %v \n", listenerGrpc.Addr())
	fmt.Printf("running server: %v \n", listenerHttp.Addr())

	g.Wait()
}

func StartGrpcServer(orderDb *database.OrderRepository, itemDb *database.ItemRepository, l net.Listener) error {
	orderService := service.NewOrderService(*orderDb, *itemDb)

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	return grpcServer.Serve(l)
}

func StartGraphQl(orderDb *database.OrderRepository, itemDb *database.ItemRepository, l net.Listener) error {

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		OrderDB: orderDb,
		ItemDB:  itemDb,
	}}))

	mux := http.NewServeMux()
	mux.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		LisOrdersHandler(w, r, orderDb, itemDb)
	})
	mux.HandleFunc("/order-create", func(w http.ResponseWriter, r *http.Request) {
		CreateOrderHandler(w, r, orderDb, itemDb)
	})
	mux.HandleFunc("/order-add-item", func(w http.ResponseWriter, r *http.Request) {
		AddItemOrderHandler(w, r, orderDb, itemDb)
	})
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	httpServer := &http.Server{Handler: mux}
	return httpServer.Serve(l)
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request, orderDb *database.OrderRepository, itemDb *database.ItemRepository) {

	if r.URL.Path != "/order-create" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var payload entity.Order

	dataRead, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err = json.Unmarshal(dataRead, &payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usecase := usecase.NewOrderUseCase(orderDb, itemDb)
	response, err := usecase.Create(payload.Data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)

}

func AddItemOrderHandler(w http.ResponseWriter, r *http.Request, orderDb *database.OrderRepository, itemDb *database.ItemRepository) {

	if r.URL.Path != "/order-add-item" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var payload entity.Item

	dataRead, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//defer body.Close()

	err = json.Unmarshal(dataRead, &payload)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usecase := usecase.NewOrderUseCase(orderDb, itemDb)
	response, err := usecase.AddItem(payload.OrderId, payload.Description, payload.Qty, payload.Value)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)

}

func LisOrdersHandler(w http.ResponseWriter, r *http.Request, orderDb *database.OrderRepository, itemDb *database.ItemRepository) {

	if r.URL.Path != "/order" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	usecase := usecase.NewOrderUseCase(orderDb, itemDb)
	response, err := usecase.GetOrders()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)

}

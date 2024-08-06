package service

import (
	"context"
	"io"

	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/database"
	"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/pb"
	//"github.com/taranttini/study/go/pos-go-expert/fundacao-21-desafio/internal/infra/pb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	OrderDB database.Order
}

func NewOrderService(orderDB database.Order) *OrderService {
	return &OrderService{
		OrderDB: orderDB,
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.Order, error) {
	order, err := o.OrderDB.Create(in.Data)
	if err != nil {
		return nil, err
	}
	orderResponse := &pb.Order{
		Id:   order.Id,
		Data: order.Data,
	}

	return orderResponse, nil
	//return nil, status.Errorf(codes.Unavailable, "method CreateOrder not implemented")
}

// list order
func (o *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrderList, error) {
	orders, err := o.OrderDB.FindAll()
	if err != nil {
		return nil, err
	}

	var ordersResponse []*pb.Order

	for _, order := range orders {
		orderResponse := &pb.Order{
			Id:   order.Id,
			Data: order.Data,
		}

		ordersResponse = append(ordersResponse, orderResponse)
	}

	return &pb.OrderList{Orders: ordersResponse}, nil
}

func (o *OrderService) GetOrder(ctx context.Context, in *pb.OrderGetRequest) (*pb.Order, error) {
	order, err := o.OrderDB.FindByItemId(in.Id)
	if err != nil {
		return nil, err
	}

	orderResponse := &pb.Order{
		Id:   order.Id,
		Data: order.Data,
	}

	return orderResponse, nil
}

func (o *OrderService) CreateOrderStream(stream pb.OrderService_CreateOrderStreamServer) error {
	orders := &pb.OrderList{}

	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(orders)
		}
		if err != nil {
			return err
		}

		orderResult, err := o.OrderDB.Create(order.Data)
		if err != nil {
			return err
		}

		orders.Orders = append(orders.Orders, &pb.Order{
			Id:   orderResult.Id,
			Data: orderResult.Data,
		})
	}
}

func (c *OrderService) CreateOrderStreamBidirectional(stream pb.OrderService_CreateOrderStreamBidirectionalServer) error {
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		orderResult, err := c.OrderDB.Create(order.Data)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.Order{
			Id:   orderResult.Id,
			Data: orderResult.Data,
		})

		if err != nil {
			return err
		}
	}
}

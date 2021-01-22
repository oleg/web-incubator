package main

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "productinfo/service/ecommerce"
)

type server struct {
	productMap map[string]*pb.Product
	pb.UnimplementedProductInfoServer
}

func (s *server) repository() map[string]*pb.Product {
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	return s.productMap
}

func (s *server) AddProduct(ctx context.Context, product *pb.Product) (*pb.ProductID, error) {
	id := uuid.New().String()
	product.Id = id
	s.repository()[id] = product
	return &pb.ProductID{Value: id}, nil
}

func (s *server) GetProduct(ctx context.Context, productId *pb.ProductID) (*pb.Product, error) {
	product, exists := s.repository()[productId.Value]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Product with id %s doesn't exist", productId.Value)
	}
	return product, nil
}

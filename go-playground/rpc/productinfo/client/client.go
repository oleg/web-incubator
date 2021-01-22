package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "productinfo/client/ecommerce"
	"time"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()

	c := pb.NewProductInfoClient(conn)

	name := "simple 10"
	description := "simple product version 10"
	price := float32(1050.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddProduct(ctx, &pb.Product{
		Name:        name,
		Description: description,
		Price:       price,
	})
	if err != nil {
		log.Fatalf("Could not add product %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(ctx, r)
	if err != nil {
		log.Fatalf("Could not get product %v", err)
	}
	log.Printf("Product: %v", product)
}

package main

import (
	"context"
	"log"
	"net"

	pb "microservices-app/consignment-service/proto-consignment"

)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Consignmnet) (*pb.Consignment,error)
}

type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated:= append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

type service struct {
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {

	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func main() {
	repo := &Repository{}

	lis,err:= net.Listen("tcp",port)
	if err != nil{
		log.Fatalf("failed to listen: %v",err)
	}
	s := grpc.NewServer()

	pb.RegisterShippingServiceServer(s, &service{repo})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve :%v",err)
	}
}
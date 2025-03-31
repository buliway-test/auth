package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit/v7"
	user_v1 "github.com/buliway-test/auth/pkg/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

type server struct {
	user_v1.UnimplementedUserV1ServiceServer
}

func (s *server) Create(ctx context.Context, req *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	log.Printf("User created:\nName: %v\nEmail: %v\nPass: %v\nPassConfirm: %v\nRole: %v\n\n",
		req.Info.GetName(), req.Info.GetEmail(), req.Info.GetPassword(),
		req.Info.GetPasswordConfirm(), req.Info.GetRole())

	response := &user_v1.CreateResponse{
		Id: gofakeit.Int64(),
	}
	return response, nil
}

func (s *server) Get(ctx context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	log.Printf("User id: %d\n\n", req.GetId())

	response := &user_v1.GetResponse{
		Info: &user_v1.UserResponse{
			Id:        req.GetId(),
			Name:      gofakeit.Name(),
			Email:     gofakeit.Email(),
			Role:      user_v1.Role_ROLE_USER,
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}
	return response, nil
}

func (s *server) Update(ctx context.Context, req *user_v1.UpdateRequest) (*user_v1.UpdateResponse, error) {
	log.Printf("User updated:\nID: %v\nName: %v\nEmail: %v\n\n",
		req.GetId(), req.GetName(), req.GetEmail())

	return &user_v1.UpdateResponse{}, nil
}

func (s *server) Delete(cto context.Context, req *user_v1.DeleteRequest) (*user_v1.DeleteResponse, error) {
	log.Printf("User id: %d\n\n", req.GetId())

	return &user_v1.DeleteResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	user_v1.RegisterUserV1ServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

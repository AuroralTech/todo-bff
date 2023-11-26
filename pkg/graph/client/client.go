package client

import pb "github.com/AuroralTech/todo-bff/pkg/grpc/generated"

type Client struct {
	TodoClient pb.TodoServiceClient
}

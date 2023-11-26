package graph

import (
	pb "github.com/AuroralTech/todo-bff/pkg/grpc/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoClient pb.TodoServiceClient
}

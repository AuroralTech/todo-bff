package graph

import "github.com/AuroralTech/todo-bff/pkg/graph/client"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client *client.Client
}

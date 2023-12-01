package client

import (
	"context"
	"strconv"

	"github.com/AuroralTech/todo-bff/pkg/graph/generated/model"
	pb "github.com/AuroralTech/todo-bff/pkg/grpc/generated"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Client) AddTodo(ctx context.Context, input model.TodoItemInput) (*model.TodoItem, error) {
	ctx = SetTokenMetadata(ctx)
	resp, err := c.TodoClient.AddTodo(ctx, &pb.TodoItem{
		Task:        input.Task,
		IsCompleted: input.IsCompleted,
	})
	if err != nil {
		return nil, err
	}

	return &model.TodoItem{
		ID:          strconv.FormatUint(resp.Id, 10),
		Task:        resp.Task,
		IsCompleted: resp.IsCompleted,
	}, nil
}

func (c *Client) UpdateTodoStatus(ctx context.Context, input *model.UpdateTodoStatusInput) (*model.UpdateTodoStatusResponse, error) {
	resp, err := c.TodoClient.UpdateTodoStatus(ctx, &pb.UpdateTodoStatusRequest{
		Id:          input.ID,
		IsCompleted: input.IsCompleted,
	})
	if err != nil {
		return nil, err
	}
	return &model.UpdateTodoStatusResponse{
		Success: resp.Success,
	}, nil
}

func (c *Client) DeleteTodoItem(ctx context.Context, input *model.DeleteTodoByIDInput) (*model.DeleteTodoByIDResponse, error) {
	resp, err := c.TodoClient.DeleteTodoById(ctx, &pb.DeleteTodoByIdRequest{
		Id: input.ID,
	})
	if err != nil {
		return nil, err
	}

	return &model.DeleteTodoByIDResponse{
		Success: resp.Success,
	}, nil
}

func (c *Client) TodoList(ctx context.Context) (*model.TodoList, error) {
	resp, err := c.TodoClient.GetTodoList(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	items := func() []*model.TodoItem {
		items := make([]*model.TodoItem, len(resp.Items))
		for i, item := range resp.Items {
			items[i] = &model.TodoItem{
				ID:          strconv.FormatUint(item.Id, 10),
				Task:        item.Task,
				IsCompleted: item.IsCompleted,
			}
		}
		return items
	}()

	return &model.TodoList{
		Items: items,
	}, nil
}

package client

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/metadata"
)

func SetTokenMetadata(ctx context.Context) context.Context {
	request := ctx.Value("httpRequest").(*http.Request)
	token := request.Header.Get("Authorization")
	fmt.Println(token)
	md := metadata.Pairs("Authorization", token)
	newCtx := metadata.NewOutgoingContext(context.Background(), md)

	return newCtx
}

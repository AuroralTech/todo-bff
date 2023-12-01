package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AuroralTech/todo-bff/pkg/graph/client"
	"github.com/rs/cors"

	graph "github.com/AuroralTech/todo-bff/pkg/graph/generated"
	pb "github.com/AuroralTech/todo-bff/pkg/grpc/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func requestContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "httpRequest", r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	// 1.gRPCサーバーのエンドポイントを環境変数から取得
	grpcEndpoint := os.Getenv("GRPC_ENDPOINT")
	if grpcEndpoint == "" {
		log.Fatal("GRPC_ENDPOINT environment variable is not set")
	}

	// 2.gRPCクライアントの接続
	conn, err := grpc.Dial(grpcEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// gRPCクライアントの作成
	todoClient := pb.NewTodoServiceClient(conn)

	// 4.GraphQLハンドラの作成
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
			Client: &client.Client{
				TodoClient: todoClient,
			},
		}}),
	)

	// 5.HTTPサーバーの設定
	http.Handle("/graphql", requestContextMiddleware(srv))
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	// CORS設定の追加
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // すべてのオリジンを許可
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})
	handler := c.Handler(http.DefaultServeMux)
	// 6.サーバーのポートを設定
	port := 4000

	// 7.HTTPサーバーの起動
	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), handler); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}

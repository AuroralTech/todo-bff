package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AuroralTech/todo-bff/pkg/graph"
	"github.com/AuroralTech/todo-bff/pkg/graph/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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

	// 3.GraphQLスキーマの定義

	// 4.GraphQLハンドラの作成
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}),
	)

	// 5.HTTPサーバーの設定
	http.Handle("/graphql", srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	// 6.サーバーのポートを設定
	port := 4000

	// 7.HTTPサーバーの起動
	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}

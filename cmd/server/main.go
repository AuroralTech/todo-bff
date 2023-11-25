package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
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
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new GraphQL schema, error: %v", err)
	}

	// 4.GraphQLハンドラの作成
	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		Playground: true,
	})

	// 5.HTTPサーバーの設定
	http.Handle("/graphql", h)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("BFF Server is running"))
	})

	// 6.サーバーのポートを設定
	port := 4000

	// 7.HTTPサーバーの起動
	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}

package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

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

	// 3.HTTPサーバーの設定
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("BFF Server is running"))
	})

	// 4.サーバーのポートを設定
	port := 4000

	// 5.HTTPサーバーの起動
	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}

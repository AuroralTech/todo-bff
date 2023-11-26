# gqlgenを実行する
.PHONY: gql-gen
gql-gen:
	go get github.com/99designs/gqlgen@v0.17.40 && go run github.com/99designs/gqlgen generate

# protoからpbを生成する
.PHONY: protoc-gen
protoc-gen:
	buf generate

# protoのlintを実行する
.PHONY: protoc-lint
protoc-lint:
	buf lint

# buf.lockファイルのアップデートをする
.PHONY: protoc-update
protoc-update:
	buf mod update

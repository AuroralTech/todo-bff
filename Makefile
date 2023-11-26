# gqlgenを実行する
.PHONY: gql-gen
gql-gen:
	gqlgen generate

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

# ネットワークの作成(ローカルでホストPCから1回だけ実行する)
.PHONY: create-network
create-network:
	docker network create bff_grpc_network

# gql-gen:
# 	go run github.com/99designs/gqlgen generate

# # factoryからprotoを取得する
# # GOPRIVATEをセットしてから実行すること
# .PHONY: factory-get
# factory-get:
# 	go get github.com/panforyou/factory/sdk/factory/go/factory_pb@$(BRANCH_NAME)

# gqlgenを実行する
.PHONY: gql-gen
gql-gen:
	go get github.com/99designs/gqlgen@v0.17.40 && go run github.com/99designs/gqlgen generate

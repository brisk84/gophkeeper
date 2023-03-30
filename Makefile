# generate:
# 	oapi-codegen --package=gophserver --generate types,gorilla -o api/gophserver/gophkeeper-srv.gen.go api/gophkeeper.yaml
# 	oapi-codegen --package=gophclient --generate types,client -o api/gophclient/gophkeeper-cli.gen.go api/gophkeeper.yaml

generate:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/gophkeeper.proto

cert:
	cd cert; ./gen.sh; cd ..

run:
	go run ./cmd/server/main.go -a :8080

.PHONY: cert
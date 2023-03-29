generate:
	oapi-codegen --package=gophserver --generate types,gorilla -o api/gophserver/gophkeeper-srv.gen.go api/gophkeeper.yaml
	oapi-codegen --package=gophclient --generate types,client -o api/gophclient/gophkeeper-cli.gen.go api/gophkeeper.yaml

run:
	go run ./cmd/server/main.go -a :8080

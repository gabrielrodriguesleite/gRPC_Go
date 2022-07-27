generate:
	@echo "\n ----- Building proto -----\n\c"
	@set PATH="$PATH:$(go env GOPATH)/bin" && protoc --proto_path=.  ./helloworld/helloworld.proto --go_out=.. --go-grpc_out=..

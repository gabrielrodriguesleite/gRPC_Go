generate:
	@echo "\n ----- Building proto -----\n\c"
	@protoc --proto_path=.  ./helloworld/helloworld.proto --go_out=.. --go-grpc_out=.. && echo OK
	@echo "\n ----- Saindo ----- \n\c"

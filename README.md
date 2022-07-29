# gRPC_Go

Um repo pra treinar conceitos de gRPC Protobuf em Golang

## Instalação do protoc

```sh
# instalação do protoc
PROTOC_ZIP=protoc-3.14.0-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP

cd /usr/local/bin/
sudo chmod +rx protoc

# instalação do plugin para golang
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

###### Referência

https://google.github.io/proto-lens/installing-protoc.html

## Geração dos arquivos Go com protoc

O arquivo `Makefile` contém utilidades que facilitam o desenvolvimento

Para gerar os arquivos Go com base no protobuffer definido execute:

`make generate`

###### EM CASO DE ERRO

Inclua a seguinte linha no final do arquivo `~/.bashrc` ou `~/.zshrc`

```sh
# PROTO GO plugins
export PATH=$PATH:$(go env GOPATH)/bin
```

Para carregar as configurações refaça o login ou execute `source ~/.bashrc` ou `source ~/.zshrc`

## Executar o Servidor

Em um terminal rode:

`go run server/main.go`

A saída padrão deve exibir:

```sh
2022/07/27 15:33:17 server listening at [::]:50051
```

Ao executar o cliente, o terminal do servidor imprime:

```sh
2022/07/27 15:33:17 server listening at [::]:50051
2022/07/27 15:33:25 Received: world
2022/07/27 15:33:36 Received: Gabriel
```

## Executar o Cliente

Em outro terminal rode:

`go run cliente/main.go`

Saída: 

```sh
2022/07/27 15:33:25 Greeting: Hello world
```

Você pode expecificar um nome com a `flag -name`

`go run cliente/main.go -name Gabriel`

Saída: 

```sh
2022/07/27 15:33:36 Greeting: Hello Gabriel
```

###### References
https://adevait.com/go/transcoding-of-http-json-to-grpc-using-go
https://github.com/felipeagger/gRPC/
https://www.youtube.com/watch?v=LuS59XHdKG8

###### Protoc generate command
```sh
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  helloworld/helloworld.proto
```

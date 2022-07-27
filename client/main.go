package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "GRPC_GO/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	// flags passadas pela linha de comando ao rodar o cliente
	//type(name, value (default?), string)
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {

	flag.Parse()

	// configura a conexão com o servidor
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// encerra a conexão ao terminar a execução de main
	defer conn.Close()

	// conecta com o servidor
	c := pb.NewGreeterClient(conn)

	// contata o servidor com timeout de 1s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// cancela a ligação ao sair da main
	defer cancel()

	// realiza o request do serviço
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// antes de finalizar imprime a mensagem
	log.Printf("Greeting: %s", r.GetMessage())

}

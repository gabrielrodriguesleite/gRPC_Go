package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "GRPC_GO/helloworld"

	"google.golang.org/grpc"
)

var (
	// flag implementa facilidade de capturar dados da linha de comando ao executar
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	// imprime o log na saída padrão do server
	log.Printf("Received: %v", in.GetName())

	// roda a resposta da requisição
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil

}

// ----- Ampliação INICIO -----
func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	//imprime um log na saída padrão do server
	log.Printf("Received: %v", in.GetName())

	// roda a resposta da nova requisição
	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}

// ----- Ampliação FIM -----

func main() {

	flag.Parse()

	// estabelece a comunicação tcp
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// estancia o servidor grpc
	s := grpc.NewServer()

	// registra o serviço
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	// inicia o servidor grpc
	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

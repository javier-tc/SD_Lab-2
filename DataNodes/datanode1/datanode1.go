package main

import (
	"context"
	"log"
	"net"

	pb "github.com/javier-tp/SD_Lab-2/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedComServServer
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func (s *server) OMSDataNode(ctx context.Context, msg *pb.IdPersona) (*pb.IdPersona, error) {
	return msg, nil
}

func main() {
	port := ":50051"
	listener, err := net.Listen("tcp", port) //conexion sincrona
	failOnError(err, "La conexion no se pudo crear")

	serv := grpc.NewServer()
	pb.RegisterComServServer(serv, &server{})

	if err == serv.Serve(listener) {
		failOnError(err, "Fallo al iniciar server")
	}

}

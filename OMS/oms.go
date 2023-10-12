package main

import (
	"context"
	"fmt"
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

func (s *server) ContOMS(ctx context.Context, msg *pb.EstadoPersona) (*pb.EstadoPersona, error) {
	fmt.Printf("Nombre: %s | Apellido: %s | Estado: %s \n", msg.Nombre, msg.Apellido, msg.Estado)
	return msg, nil
}

func envioDatanode(port string, id string, nombre string, apellido string) {
	//envio a datanode
	//port
	id_persona := id
	nombre_persona := nombre
	apellido_persona := apellido

	conn, err := grpc.Dial(port, grpc.WithInsecure()) //conectamos con el servidor
	failOnError(err, "No se pudo conectar con el servidor")
	service_client := pb.NewComServClient(conn)
	_, err = service_client.OMSDataNode(context.Background(), //enviamos mensaje con id de persona
		&pb.IdPersona{
			Id:       id_persona,
			Nombre:   nombre_persona,
			Apellido: apellido_persona,
		})
	conn.Close()
}

func main() {
	port := ":50051"
	listener, err := net.Listen("tcp", port) //conexion sincrona
	failOnError(err, "La conexion no se pudo crear")

	serv := grpc.NewServer()
	pb.RegisterComServServer(serv, &server{})

	fmt.Printf("OMS escuchando en %s\n", port)

	if err == serv.Serve(listener) {
		failOnError(err, "Fallo al iniciar server")
	}

	//envioDatanode(port)
}

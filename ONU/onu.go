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

func consultaNamenode(port string, estado string) { //FALTA IMPLEMENTAR LA FUNCION DE RESPUESTA EN OMS.go
	conn, err := grpc.Dial(port, grpc.WithInsecure()) //conectamos con el servidor
	failOnError(err, "No se pudo conectar con el servidor")
	service_client := pb.NewComServClient(conn)

	_, err = service_client.ConsultaOMS(context.Background(), //enviamos mensaje
		&pb.EstadoPersona{
			Nombre:   "",
			Apellido: "",
			Estado:   estado,
		})
	conn.Close()
}

func (s *server) OMSDataNode(ctx context.Context, msg *pb.IdPersona) (*pb.IdPersona, error) {
	fmt.Printf("Solicitud de OMS recibida, mensaje recibido: %s %s\n", msg.Nombre, msg.Apellido)
	return msg, nil
}

func connection() {
	port := ":50053"
	listener, err := net.Listen("tcp", port) //conexion sincrona
	failOnError(err, "La conexion no se pudo crear")

	serv := grpc.NewServer()
	pb.RegisterComServServer(serv, &server{})

	//fmt.Printf("ONU escuchando en %s\n", port)

	if err == serv.Serve(listener) {
		failOnError(err, "Fallo al iniciar server")
	}
}

func consulta() {
	var entrada string
	fmt.Print("Consulta por <Muerto> o <Infectado>:\n")
	_, err := fmt.Scanln(&entrada)
	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return
	}

	puerto_oms := "dist023:50050"
	//consulta
	consultaNamenode(puerto_oms, entrada)
}

func main() {

	go connection()
	for {
		consulta()
	}

}

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

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

func writeData(id string, nombre string, apellido string) {
	line := id + " " + nombre + " " + apellido + "\n"
	// Abre o crea el archivo "DATA.txt" en modo de apertura y escritura (append)
	archivo, err := os.OpenFile("DATA.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()

	_, err = archivo.WriteString(line)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
	}
}

func obtenerNombreApellidoPorID(archivo string, idBuscado string) (string, string, error) {
	file, err := os.Open(archivo)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		campos := strings.Fields(linea) // Divide la línea en palabras
		if len(campos) < 3 {
			continue
		}

		id := campos[0]

		if id == idBuscado {
			return campos[1], campos[2], nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", "", err
	}

	return "", "", fmt.Errorf("ID no encontrado")
}

// funcion en que recibe mensajes de la oms
func (s *server) OMSDataNode(ctx context.Context, msg *pb.IdPersona) (*pb.IdPersona, error) {
	id := strconv.FormatInt(msg.Id, 10)
	nombre := msg.Nombre
	apellido := msg.Apellido
	//almacenar msg en data.txt
	writeData(id, nombre, apellido)
	return msg, nil
}

func (s *server) ConsultaDN(ctx context.Context, msg *pb.IdPersona) (*pb.IdPersona, error) {
	id := strconv.FormatInt(msg.Id, 10)
	nombre, apellido, _ := obtenerNombreApellidoPorID("DATA.txt", id)
	fmt.Printf("Solicitud de NameNode recibida, mensaje enviado: %s %s\n", nombre, apellido)
	response := &pb.IdPersona{
		Id:       0,
		Nombre:   nombre,
		Apellido: apellido,
	}
	return response, nil
}

func main() {
	port := ":50051"
	listener, err := net.Listen("tcp", port) //conexion sincrona
	failOnError(err, "La conexion no se pudo crear")

	serv := grpc.NewServer()
	pb.RegisterComServServer(serv, &server{})

	//fmt.Printf("DN1 escuchando en %s\n", port)

	if err == serv.Serve(listener) {
		failOnError(err, "Fallo al iniciar server")
	}

}

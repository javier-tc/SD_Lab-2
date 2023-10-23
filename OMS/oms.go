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

var idCounter int

func generarIDUnico() int {
	idCounter++
	return idCounter
}

func writeData(dn string, id int, estado string) {
	line := strconv.Itoa(id) + " " + dn + " " + estado + "\n"
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

type Registro struct {
	ID        int
	NDatanode int
}

type Persona struct {
	Nombre   string
	Apellido string
}

var personas []Persona

func buscarRegistrosPorEstado(archivo, estadoBuscado string) ([]Registro, error) {
	file, err := os.Open(archivo)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	registros := make([]Registro, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		campos := strings.Fields(linea) // Divide la línea en palabras
		if len(campos) != 3 {
			continue
		}

		id, err := strconv.Atoi(campos[0])
		if err != nil {
			continue
		}

		ndatanode, err := strconv.Atoi(campos[1])
		if err != nil {
			continue
		}

		estado := campos[2]
		if estado == estadoBuscado {
			registros = append(registros, Registro{ID: id, NDatanode: ndatanode})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return registros, nil
}

func (s *server) ConsultaOMS(ctx context.Context, msg *pb.EstadoPersona) (*pb.EstadoPersona, error) {
	estado := msg.Estado
	registros, _ := buscarRegistrosPorEstado("DATA.txt", estado)

	for _, registro := range registros {
		//puerto datanodes
		puerto := "dist02" + strconv.Itoa(registro.NDatanode) + ":5005" + strconv.Itoa(registro.NDatanode)
		nombre, apellido := envioIdReturnNombreApellido(puerto, registro.ID)
		//func para enviar a onu nombre y apellido
		persona := Persona{
			Nombre:   nombre,
			Apellido: apellido,
		}
		// Agregar la persona a la slice
		personas = append(personas, persona)
	}
	//puerto onu
	puerto_onu := "dist024:50053"
	for _, p := range personas {
		fmt.Printf("Solicitud de ONU recibida, mensaje enviado: %s %s\n", p.Nombre, p.Apellido)
		envioIdNombreApellido(puerto_onu, 0, p.Nombre, p.Apellido)
	}

	return msg, nil
}

// consulta por id a datanode
func envioIdReturnNombreApellido(port string, id int) (string, string) {
	conn, err := grpc.Dial(port, grpc.WithInsecure()) //conectamos con el servidor
	failOnError(err, "No se pudo conectar con el servidor")
	service_client := pb.NewComServClient(conn)
	response, err := service_client.ConsultaDN(context.Background(), //enviamos mensaje con id de persona
		&pb.IdPersona{
			Id:       int64(id),
			Nombre:   "",
			Apellido: "",
		})
	conn.Close()

	nombre := response.Nombre
	apellido := response.Apellido
	return nombre, apellido
}

// envio a datanode | onu
func envioIdNombreApellido(port string, id int, nombre string, apellido string) {
	conn, err := grpc.Dial(port, grpc.WithInsecure()) //conectamos con el servidor
	failOnError(err, "No se pudo conectar con el servidor")
	service_client := pb.NewComServClient(conn)
	_, err = service_client.OMSDataNode(context.Background(), //enviamos mensaje con id de persona
		&pb.IdPersona{
			Id:       int64(id),
			Nombre:   nombre,
			Apellido: apellido,
		})
	conn.Close()
}

func (s *server) ContOMS(ctx context.Context, msg *pb.EstadoPersona) (*pb.EstadoPersona, error) {
	id := generarIDUnico()
	dn := "0"
	nombre := msg.Nombre
	apellido := msg.Apellido
	inicial_ap := apellido[0]
	estado := msg.Estado

	if inicial_ap >= 'A' && inicial_ap <= 'M' {
		dn = "1"
	} else {
		dn = "2"
	}

	writeData(dn, id, estado)
	puerto_dn := "dist02" + dn + ":5005" + dn
	//envio info a datanode -> "<id> <nombre> <apellido>"
	envioIdNombreApellido(puerto_dn, id, nombre, apellido)
	return msg, nil
}

func main() {
	port := ":50050"
	listener, err := net.Listen("tcp", port) //conexion sincrona
	failOnError(err, "La conexion no se pudo crear")

	serv := grpc.NewServer()
	pb.RegisterComServServer(serv, &server{})

	fmt.Printf("OMS escuchando en %s\n", port)

	if err == serv.Serve(listener) {
		failOnError(err, "Fallo al iniciar server")
	}
}

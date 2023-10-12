package main

import (
	"bufio"
	"context"
	"log"
	"math/rand"
	"os"
	"strings"

	pb "github.com/javier-tp/SD_Lab-2/proto"
	"google.golang.org/grpc"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func envioOMS(port string, nombre_persona string, apellido_persona string, estado_persona string) {
	conn, err := grpc.Dial(port, grpc.WithInsecure()) //conectamos con el servidor
	failOnError(err, "No se pudo conectar con el servidor")
	service_client := pb.NewComServClient(conn)

	println("enviando mensaje...", nombre_persona, apellido_persona, estado_persona)

	_, err = service_client.ContOMS(context.Background(), //enviamos mensaje con estado de persona
		&pb.EstadoPersona{
			Nombre:   nombre_persona,
			Apellido: apellido_persona,
			Estado:   estado_persona,
		})
	conn.Close()
}

func estadoPersona(port string) {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ruta_txt := currentDir + "/../data.txt"
	file, ferr := os.Open(ruta_txt) //Se abre el archivo que contiene nombres
	if ferr != nil {
		panic(ferr)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var estado string
	var nombre string
	var apellido string

	for scanner.Scan() {
		line := scanner.Text()
		nombre = strings.Split(line, " ")[0]
		apellido = strings.Split(line, " ")[1]

		probabilidad := rand.Float64()

		// Asigna "Infectado" o "Muerto" según la probabilidad
		if probabilidad < 0.55 {
			estado = "Infectado"
		} else {
			estado = "Muerto"
		}
		envioOMS(port, nombre, apellido, estado)
	}
}

func main() {
	//conexion grpc
	//s := regToIp[k]       //servidor al que conectarse
	//port := s + comp_port //puerto

	//persona
	port := ":50051"
	estadoPersona(port)
}

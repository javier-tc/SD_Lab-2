package main

import (
	"bufio"
	"context"
	"log"
	"os"

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
	_, err = service_client.ContOMS(context.Background(), //enviamos mensaje con estado de persona
		&pb.EstadoPersona{
			Nombre:   nombre_persona,
			Apellido: apellido_persona,
			Estado:   estado_persona,
		})
	conn.Close()
}

func estadoPersona() {
	file, ferr := os.Open("./Continentes/data.txt") //Se abre el archivo que contiene nombres
	if ferr != nil {
		panic(ferr)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		println(line)
	}
}

func main() {
	//conexion grpc
	//s := regToIp[k]       //servidor al que conectarse
	//port := s + comp_port //puerto

	//persona
	/*
		nombre_persona := "nombre"
		apellido_persona := "apellido"
		estado_persona := "estado"

		port := "50051"
	*/
	estadoPersona()
	//envioOMS(port, nombre_persona, apellido_persona, estado_persona)
}

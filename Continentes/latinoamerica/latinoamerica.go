package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

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

	fmt.Printf("Estado enviado: %s %s %s \n", nombre_persona, apellido_persona, estado_persona)

	_, err = service_client.ContOMS(context.Background(), //enviamos mensaje con estado de persona
		&pb.EstadoPersona{
			Nombre:   nombre_persona,
			Apellido: apellido_persona,
			Estado:   estado_persona,
		})
	conn.Close()
}

func obtenerLineaAleatoria(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	rand.Seed(time.Now().UnixNano())

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	if len(lines) == 0 {
		return []string{}, fmt.Errorf("El archivo está vacío")
	}

	return lines, nil
}

func estadoPersona(port string) {
	var estado string
	var nombre string
	var apellido string
	var lines []string

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ruta_txt := currentDir + "/../names.txt"
	lines, ferr := obtenerLineaAleatoria(ruta_txt)
	if ferr != nil {
		panic(ferr)
	}

	len_txt := len(lines)
	count := 0

	for n := 0; n < len_txt; n++ {
		line := lines[rand.Intn(len_txt)]
		nombre = strings.Split(line, " ")[0]
		apellido = strings.Split(line, " ")[1]

		probabilidad := rand.Float64()

		// Asigna "Infectado" o "Muerto" según la probabilidad
		if probabilidad < 0.55 {
			estado = "Infectado"
		} else {
			estado = "Muerto"
		}

		if count > 4 {
			time.Sleep(3 * time.Second)
			envioOMS(port, nombre, apellido, estado)
		} else {
			envioOMS(port, nombre, apellido, estado)
		}
		count += 1
	}
}

func main() {
	//conexion grpc
	port := "dist023:50050"
	estadoPersona(port)
}

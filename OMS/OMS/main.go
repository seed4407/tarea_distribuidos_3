/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"net" 
    "os"
	"bufio"
	// "math/rand"
	"sync"
	"time"
	"strconv"
	"google.golang.org/grpc"
	// "github.com/streadway/amqp"
	pb "github.com/seed4407/Tarea_distribuidos_2/proto"
)

var (
	port = flag.Int("port", 80, "The server port")
)

var id_datos int
var err error
var aux string
var lock = &sync.RWMutex{}

type server struct {
	pb.UnimplementedNameNodeServer
}

func enviarDatosAlDataNode(dataNodeAddr string, id string, informe *pb.Datos) {
    conn, err := grpc.Dial(dataNodeAddr, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("No se pudo conectar al DataNode: %v", err)
    }
    defer conn.Close()

	if dataNodeAddr == "10.6.46.109:8080"{
		fmt.Println("Mensaje enviado a DataNode1",id, informe)
	} else{
	    fmt.Println("Mensaje enviado a DataNode2",id, informe)
	}
    client := pb.NewDataNodeClient(conn)

    _, err = client.RegistrarNombre(context.Background(), &pb.Registro{
        Id:       id,
        Nombre:   informe.Nombre,
        Apellido: informe.Apellido,
    })
    if err != nil {
        log.Fatalf("Error al registrar nombre en el DataNode: %v", err)
    }
}

func (s *server) Recepcion_Info(ctx context.Context, in *pb.Datos) (*pb.Recepcion, error) {
	lock.Lock()
    defer lock.Unlock()
	var datos_persona = in.GetApellido()
	var nodo int
	filePath := "/app/Data.txt"

	fmt.Println("Mensaje recibido de Continente",in)
	if datos_persona[0] <= 77 {
		enviarDatosAlDataNode("10.6.46.109:8080",strconv.Itoa(id_datos),in)
		nodo = 1
	} else{
		enviarDatosAlDataNode("10.6.46.110:8080",strconv.Itoa(id_datos),in)
		nodo = 2
	}

	dato_escribir := strconv.Itoa(id_datos) + " " + strconv.Itoa(nodo) + " " +  in.GetEstado()
	id_datos = id_datos + 1

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	_, err = file.WriteString(dato_escribir + "\n")
	if err != nil {
		return nil, err
	}
	
	err = file.Sync()
	if err != nil {
		return nil, err
	}
	file.Close()

	return &pb.Recepcion{Ok:aux}, nil
}

func (s *server) ConsultarNombres(ctx context.Context, in *pb.Estado_Persona) (*pb.Lista_Datos_DataNode, error) {
    var estado_persona = strings.ToUpper(in.Estado)
	fmt.Println("Mensaje recibido de ONU",estado_persona)
	var linea_data string
	var id_1 []string
	var id_2 []string

	filePath := "/app/Data.txt"
	var file *os.File

	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		linea_data = scanner.Text() 
        // fmt.Println(linea_data)
		if strings.Contains(linea_data, estado_persona) {
			if strings.Contains(linea_data, " 1 ") {
				id_1 = append(id_1,strings.Split(linea_data, " ")[0])
			} else {
				id_2 = append(id_2,strings.Split(linea_data, " ")[0])
			}
		}
    }

	fmt.Println("Mensaje enviado a DataNode1",id_1)
	fmt.Println("Mensaje enviado a DataNode2",id_2)
	
	if err = scanner.Err(); err != nil {
        log.Fatal(err)
    }

	file.Close()

	//llamar a funcion para solicitar a dataNode y se retorna lo que devuelva

	conn, err := grpc.Dial("10.6.46.109:8080",grpc.WithInsecure())

	if err != nil {
        log.Fatal(err)
    }

	defer conn.Close()

	datanode := pb.NewDataNodeClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    respuesta1, err := datanode.Solicitud_Info_DataNode(ctx, &pb.Id{ListaId: id_1})
    if err != nil{
        log.Print("No hay respuesta del datanode1")
    }else{
		fmt.Println("Mensaje recibido de DataNode1",respuesta1.Datos)
    }

	conn, err = grpc.Dial("10.6.46.110:8080",grpc.WithInsecure())

	if err != nil {
        log.Fatal(err)
    }

	defer conn.Close()

	datanode = pb.NewDataNodeClient(conn)
    ctx, cancel = context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    respuesta2, err := datanode.Solicitud_Info_DataNode(ctx, &pb.Id{ListaId: id_2})
    if err != nil{
        log.Print("No hay respuesta del Datanode2")
    }else{
		fmt.Println("Mensaje recibido de DataNode2",respuesta2.Datos)
    }

	var respuesta []*pb.Datos_DataNode

	for _, dato := range respuesta1.Datos {
		persona := &pb.Datos_DataNode{
			Nombre:   dato.Nombre,
			Apellido: dato.Apellido,
		}
		respuesta = append(respuesta,persona)
	}

	for _, dato := range respuesta2.Datos {
		persona := &pb.Datos_DataNode{
			Nombre:   dato.Nombre,
			Apellido: dato.Apellido,
		}
		respuesta = append(respuesta,persona)
	}

	return &pb.Lista_Datos_DataNode{Datos:respuesta}, nil
}


func main() {
    // Abrir el archivo en modo lectura
	filePath := "/app/Data.txt"

	var linea_data string
	var id []string
	var estado_persona = "Infectado"
	var file *os.File

	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		linea_data = scanner.Text() 
		if strings.Contains(linea_data, estado_persona) {
			id = append(id,strings.Split(linea_data, " ")[0])
		}
    }
	if(linea_data == ""){
		id_datos = 1
	} else{
		id_datos, err = strconv.Atoi(strings.Split(linea_data, " ")[0])
		id_datos = id_datos + 1
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d",*port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNameNodeServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

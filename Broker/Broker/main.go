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
	// "strings"
	"net" 
    // "os"
	// "bufio"
	"math/rand"
	// "sync"
	"time"
	// "strconv"
	"google.golang.org/grpc"
	// "github.com/streadway/amqp"
	pb "github.com/seed4407/Tarea_distribuidos_3/proto"
)

var (
	port = flag.Int("port", 80, "The server port")
)

// var id_datos int
var err error
var aux string
// var lock = &sync.RWMutex{}

type server struct {
	pb.UnimplementedVanguardiaServer
	pb.UnimplementedInformantesServer
}

func (s *server) GetSoldados(ctx context.Context, in *pb.ParametroSectorBase) (*pb.RetornoCantSoldOscuridadReloj, error) {
	
	rand.Seed(time.Now().UnixNano())
	var numeroAleatorio = rand.Intn(3)
	
	var nombre_sector = in.GetNombreSector()
	var nombre_base = in.GetNombreBase()

	log.Printf("%s",nombre_sector)
	log.Printf("%s",nombre_base)

	if(numeroAleatorio == 0){
		log.Printf("Enviar consulta a server 0")
	} else if(numeroAleatorio == 1){
		log.Printf("Enviar consulta a server 1")
	} else{
		log.Printf("Enviar consulta a server 2")
	}

	// var respuesta *pb.RetornoCantSoldOscuridadReloj

	var respuesta = &pb.RetornoCantSoldOscuridadReloj{
		CantSoldadosOscuridad:   1,
		Reloj: []int32{1, 2, 3},
	}
	// respuesta.CantSoldadosOscuridad = 23
	// respuesta.Reloj = 

	// for _, dato := range respuesta1.Datos {
	// 	persona := &pb.Datos_DataNode{
	// 		Nombre:   dato.Nombre,
	// 		Apellido: dato.Apellido,
	// 	}
	// 	respuesta = append(respuesta,persona)
	// }


	return &pb.RetornoCantSoldOscuridadReloj{CantSoldadosOscuridad:respuesta.CantSoldadosOscuridad ,Reloj:respuesta.Reloj}, nil
}

func (s *server) Obtener_servidor_consultar_aleatorio(ctx context.Context, in *pb.ParamamtroVacio) (*pb.ServidorAConsultar, error) {
	
	rand.Seed(time.Now().UnixNano())
	var numeroAleatorio = rand.Intn(3)

	if(numeroAleatorio == 0){
		log.Printf("Enviar ip server 0")
	} else if(numeroAleatorio == 1){
		log.Printf("Enviar ip server 1")
	} else{
		log.Printf("Enviar ip server 2")
	}

	var respuesta = &pb.ServidorAConsultar{
		Ip:   "100.100.100",
	}

	return &pb.ServidorAConsultar{Ip:respuesta.Ip}, nil
}



func main() {
    // // Abrir el archivo en modo lectura
	// filePath := "/app/Data.txt"

	// var linea_data string
	// var id []string
	// var estado_persona = "Infectado"
	// var file *os.File

	// file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)

	// scanner := bufio.NewScanner(file)
    // for scanner.Scan() {
	// 	linea_data = scanner.Text() 
	// 	if strings.Contains(linea_data, estado_persona) {
	// 		id = append(id,strings.Split(linea_data, " ")[0])
	// 	}
    // }
	// if(linea_data == ""){
	// 	id_datos = 1
	// } else{
	// 	id_datos, err = strconv.Atoi(strings.Split(linea_data, " ")[0])
	// 	id_datos = id_datos + 1
	// }

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d",*port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVanguardiaServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

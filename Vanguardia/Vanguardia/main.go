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
	// "net" 
    // "os"
	// "bufio"
	// "math/rand"
	// "sync"
	// "time"
	// "strconv"
	"google.golang.org/grpc"
	// "github.com/streadway/amqp"
	pb "github.com/seed4407/Tarea_distribuidos_2/proto"
)

var (
	port = flag.Int("port", 80, "The server port")
)



func GetSoldados(nombre_sector string,nombre_base string) (*pb.RetornoCantSoldOscuridadReloj, error){

	conn, err := grpc.Dial("172.17.0.2:80", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al broker: %v", err)
	}
	defer conn.Close()

	client := pb.NewVanguardiaClient(conn)

	res, err := client.GetSoldados(context.Background(), &pb.ParametroSectorBase{
		NombreSector: nombre_sector,
		NombreBase: nombre_base,
	})

	if err != nil {
		log.Fatalf("Error al pedir datos: %v", err)
	} else {
		log.Printf("%d\n", res.CantSoldadosOscuridad)
		log.Printf("%v\n", res.Reloj)
	}

	reloj_almacenado := buscar_coincidencia(nombre_sector)

	for i:=0;i<3;i++{
		if(reloj_almacenado[i]>res.Reloj[i]){
			log.Println("Error de consisntencia, leyendo dato mas antiguo")
		}
	}

	return res, nil
}

func buscar_coincidencia(nombre_sector string) []int32{
	filePath := "/app/Data.txt"

	var linea_data string
	var file *os.File
	var reloj []int32
	
	file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea_data = scanner.Text() 
		if strings.Contains(linea_data, nombre_sector) {
			result := strings.Split(linea_data, " ")
			err := json.Unmarshal([]byte(result[1]), &reloj)
			if err != nil {
				log.Fatal(err)
			}
			return reloj
		}
	}
	return ""
}

func main() {
	var datos []string
	for {
		var sector string
		var base string
		log.Println("Ingrese sector y base de los que se quiere conseguir cantidad de solado (Ej 'sector1 base1'):")
		cant_parametros, err := fmt.Scanln(&sector,&base)
		if err != nil {
			log.Printf("Error al leer la entrada: %v\n", err)
			return
		}

		if cant_parametros == 2 {
			log.Printf("Enviar consulta")
			GetSoldados(sector, base)
		} else {
			if len(datos) < 2 {
				log.Printf("Error: Faltan datos")
			} else if len(datos) > 2 {
				log.Printf("Error: Sobran datos")
			}
		}
	}
}

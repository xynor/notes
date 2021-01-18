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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	wg := sync.WaitGroup{}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	for i := 0; i < 10; i++ {
		go func(k int) {
			wg.Add(1)
			defer wg.Done()
			for {
				r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
				if err != nil {
					log.Printf("could not greet(%d): %v", k, err)
					time.Sleep(1 * time.Second)
				}
				log.Printf("Greeting(%d): %s", k, r.GetMessage())
				time.Sleep(1 * time.Second)
			}
		}(i)
	}
	wg.Wait()
}

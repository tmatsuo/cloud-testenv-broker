/*
Copyright 2014 Google Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package borker implements the broker for cloud emmulators.
package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	broker "cloud-testenv-broker/broker"
	"google.golang.org/grpc/credentials"
	emulators "google/emulators"
	pb "google/protobuf"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "server1.pem", "The TLS cert file")
	keyFile    = flag.String("key_file", "server1.key", "The TLS key file")
	port       = flag.Int("port", 10000, "The server port")
	configFile = flag.String("config_file", "", "The json config file of the Cloud Broker.")
	EMPTY      = &pb.Empty{}
)
var config *broker.Config

func main() {
	log.Printf("Emulator broker starting up...")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v.", err)
	}
	if *configFile != "" {
		config, err = broker.Decode(*configFile)
		if err != nil {
			log.Fatalf("Could not parse config file: %v", err)
		}
	}
	grpcServer := grpc.NewServer()
	server := broker.New()
	emulators.RegisterBrokerServer(grpcServer, server)
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v.", err)
		}
		log.Printf("Gateway listening with TLS on :%d.", *port)
		grpcServer.Serve(creds.NewListener(lis))
	} else {
		log.Printf("Gateway listening on :%d.", *port)
		grpcServer.Serve(lis)
	}
}

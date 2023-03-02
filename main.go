package main

import (
	"log"

	"github.com/pachyderm/pachyderm/src/client"
	v2c "github.com/pachyderm/pachyderm/v2/src/client"
)

func main() {
	pachOne, err := client.NewFromAddress("localhost:80")
	if err != nil {
		log.Fatalf("new pach1 client: %v", err)
	}
	if v, err := pachOne.Version(); err != nil {
		log.Printf("pach1 version: %v", err)
	} else {
		log.Printf("pach1 version: %v", v)
	}

	pachTwo, err := v2c.NewFromURI("grpc://localhost:80")
	if err != nil {
		log.Fatalf("new pach2 client: %v", err)
	}
	if v, err := pachTwo.Version(); err != nil {
		log.Printf("pach2 version: %v", err)
	} else {
		log.Printf("pach2 version: %v", v)
	}
}

package main

import (
	"fmt"

	"github.com/gofrs/uuid"
	//"log"
)

func main() {

	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	// u1 := uuid.Must(uuid.NewV4())
	// fmt.Printf("UUIDv4: %s\n", u1)

	// u2, err := uuid.NewV4()
	// if err != nil {
	// 	log.Fatalf("failed to generate UUID: %v", err)
	// }
	// log.Printf("generated Version 4 UUID %v", u2)

}

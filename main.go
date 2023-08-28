package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Printf("hi with this uuid: %v", uuid.Must(uuid.NewRandom()).String())
}

package main

import (
	"AlvinSamuelsson/chip8go/internal/chip8"
	"fmt"
	"log"
)

func main() {
	cmp, err := chip8.NewPC()
	if err != nil {
		log.Fatalf("\n error creating a new PC: %v \n", err)
	}

	fmt.Printf("%v", cmp.Memory)
}

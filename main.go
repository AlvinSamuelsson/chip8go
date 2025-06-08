package main

import (
	"AlvinSamuelsson/chip8go/internal/chip8"
	"log"

	"github.com/gopxl/pixel/v2/backends/opengl"
)

func main() {
	opengl.Run(run)
}

func run() {
	cmp, err := chip8.NewPC("roms/ibm.ch8")
	if err != nil {
		log.Fatalf("\n error creating a new PC: %v \n", err)
		panic("REE")
	}
	cmp.RunProgram()
}

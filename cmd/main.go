package main

import (
	"boids/parser"
	"boids/sim"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	count, width, height, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		panic(err)
	}

	bound := rl.NewRectangle(0, 0, float32(width), float32(height))
	s := sim.NewSim(width, height, count, &bound)
	s.Run()
}

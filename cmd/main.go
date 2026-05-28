package main

import (
	"boids/sim"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	defCount = 2000
)

func main() {
	w, h := int32(2400), int32(1300)
	bound := rl.NewRectangle(0, 0, float32(w), float32(h))
	s := sim.NewSim(w, h, defCount, &bound)
	s.Run()
}

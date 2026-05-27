package main

import (
	"boids/boid/logic"
	"boids/boid/render"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	numBoids = 200
)

func main() {
	windowHeight, windowWidth := int32(1080), int32(1920)
	rl.InitWindow(windowWidth, windowHeight, "boids")
	defer rl.CloseWindow()
	cam := newCamera3D()

	var frameDelta float32 = 0.0
	rl.SetTargetFPS(60)

	boids := initBoids()
	boidColors := make([]rl.Color, numBoids)
	for i := range boidColors {
		boidColors[i] = rl.NewColor(
			uint8(rand.Intn(256)),
			uint8(rand.Intn(256)),
			uint8(rand.Intn(256)),
			255,
		)
	}

	minBound := rl.NewVector3(-50, -50, -50)
	maxBound := rl.NewVector3(50, 50, 50)
	bounds := rl.NewBoundingBox(minBound, maxBound)

	for !rl.WindowShouldClose() {
		frameDelta = rl.GetFrameTime()

		for _, boid := range boids {
			boid.Update(boids, &bounds, frameDelta)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		rl.DrawFPS(10, 10)

		rl.BeginMode3D(cam)
		for i, boid := range boids {
			render.DrawBoid(boid.GetPos(), boidColors[i])
		}
		render.DrawBounds(&bounds)

		rl.EndMode3D()

		rl.EndDrawing()
	}
}

func newCamera3D() rl.Camera3D {
	return rl.Camera3D{
		Position:   rl.NewVector3(200, 100, 0),
		Target:     rl.NewVector3(0, 0, 0),
		Up:         rl.NewVector3(0, 1, 0),
		Fovy:       45,
		Projection: rl.CameraPerspective,
	}
}

func initBoids() []*logic.Boid {
	boids := make([]*logic.Boid, 0)
	for range numBoids {
		boids = append(
			boids,
			logic.NewBoid(rl.NewVector3(
				rand.Float32()*100-50,
				rand.Float32()*100-50,
				rand.Float32()*100-50,
			)),
		)
	}
	return boids
}

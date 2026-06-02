package sim

import (
	"runtime"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	bgColor   = rl.NewColor(3, 10, 36, 255)
	boidColor = rl.NewColor(209, 192, 155, 255)
)

type Sim struct {
	windowWidth  int32
	windowHeight int32
	boidCount    int
	bound        *rl.Rectangle
	boids        []*Boid
}

func NewSim(w, h int32, count int, bound *rl.Rectangle) *Sim {
	return &Sim{
		windowWidth:  w,
		windowHeight: h,
		boidCount:    count,
		bound:        bound,
		boids:        spawnBoids(count, bound),
	}
}

func (s *Sim) Run() {
	rl.InitWindow(s.windowWidth, s.windowHeight, "Boids")
	defer rl.CloseWindow()

	var df float32 = 0.0
	var accDf float32 = 0.0
	var frames int = 0
	var frameThreshold int = 100
	var maxRad float32 = max(separationRadius, max(alignmentRadius, cohesionRadius))
	var tree *Tree = NewTree(*s.bound)
	var wg sync.WaitGroup
	var numWorkers = runtime.NumCPU()
	var chunkSize = (len(s.boids) + numWorkers - 1) / numWorkers
	for _, b := range s.boids {
		tree.Insert(b)
	}

	for !rl.WindowShouldClose() {
		df = rl.GetFrameTime()
		accDf += df

		if frames%frameThreshold != 0 {
			tree.Clear()
			for _, b := range s.boids {
				tree.Insert(b)
			}

			wg.Add(numWorkers)
			for i := range numWorkers {
				start := i * chunkSize
				end := min(start+chunkSize, len(s.boids))
				go func(chunk []*Boid) {
					defer wg.Done()
					var inRange []*Boid
					for _, b := range chunk {
						inRange = inRange[:0]
						tree.QueryRange(
							rl.NewRectangle(
								b.Position.X-maxRad,
								b.Position.Y-maxRad,
								maxRad*2,
								maxRad*2,
							),
							&inRange,
						)
						b.UpdateDir(inRange, accDf)
					}
				}(s.boids[start:end])
			}
			wg.Wait()
			accDf = 0.0
		}

		rl.BeginDrawing()
		rl.ClearBackground(bgColor)
		for _, b := range s.boids {
			b.Move(s.bound, df)
			renderBoid(b)
		}
		rl.DrawFPS(10, 10)
		rl.EndDrawing()

		frames += 1 % frameThreshold
	}
}

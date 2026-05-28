package sim

import rl "github.com/gen2brain/raylib-go/raylib"

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
		boids:        SpawnBoids(count, bound),
	}
}

func (s *Sim) Run() {
	rl.InitWindow(s.windowWidth, s.windowHeight, "Boids")
	defer rl.CloseWindow()

	var df float32 = 0.0

	for !rl.WindowShouldClose() {
		df = rl.GetFrameTime()

		rl.BeginDrawing()
		rl.ClearBackground(bgColor)

		for _, b := range s.boids {
			b.UpdateDir(s.boids, df)
			b.Move(s.bound, df)
			renderBoid(b)
			//rl.DrawCircleLinesV(b.Position, PerceptionRadius, boidColor)
		}

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}

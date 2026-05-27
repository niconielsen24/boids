package render

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	radius = 1
)

func DrawBoid(pos rl.Vector3, color rl.Color) {
	rl.DrawSphere(pos, radius, color)
}

func DrawBounds(bounds *rl.BoundingBox) {
	rl.DrawBoundingBox(*bounds, rl.Red)
}

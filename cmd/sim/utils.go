package sim

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func SpawnBoids(count int, bound *rl.Rectangle) []*Boid {
	boids := make([]*Boid, count)
	for i := range count {
		boids[i] = NewBoid(
			rl.NewVector2(
				bound.X+rand.Float32()*bound.Width,
				bound.Y+rand.Float32()*bound.Height,
			),
			rl.NewVector2(
				rand.Float32()*2-1,
				rand.Float32()*2-1,
			),
		)
	}
	return boids
}

func densityColor(d uint8) rl.Color {
	t := float32(d) / 255.0
	r := uint8(80 + t*80)
	g := uint8(200 - t*160)
	b := uint8(255 - t*35)
	return rl.NewColor(r, g, b, 255)
}

func renderBoid(b *Boid) {
	dir := b.Direction
	perp := rl.NewVector2(-dir.Y, dir.X)
	color := densityColor(b.Density)

	tip   := rl.Vector2Add(b.Position, rl.Vector2Scale(dir, Radius*1.5))
	baseL := rl.Vector2Add(rl.Vector2Subtract(b.Position, rl.Vector2Scale(dir, Radius*0.5)), rl.Vector2Scale(perp, Radius))
	baseR := rl.Vector2Subtract(rl.Vector2Subtract(b.Position, rl.Vector2Scale(dir, Radius*0.5)), rl.Vector2Scale(perp, Radius))

	rl.DrawTriangle(tip, baseR, baseL, color)
}

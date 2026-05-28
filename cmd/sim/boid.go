package sim

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	defVel = 120.0

	drawRadius = 8.0
	// Perception radii for the three main behaviors
	// Tweak to change how far boids can see for each behavior
	separationRadius = drawRadius * 4
	alignmentRadius  = drawRadius * 6
	cohesionRadius   = drawRadius * 8
	fovThreshold     = 0.0

	// Weights for the three main behaviors
	// Reynolds' original values were around 1.5 for separation, 1.0 for cohesion, and 1.0 for alignment
	separationWeight = 1.5
	cohesionWeight   = 1.0
	alignmentWeight  = 1.0

	turnSpeed = 5.0
)

type Boid struct {
	Position  rl.Vector2
	Direction rl.Vector2
	velocity  float32
	Density   uint8
}

func NewBoid(p, d rl.Vector2) *Boid {
	return &Boid{
		Position:  p,
		Direction: d,
		velocity:  defVel,
		Density:   1.0,
	}
}

func (b *Boid) Move(bounds *rl.Rectangle, df float32) {
	b.wrapAround(bounds)
	b.Position.X += b.Direction.X * b.velocity * df
	b.Position.Y += b.Direction.Y * b.velocity * df
}

func (b *Boid) UpdateDir(others []*Boid, df float32) {
	if len(others) == 0 {
		return
	}

	alignCount, cohesionCount := 0, 0
	separation := rl.NewVector2(0, 0)
	alignment := rl.NewVector2(0, 0)
	cohesion := rl.NewVector2(0, 0)

	for _, o := range others {
		if o == b {
			continue
		}
		dist := rl.Vector2Distance(b.Position, o.Position)
		if dist <= 0 || dist > cohesionRadius {
			continue
		}

		toOther := rl.Vector2Normalize(rl.Vector2Subtract(o.Position, b.Position))
		if rl.Vector2DotProduct(b.Direction, toOther) < fovThreshold {
			continue
		}

		if dist < separationRadius {
			diff := rl.Vector2Normalize(rl.Vector2Subtract(b.Position, o.Position))
			scale := separationWeight * (1.0 - dist/separationRadius)
			separation = rl.Vector2Add(separation, rl.Vector2Scale(diff, scale))
		}
		if dist < alignmentRadius {
			alignCount++
			alignment = rl.Vector2Add(alignment, o.Direction)
		}
		cohesionCount++
		cohesion = rl.Vector2Add(cohesion, o.Position)
	}

	newDir := rl.Vector2Add(b.Direction, separation)
	if alignCount > 0 {
		alignment = rl.Vector2Scale(alignment, 1.0/float32(alignCount))
		newDir = rl.Vector2Add(newDir, rl.Vector2Scale(alignment, alignmentWeight))
	}
	if cohesionCount > 0 {
		cohesion = rl.Vector2Scale(cohesion, 1.0/float32(cohesionCount))
		cohesionDir := rl.Vector2Normalize(rl.Vector2Subtract(cohesion, b.Position))
		newDir = rl.Vector2Add(newDir, rl.Vector2Scale(cohesionDir, cohesionWeight))
	}

	lerpedNewDir := rl.Vector2Lerp(b.Direction, newDir, turnSpeed*df)
	b.Direction = rl.Vector2Normalize(lerpedNewDir)
	b.mapNeighborsToDensity(cohesionCount)
	//b.mapNeighborsToDensityBitwise(neighborCount)
}

func (b *Boid) wrapAround(bounds *rl.Rectangle) {
	if b.Position.X < bounds.X {
		b.Position.X += bounds.Width
	} else if b.Position.X > bounds.X+bounds.Width {
		b.Position.X -= bounds.Width
	}
	if b.Position.Y < bounds.Y {
		b.Position.Y += bounds.Height
	} else if b.Position.Y > bounds.Y+bounds.Height {
		b.Position.Y -= bounds.Height
	}
}

func (b *Boid) mapNeighborsToDensity(neighborsCount int) {
	const maxNeighbors = 20
	if neighborsCount > maxNeighbors {
		neighborsCount = maxNeighbors
	}
	t := math.Sqrt(float64(neighborsCount) / maxNeighbors)
	b.Density = uint8(50 + t*205)
}

func (b *Boid) mapNeighborsToDensityBitwise(neighborsCount int) {
	if neighborsCount > 8 {
		neighborsCount = 8
	}
	b.Density = ^(uint8(0xFF) >> uint(neighborsCount))
}

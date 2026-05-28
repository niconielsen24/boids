package sim

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	defVel = 120.0

	Radius           = 8.0
	PerceptionRadius = 80.0
	fovThreshold     = 0.0

	separationWeight = 3.0
	cohesionWeight   = 1.1
	alignmentWeight  = 0.8

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

	neighborCount := 0
	separation := rl.NewVector2(0, 0)
	alignment := rl.NewVector2(0, 0)
	cohesion := rl.NewVector2(0, 0)

	for _, o := range others {
		if o == b {
			continue
		}
		dist := rl.Vector2Distance(b.Position, o.Position)
		if dist <= 0 {
			continue
		}
		if dist > PerceptionRadius {
			continue
		}

		toOther := rl.Vector2Normalize(rl.Vector2Subtract(o.Position, b.Position))
		dot := rl.Vector2DotProduct(b.Direction, toOther)

		if dot < fovThreshold {
			continue
		}

		diff := rl.Vector2Subtract(b.Position, o.Position)
		diff = rl.Vector2Normalize(diff)
		scale := separationWeight * (1.0 - dist/PerceptionRadius)
		separation = rl.Vector2Add(separation, rl.Vector2Scale(diff, scale))

		neighborCount++
		alignment = rl.Vector2Add(alignment, o.Direction)
		cohesion = rl.Vector2Add(cohesion, o.Position)
	}

	newDir := rl.Vector2Add(b.Direction, separation)
	if neighborCount > 0 {
		alignment = rl.Vector2Scale(alignment, 1.0/float32(neighborCount))
		cohesion = rl.Vector2Scale(cohesion, 1.0/float32(neighborCount))
		cohesionDir := rl.Vector2Normalize(rl.Vector2Subtract(cohesion, b.Position))
		newDir = rl.Vector2Add(newDir, rl.Vector2Scale(alignment, alignmentWeight))
		newDir = rl.Vector2Add(newDir, rl.Vector2Scale(cohesionDir, cohesionWeight))
	}

	lerpedNewDir := rl.Vector2Lerp(b.Direction, newDir, turnSpeed*df)
	b.Direction = rl.Vector2Normalize(lerpedNewDir)
	b.mapNeighborsToDensity(neighborCount)
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
	b.Density = uint8(50 + float32(neighborsCount)/maxNeighbors*205)
}

package logic

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	separationWeight = 1.2
	alignmentWeight  = 0.9
	cohesionWeight   = 0.8
	defSpeed         = 20
	viewRadius       = 5
	sqViewRadius     = viewRadius * viewRadius
)

var (
	defPos = rl.NewVector3(0, 0, 0)
	defDir = rl.NewVector3(1, 0, 0)
)

const (
	Neutral uint8 = iota
	Normal
	Warning
	Alert
	Fleeing
)

type Boid struct {
	position  rl.Vector3
	direction rl.Vector3
	speed     float32
}

func NewBoid(pos rl.Vector3) *Boid {
	return &Boid{
		position:  pos,
		direction: defDir,
		speed:     defSpeed,
	}
}

func (b *Boid) GetPos() rl.Vector3 {
	return b.position
}

func (b *Boid) Update(others []*Boid, bounds *rl.BoundingBox, delta float32) {
	var separationSteer rl.Vector3 = rl.Vector3Zero()
	var alignmentSteer rl.Vector3 = rl.Vector3Zero()
	var cohesionCenter rl.Vector3 = rl.Vector3Zero()
	neighborCount := 0

	for _, other := range others {
		if other == b {
			continue
		}

		if b.inView(other) {
			neighborCount++

			// separation
			away := rl.Vector3Subtract(b.position, other.position)
			dist := rl.Vector3Length(away)
			if dist > 0 {
				separationSteer = rl.Vector3Add(separationSteer, rl.Vector3Scale(away, separationWeight/dist))
			}

			// alignment
			alignmentSteer = rl.Vector3Add(alignmentSteer, rl.Vector3Scale(other.direction, alignmentWeight))

			// cohesion: accumulate neighbor positions
			cohesionCenter = rl.Vector3Add(cohesionCenter, other.position)
		}
	}

	if neighborCount > 0 {
		// steer toward center of mass of neighbors
		cohesionCenter = rl.Vector3Scale(cohesionCenter, 1.0/float32(neighborCount))
		cohesionSteer := rl.Vector3Scale(rl.Vector3Subtract(cohesionCenter, b.position), cohesionWeight)

		totalSteer := rl.Vector3Add(separationSteer, rl.Vector3Add(alignmentSteer, cohesionSteer))
		if rl.Vector3Length(totalSteer) > 0 {
			b.direction = rl.Vector3Normalize(rl.Vector3Add(b.direction, totalSteer))
		}
	}

	scaled := rl.Vector3Scale(b.direction, b.speed*delta)
	newPos := rl.Vector3Add(b.position, scaled)

	b.avoidWalls(bounds)
	b.position = newPos
}

func (b *Boid) avoidWalls(bounds *rl.BoundingBox) {
	if b.position.X < bounds.Min.X {
		b.direction.X = 1
		return
	} else if b.position.X > bounds.Max.X {
		b.direction.X = -1
		return
	}

	if b.position.Y < bounds.Min.Y {
		b.direction.Y = 1
		return
	} else if b.position.Y > bounds.Max.Y {
		b.direction.Y = -1
		return
	}

	if b.position.Z < bounds.Min.Z {
		b.direction.Z = 1
		return
	} else if b.position.Z > bounds.Max.Z {
		b.direction.Z = -1
		return
	}
}

func (b *Boid) inView(other *Boid) bool {
	dx := other.position.X - b.position.X
	dy := other.position.Y - b.position.Y
	dz := other.position.Z - b.position.Z
	sqDist := dx*dx + dy*dy + dz*dz
	return sqDist <= sqViewRadius
}

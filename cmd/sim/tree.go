package sim

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	maxBoidsPerNode = 8
)

type Tree struct {
	bounds   rl.Rectangle
	boids    []*Boid
	children [4]*Tree
}

func NewTree(bounds rl.Rectangle) *Tree {
	return &Tree{
		bounds: bounds,
		boids:  make([]*Boid, 0),
	}
}

func (t *Tree) subdivide() {
	halfWidth := t.bounds.Width / 2
	halfHeight := t.bounds.Height / 2

	t.children[0] = NewTree(rl.NewRectangle(t.bounds.X, t.bounds.Y, halfWidth, halfHeight))                      // NW
	t.children[1] = NewTree(rl.NewRectangle(t.bounds.X+halfWidth, t.bounds.Y, halfWidth, halfHeight))            // NE
	t.children[2] = NewTree(rl.NewRectangle(t.bounds.X, t.bounds.Y+halfHeight, halfWidth, halfHeight))           // SW
	t.children[3] = NewTree(rl.NewRectangle(t.bounds.X+halfWidth, t.bounds.Y+halfHeight, halfWidth, halfHeight)) // SE
}

func (t *Tree) Insert(boid *Boid) {
	if !rl.CheckCollisionPointRec(boid.Position, t.bounds) {
		return
	}

	if len(t.boids) < maxBoidsPerNode {
		t.boids = append(t.boids, boid)
		return
	}

	if t.children[0] == nil {
		t.subdivide()
	}

	for _, child := range t.children {
		if rl.CheckCollisionPointRec(boid.Position, child.bounds) {
			child.Insert(boid)
			return
		}
	}
}

func (t *Tree) QueryRange(rangeRect rl.Rectangle, found *[]*Boid) {
	if !rl.CheckCollisionRecs(t.bounds, rangeRect) {
		return
	}

	for _, b := range t.boids {
		if rl.CheckCollisionPointRec(b.Position, rangeRect) {
			*found = append(*found, b)
		}
	}

	if t.children[0] != nil {
		for _, child := range t.children {
			child.QueryRange(rangeRect, found)
		}
	}
}

func (t *Tree) Clear() {
	t.boids = t.boids[:0]
	for _, child := range t.children {
		if child != nil {
			child.Clear()
		}
	}
}

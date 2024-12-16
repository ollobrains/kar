package arc

import (
	"strconv"
	"time"
)

type Rect struct {
	X, Y, W, H float64
}

// Overlaps checks if the rectangle overlaps with another rectangle
func (r *Rect) Overlaps(x, y, w, h float64) bool {
	return r.X+r.W > x && x+w > r.X && r.Y+r.H > y && y+h > r.Y
}

// Overlaps checks if the rectangle overlaps with another rectangle
func (r *Rect) OverlapsR(box *Rect) bool {
	return r.X+r.W > box.X && box.X+box.W > r.X && r.Y+r.H > box.Y && box.Y+box.H > r.Y
}

func (r *Rect) String() string {
	x := "X: " + strconv.FormatFloat(r.X, 'f', -1, 64)
	y := "Y: " + strconv.FormatFloat(r.X, 'f', -1, 64)
	return x + y
}

type DrawOptions struct {
	Scale float64
	FlipX bool
}

type Timer struct {
	Duration time.Duration
	Elapsed  time.Duration
}

type AnimationFrameIndex struct {
	Index int
}

type ItemID struct {
	ID uint16
}

type Health struct {
	Health    float64
	MaxHealth float64
}

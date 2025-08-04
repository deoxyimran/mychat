package animation

import (
	"image/color"
	"time"

	"gioui.org/f32"
)

type Animator struct {
	start    time.Time
	duration time.Duration
	running  bool
	progress float32
}

func (a *Animator) Start(d time.Duration) {
	a.start = time.Now()
	a.duration = d
	a.running = true
	a.progress = 0
}

func (a *Animator) Update() {
	if !a.running {
		return
	}
	elapsed := time.Since(a.start)
	t := float32(elapsed) / float32(a.duration)
	if t >= 1 {
		t = 1
		a.running = false
	}
	a.progress = t
}

func (a *Animator) Running() bool {
	return a.running
}

func (a *Animator) lerp() float32 {
	return a.progress
}

func (a *Animator) LerpFloat(aVal, bVal float32) float32 {
	t := a.lerp()
	return aVal + (bVal-aVal)*t
}

func (a *Animator) LerpColor(c1, c2 color.NRGBA) color.NRGBA {
	t := a.lerp()
	return color.NRGBA{
		R: uint8(float32(c1.R) + (float32(c2.R)-float32(c1.R))*t),
		G: uint8(float32(c1.G) + (float32(c2.G)-float32(c1.G))*t),
		B: uint8(float32(c1.B) + (float32(c2.B)-float32(c1.B))*t),
		A: uint8(float32(c1.A) + (float32(c2.A)-float32(c1.A))*t),
	}
}

func (a *Animator) LerpPos(p1, p2 f32.Point) f32.Point {
	t := a.lerp()
	return f32.Point{
		X: p1.X + (p2.X-p1.X)*t,
		Y: p1.Y + (p2.Y-p1.Y)*t,
	}
}

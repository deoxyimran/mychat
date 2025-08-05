package animation

import (
	"image/color"

	"gioui.org/f32"
)

// Position interpolator
func InterpPoint(a, b f32.Point, t float32) f32.Point {
	return f32.Point{
		X: a.X + (b.X-a.X)*t,
		Y: a.Y + (b.Y-a.Y)*t,
	}
}

// Scale interpolator (same as position)
var InterpScale = InterpPoint

// Rotation interpolator (angle in radians)
func InterpRotation(a, b float32, t float32) float32 {
	return a + (b-a)*t
}

// Opacity: float32 (0 to 1)
func InterpOpacity(a, b float32, t float32) float32 {
	return a + (b-a)*t
}

// Color interpolator (NRGBA)
func InterpColor(a, b color.NRGBA, t float32) color.NRGBA {
	return color.NRGBA{
		R: uint8(float32(a.R) + (float32(b.R)-float32(a.R))*t),
		G: uint8(float32(a.G) + (float32(b.G)-float32(a.G))*t),
		B: uint8(float32(a.B) + (float32(b.B)-float32(a.B))*t),
		A: uint8(float32(a.A) + (float32(b.A)-float32(a.A))*t),
	}
}

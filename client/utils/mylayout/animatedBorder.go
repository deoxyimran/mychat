package mylayout

import (
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/deoxyimran/mychat/client/utils/animation"
)

type BorderSide int

const (
	BORDER_SIDE_TOP BorderSide = iota
	BORDER_SIDE_BOT
	BORDER_SIDE_LEFT
	BORDER_SIDE_RIGHT
)

type AnimatedBorderSide struct {
	shouldRenderBorder bool
	side               BorderSide
	color              color.NRGBA
	width, padding     float32
	animator           animation.Animator
	posKeyFrame        []animation.Keyframe[f32.Point]
}

func NewAnimatedBorderSide(side BorderSide, color color.NRGBA, width, padding float32) AnimatedBorderSide {
	return AnimatedBorderSide{
		side:     side,
		color:    color,
		width:    width,
		padding:  padding,
		animator: animation.Animator{},
	}
}

func (a *AnimatedBorderSide) Layout(gtx C) D {
}

var posKeyFramed = animation.KeyframedProperty[f32.Point]{
	Keyframes: []animation.Keyframe[f32.Point]{
		{Time: 0, Value: f32.Pt(0, 0)}, // dummy
		{Time: 1, Value: f32.Pt(0, 0)}, // dummy
	},
	Interp: animation.InterpPoint,
}

func getInterpPos(p1, p2 f32.Point, t float32) f32.Point {
	posKeyFramed.Keyframes[0].Value = p1
	posKeyFramed.Keyframes[1].Value = p2
	return posKeyFramed.Sample(t)
}

func BorderOneSide(gtx C, inner layout.Widget, focused bool, borderSide BorderSide, anim *animation.Animator, width float32, padding unit.Dp, color color.NRGBA) D {
}

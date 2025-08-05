package mylayout

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/deoxyimran/mychat/client/utils/animation"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// Layout utils
func Border(gtx C, inner layout.Widget, width float32, padding unit.Dp, color color.NRGBA) D {
	// Layout inner first
	dims := layout.UniformInset(padding).Layout(gtx, inner)

	// Layout border
	defer clip.Stroke{
		Width: width,
		Path:  clip.UniformRRect(image.Rectangle{Max: dims.Size}, int(width)).Path(gtx.Ops),
	}.Op().Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return dims
}

type BorderSide int

const (
	BORDER_SIDE_TOP BorderSide = iota
	BORDER_SIDE_BOT
	BORDER_SIDE_LEFT
	BORDER_SIDE_RIGHT
)

var shouldRenderBorder bool

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

func BorderOneSide(gtx C, inner layout.Widget, borderSide BorderSide, anim *animation.Animator, width float32, padding unit.Dp, color color.NRGBA) D {
	if anim.Running() {
		anim.Update()
		gtx.Execute(op.InvalidateCmd{})
	}
	// Layout inner first
	dims := layout.UniformInset(padding).Layout(gtx, inner)
	// Layout border
	var outline clip.PathSpec
	if borderSide == BORDER_SIDE_TOP {
		var path clip.Path
		p1 := f32.Pt(0, 0)
		p2 := getInterpPos(p1, f32.Pt(float32(dims.Size.X), 0), anim.Progress())
		if p2.X-p1.X <= 0.01 {
			shouldRenderBorder = false
		} else {
			shouldRenderBorder = true
			path.Begin(gtx.Ops)
			path.MoveTo(p1)
			path.LineTo(p2)
			outline = path.End()
		}
	} else if borderSide == BORDER_SIDE_BOT {
		var path clip.Path
		p1 := f32.Pt(0, float32(dims.Size.Y))
		p2 := getInterpPos(p1, f32.Pt(float32(dims.Size.X), float32(dims.Size.Y)), anim.Progress())
		if p2.X-p1.X <= 0.01 {
			shouldRenderBorder = false
		} else {
			shouldRenderBorder = true
			path.Begin(gtx.Ops)
			path.MoveTo(p1)
			path.LineTo(p2)
			outline = path.End()
		}
	} else if borderSide == BORDER_SIDE_LEFT {
		var path clip.Path
		p1 := f32.Pt(0, 0)
		p2 := getInterpPos(p1, f32.Pt(0, float32(dims.Size.Y)), anim.Progress())
		if p2.X-p1.X <= 0.01 {
			shouldRenderBorder = false
		} else {
			shouldRenderBorder = true
			path.Begin(gtx.Ops)
			path.MoveTo(p1)
			path.LineTo(p2)
			outline = path.End()
		}
	} else {
		var path clip.Path
		p1 := f32.Pt(float32(dims.Size.X), 0)
		p2 := getInterpPos(p1, f32.Pt(float32(dims.Size.X), float32(dims.Size.Y)), anim.Progress())
		if p2.X-p1.X <= 0.01 {
			shouldRenderBorder = false
		} else {
			shouldRenderBorder = true
			path.Begin(gtx.Ops)
			path.MoveTo(p1)
			path.LineTo(p2)
			outline = path.End()
		}
	}
	if shouldRenderBorder {
		defer clip.Stroke{
			Width: width,
			Path:  outline,
		}.Op().Push(gtx.Ops).Pop()
		paint.ColorOp{Color: color}.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)
	}

	return dims
}

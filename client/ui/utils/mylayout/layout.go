package mylayout

import (
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
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

func BorderOneSide(gtx C, inner layout.Widget, borderSide BorderSide, width float32, padding unit.Dp, color color.NRGBA) D {
	// Layout inner first
	dims := layout.UniformInset(padding).Layout(gtx, inner)

	// Layout border
	var outline clip.PathSpec
	if borderSide == BORDER_SIDE_TOP {
		var path clip.Path
		path.Begin(gtx.Ops)
		path.MoveTo(f32.Pt(0, 0))
		path.LineTo(f32.Pt(float32(dims.Size.X), 0))
		outline = path.End()
	} else if borderSide == BORDER_SIDE_BOT {
		var path clip.Path
		path.Begin(gtx.Ops)
		path.MoveTo(f32.Pt(0, float32(dims.Size.Y)))
		path.LineTo(f32.Pt(float32(dims.Size.X), float32(dims.Size.Y)))
		outline = path.End()
	} else if borderSide == BORDER_SIDE_LEFT {
		var path clip.Path
		path.Begin(gtx.Ops)
		path.MoveTo(f32.Pt(0, 0))
		path.LineTo(f32.Pt(0, float32(dims.Size.Y)))
		outline = path.End()
	} else {
		var path clip.Path
		path.Begin(gtx.Ops)
		path.MoveTo(f32.Pt(float32(dims.Size.X), 0))
		path.LineTo(f32.Pt(float32(dims.Size.X), float32(dims.Size.Y)))
		outline = path.End()
	}
	defer clip.Stroke{
		Width: width,
		Path:  outline,
	}.Op().Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return dims
}

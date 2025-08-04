package ui

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/deoxyimran/mychat/client/ui/theme"
	"github.com/deoxyimran/mychat/client/ui/utils/mylayout"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type LoginScreen struct {
	th      *material.Theme
	isLogin bool
	// login ui
	loginBtn          *widget.Clickable
	loginUsernameEdit *widget.Editor
	loginPasswordEdit *widget.Editor
	// signin ui
	signinBtn                 *widget.Clickable
	signinUsernameEdit        *widget.Editor
	signinPasswordEdit        *widget.Editor
	signinConfirmPasswordEdit *widget.Editor
}

func NewLoginScreen(isLogin bool) *LoginScreen {
	return &LoginScreen{
		th:       material.NewTheme(),
		isLogin:  isLogin,
		loginBtn: &widget.Clickable{},

		loginUsernameEdit: &widget.Editor{
			SingleLine: true,
		},
		loginPasswordEdit: &widget.Editor{
			SingleLine: true,
		},

		signinBtn: &widget.Clickable{},
		signinUsernameEdit: &widget.Editor{
			SingleLine: true,
		},
		signinPasswordEdit: &widget.Editor{
			SingleLine: true,
		},
		signinConfirmPasswordEdit: &widget.Editor{
			SingleLine: true,
		},
	}
}

func (l *LoginScreen) Layout(gtx C, screenPointer *Screen) D {
	var contentLayout func(C) D
	if l.isLogin {
		contentLayout = func(gtx C) D {
			gtx.Constraints.Min = image.Pt(0, 0) // Reset Constraints Min
			return layout.Flex{                  // Horizontally align
				Axis:      layout.Horizontal,
				Alignment: layout.Middle,
				Spacing:   layout.SpaceAround,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Flex{ // Vertically align
						Axis:      layout.Vertical,
						Alignment: layout.Middle,
					}.Layout(gtx,
						// Heading
						layout.Rigid(func(gtx C) D {
							h4 := material.H4(l.th, "Login")
							h4.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
							h4.Alignment = text.Middle
							return layout.Inset{Bottom: unit.Dp(5)}.Layout(gtx, h4.Layout)
						}),
						// Username edit
						layout.Rigid(func(gtx C) D {
							c := gtx.Constraints
							c.Max.X, c.Min.X = 400, 330
							gtx.Constraints = c
							edit := material.Editor(l.th, l.loginUsernameEdit, "Enter your email")
							edit.Color = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
							edit.HintColor = color.NRGBA{R: 135, G: 135, B: 135, A: 220}
							edit.TextSize = unit.Sp(14)
							return layout.UniformInset(unit.Dp(10)).Layout(gtx,
								func(gtx C) D {
									return mylayout.BorderOneSideLayout(gtx, edit.Layout, mylayout.BORDER_SIDE_BOT, 2, 4, color.NRGBA{R: 200, G: 5, B: 30, A: 255}) // rgb(200, 5, 30)
								},
							)
						}),
						// Password edit
						layout.Rigid(func(gtx C) D {
							c := gtx.Constraints
							c.Max.X, c.Min.X = 400, 330
							gtx.Constraints = c
							edit := material.Editor(l.th, l.loginPasswordEdit, "Enter your password")
							edit.Color = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
							edit.HintColor = color.NRGBA{R: 135, G: 135, B: 135, A: 220}
							edit.TextSize = unit.Sp(14)
							edit.Editor.Mask = '*'
							return layout.UniformInset(unit.Dp(10)).Layout(gtx,
								func(gtx C) D {
									return mylayout.BorderOneSideLayout(gtx, edit.Layout, mylayout.BORDER_SIDE_BOT, 2, 4, color.NRGBA{R: 200, G: 5, B: 30, A: 255})
								},
							)
						}),
					)
				}),
			)
		}
	} else {
		contentLayout = func(gtx C) D {
			gtx.Constraints.Min = image.Pt(0, 0) // Reset Constraints Min
			return layout.Flex{                  // Horizontally align
				Axis:      layout.Horizontal,
				Alignment: layout.Middle,
				Spacing:   layout.SpaceAround,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return layout.Flex{ // Vertically align
						Axis:      layout.Vertical,
						Alignment: layout.Middle,
					}.Layout(gtx,
						// Heading
						layout.Rigid(func(gtx C) D {
							h4 := material.H4(l.th, "Signup")
							h4.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
							h4.Alignment = text.Middle
							return layout.Inset{Bottom: unit.Dp(5)}.Layout(gtx, h4.Layout)
						}),
						// Username edit
						layout.Rigid(func(gtx C) D {
							c := gtx.Constraints
							c.Max.X, c.Min.X = 400, 330
							gtx.Constraints = c
							edit := material.Editor(l.th, l.signinUsernameEdit, "Enter your email")
							edit.Color = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
							edit.HintColor = color.NRGBA{R: 135, G: 135, B: 135, A: 220}
							edit.TextSize = unit.Sp(14)
							return layout.UniformInset(unit.Dp(10)).Layout(gtx,
								func(gtx C) D {
									return mylayout.BorderOneSideLayout(gtx, edit.Layout, mylayout.BORDER_SIDE_BOT, 2, 4, color.NRGBA{R: 200, G: 5, B: 30, A: 255}) // rgb(200, 5, 30)
								},
							)
						}),
						// Password edit
						layout.Rigid(func(gtx C) D {
							c := gtx.Constraints
							c.Max.X, c.Min.X = 400, 330
							gtx.Constraints = c
							edit := material.Editor(l.th, l.signinPasswordEdit, "Enter a password")
							edit.Color = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
							edit.HintColor = color.NRGBA{R: 135, G: 135, B: 135, A: 220}
							edit.TextSize = unit.Sp(14)
							edit.Editor.Mask = '*'
							gtx.Focused(l.signinPasswordEdit)
							return layout.UniformInset(unit.Dp(10)).Layout(gtx,
								func(gtx C) D {
									return mylayout.BorderOneSideLayout(gtx, edit.Layout, mylayout.BORDER_SIDE_BOT, 2, 4, color.NRGBA{R: 200, G: 5, B: 30, A: 255})
								},
							)
						}),
						// Confirm Password edit
						layout.Rigid(func(gtx C) D {
							c := gtx.Constraints
							c.Max.X, c.Min.X = 400, 330
							gtx.Constraints = c
							edit := material.Editor(l.th, l.signinConfirmPasswordEdit, "Confirm password")
							edit.Color = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
							edit.HintColor = color.NRGBA{R: 135, G: 135, B: 135, A: 220}
							edit.TextSize = unit.Sp(14)
							edit.Editor.Mask = '*'
							return layout.UniformInset(unit.Dp(10)).Layout(gtx,
								func(gtx C) D {
									return mylayout.BorderOneSideLayout(gtx, edit.Layout, mylayout.BORDER_SIDE_BOT, 2, 4, color.NRGBA{R: 200, G: 5, B: 30, A: 255})
								},
							)
						}),
						// Button
						layout.Rigid(func(gtx C) D {
							c := gtx.Constraints
							c.Max.X, c.Min.X = 400, 330
							gtx.Constraints = c
							btn := material.Button(l.th, l.signinBtn, "Sign in")
							return layout.UniformInset(unit.Dp(10)).Layout(gtx, btn.Layout)
						}),
					)
				}),
			)
		}
	}
	return layout.Background{}.Layout(gtx,
		// Fullscreen background
		func(gtx C) D {
			defer clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops).Pop()
			color := theme.BackgroundColor()
			paint.ColorOp{Color: color}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
			return layout.Dimensions{Size: gtx.Constraints.Max}
		},
		contentLayout,
	)
}

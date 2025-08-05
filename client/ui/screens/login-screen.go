package screens

import (
	"image"
	"image/color"
	"time"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/deoxyimran/mychat/client/ui/theme"
	"github.com/deoxyimran/mychat/client/utils/animation"
	"github.com/deoxyimran/mychat/client/utils/mylayout"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type login struct {
	btn *widget.Clickable

	usernameEdit        *widget.Editor
	usernameEditExState editorExState
	animUsernameEdit    animation.Animator

	passEdit        *widget.Editor
	passEditExState editorExState
	animPassEdit    animation.Animator
}

type signin struct {
	btn *widget.Clickable

	usernameEdit        *widget.Editor
	usernameEditExState editorExState
	animUsernameEdit    animation.Animator

	passEdit        *widget.Editor
	passEditExState editorExState
	animPassEdit    animation.Animator

	confirmPassEdit        *widget.Editor
	confirmPassEditExState editorExState
	animConfirmPassEdit    animation.Animator
}

type editorExState struct {
	pointerEntered  bool
	pointerReleased bool
}

type LoginScreen struct {
	th      *material.Theme
	isLogin bool
	login   login
	signin  signin
}

func NewLoginScreen(isLogin bool) *LoginScreen {
	return &LoginScreen{
		th:      material.NewTheme(),
		isLogin: isLogin,
		login: login{
			btn: &widget.Clickable{},
			usernameEdit: &widget.Editor{
				SingleLine: true,
			},
			passEdit: &widget.Editor{
				SingleLine: true,
			},
		},
		signin: signin{
			btn: &widget.Clickable{},
			usernameEdit: &widget.Editor{
				SingleLine: true,
			},
			passEdit: &widget.Editor{
				SingleLine: true,
			},
			confirmPassEdit: &widget.Editor{
				SingleLine: true,
			},
		},
	}
}

func (l *LoginScreen) Layout(gtx C, screenPointer *Screen) D {
	var contentLayout func(C) D
	if l.isLogin { // Login
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
							sz := 350
							c := gtx.Constraints
							c.Max.X, c.Min.X = sz, sz
							gtx.Constraints = c
							edit := material.Editor(l.th, l.login.usernameEdit, "Enter your email")
							edit.Color = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
							edit.HintColor = color.NRGBA{R: 135, G: 135, B: 135, A: 220}
							edit.TextSize = unit.Sp(14)
							dims := layout.UniformInset(unit.Dp(10)).Layout(gtx,
								func(gtx C) D {
									return mylayout.BorderOneSide(gtx, edit.Layout, mylayout.BORDER_SIDE_BOT, &l.login.animUsernameEdit, 2, 4, color.NRGBA{R: 200, G: 5, B: 30, A: 255}) // rgb(200, 5, 30)
								},
							)
							// Check for focused event
							defer clip.Rect{Max: dims.Size}.Push(gtx.Ops).Pop()
							event.Op(gtx.Ops, &l.login.usernameEdit)
							var entered bool
							for {
								ev, ok := gtx.Source.Event(pointer.Filter{
									Target: &l.login.usernameEdit,
									Kinds:  pointer.Press | pointer.Enter | pointer.Leave,
								})
								if !ok {
									break
								}
								if x, ok := ev.(pointer.Event); ok {
									switch x.Kind {
									case pointer.Enter:
										entered = true
									case pointer.Release:
										if entered {
											l.login.animUsernameEdit.Start(time.Millisecond*250, animation.Once)
											gtx.Execute(op.InvalidateCmd{})
										}
									}
								}
							}
							return dims
						}),
						// Password edit
						layout.Rigid(func(gtx C) D {
							sz := 350
							c := gtx.Constraints
							c.Max.X, c.Min.X = sz, sz
							gtx.Constraints = c
							edit := material.Editor(l.th, l.login.passEdit, "Enter your password")
							edit.Color = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
							edit.HintColor = color.NRGBA{R: 135, G: 135, B: 135, A: 220}
							edit.TextSize = unit.Sp(14)
							edit.Editor.Mask = '*'
							dims := layout.UniformInset(unit.Dp(10)).Layout(gtx,
								func(gtx C) D {
									return mylayout.BorderOneSide(gtx, edit.Layout, mylayout.BORDER_SIDE_BOT, &l.login.animPassEdit, 2, 4, color.NRGBA{R: 200, G: 5, B: 30, A: 255})
								},
							)
							// Check for focused event
							defer clip.Rect{Max: dims.Size}.Push(gtx.Ops).Pop()
							event.Op(gtx.Ops, &l.login.passEdit)
							var entered bool
							for {
								ev, ok := gtx.Source.Event(pointer.Filter{
									Target: &l.login.passEdit,
									Kinds:  pointer.Press | pointer.Enter | pointer.Leave,
								})
								if !ok {
									break
								}
								if x, ok := ev.(pointer.Event); ok {
									switch x.Kind {
									case pointer.Enter:
										entered = true
									case pointer.Release:
										if entered {
											l.login.animPassEdit.Start(time.Millisecond*250, animation.Once)
											gtx.Execute(op.InvalidateCmd{})
										}
									}
								}
							}
							return dims
						}),
					)
				}),
			)
		}
	} else { // Signin
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
							sz := 350
							c := gtx.Constraints
							c.Max.X, c.Min.X = sz, sz
							gtx.Constraints = c
							edit := material.Editor(l.th, l.signin.usernameEdit, "Enter your email")
							edit.Color = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
							edit.HintColor = color.NRGBA{R: 135, G: 135, B: 135, A: 220}
							edit.TextSize = unit.Sp(14)
							defer pointer.PassOp{}.Push(gtx.Ops).Pop()
							dims := layout.UniformInset(unit.Dp(10)).Layout(gtx,
								func(gtx C) D {
									return mylayout.BorderOneSide(gtx,
										edit.Layout,
										l.signin.passEditExState.pointerEntered && ,
										mylayout.BORDER_SIDE_BOT,
										&l.signin.animUsernameEdit,
										2, 4,
										color.NRGBA{R: 200, G: 5, B: 30, A: 255}, // rgb(200, 5, 30)
									)
								},
							)
							// Check for focused event
							defer clip.Rect{Max: dims.Size}.Push(gtx.Ops).Pop()
							event.Op(gtx.Ops, &l.signin.usernameEdit)
							entered := true
							for {
								ev, ok := gtx.Source.Event(pointer.Filter{
									Target: &l.signin.usernameEdit,
									Kinds:  pointer.Press | pointer.Release | pointer.Enter | pointer.Leave,
								})
								if !ok {
									break
								}
								if x, ok := ev.(pointer.Event); ok {
									switch x.Kind {
									case pointer.Enter:
										entered = true
									case pointer.Release:
										if entered {
											l.signin.animUsernameEdit.Start(time.Millisecond*250, animation.Once)
											gtx.Execute(op.InvalidateCmd{})
										}
									}
								}
							}
							return dims
						}),
						// Password edit
						layout.Rigid(func(gtx C) D {
							sz := 350
							c := gtx.Constraints
							c.Max.X, c.Min.X = sz, sz
							gtx.Constraints = c
							edit := material.Editor(l.th, l.signin.passEdit, "Enter a password")
							edit.Color = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
							edit.HintColor = color.NRGBA{R: 135, G: 135, B: 135, A: 220}
							edit.TextSize = unit.Sp(14)
							edit.Editor.Mask = '*'
							dims := layout.UniformInset(unit.Dp(10)).Layout(gtx,
								func(gtx C) D {
									return mylayout.BorderOneSide(gtx, edit.Layout, mylayout.BORDER_SIDE_BOT, &l.signin.animPassEdit, 2, 4, color.NRGBA{R: 200, G: 5, B: 30, A: 255})
								},
							)
							// Check for focused event
							// defer clip.Rect{Max: dims.Size}.Push(gtx.Ops).Pop()
							// event.Op(gtx.Ops, &l.signin.passEdit)
							// var entered bool
							// for {
							// 	ev, ok := gtx.Source.Event(pointer.Filter{
							// 		Target: &l.signin.passEdit,
							// 		Kinds:  pointer.Press | pointer.Enter | pointer.Leave,
							// 	})
							// 	if !ok {
							// 		break
							// 	}
							// 	if x, ok := ev.(pointer.Event); ok {
							// 		switch x.Kind {
							// 		case pointer.Enter:
							// 			entered = true
							// 		case pointer.Release:
							// 			if entered {
							// 				l.signin.animPassEdit.Start(time.Millisecond * 500)
							// 				gtx.Execute(op.InvalidateCmd{})
							// 			}
							// 		}
							// 	}
							// }
							return dims
						}),
						// Confirm Password edit
						layout.Rigid(func(gtx C) D {
							sz := 350
							c := gtx.Constraints
							c.Max.X, c.Min.X = sz, sz
							gtx.Constraints = c
							edit := material.Editor(l.th, l.signin.confirmPassEdit, "Confirm password")
							edit.Color = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
							edit.HintColor = color.NRGBA{R: 135, G: 135, B: 135, A: 220}
							edit.TextSize = unit.Sp(14)
							edit.Editor.Mask = '*'
							dims := layout.UniformInset(unit.Dp(10)).Layout(gtx,
								func(gtx C) D {
									return mylayout.BorderOneSide(gtx, edit.Layout, mylayout.BORDER_SIDE_BOT, &l.signin.animConfirmPassEdit, 2, 4, color.NRGBA{R: 200, G: 5, B: 30, A: 255})
								},
							)
							// // Check for focused event
							// defer clip.Rect{Max: dims.Size}.Push(gtx.Ops).Pop()
							// event.Op(gtx.Ops, &l.signin.confirmPassEdit)
							// var entered bool
							// for {
							// 	ev, ok := gtx.Source.Event(pointer.Filter{
							// 		Target: &l.signin.confirmPassEdit,
							// 		Kinds:  pointer.Press | pointer.Enter | pointer.Leave,
							// 	})
							// 	if !ok {
							// 		break
							// 	}
							// 	if x, ok := ev.(pointer.Event); ok {
							// 		switch x.Kind {
							// 		case pointer.Enter:
							// 			entered = true
							// 		case pointer.Release:
							// 			if entered {
							// 				l.signin.animConfirmPassEdit.Start(time.Millisecond * 500)
							// 				gtx.Execute(op.InvalidateCmd{})
							// 			}
							// 		}
							// 	}
							// }
							return dims
						}),
						// Button
						layout.Rigid(func(gtx C) D {
							c := gtx.Constraints
							c.Max.X, c.Min.X = 350, 350
							gtx.Constraints = c
							btn := material.Button(l.th, l.signin.btn, "Sign in")
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

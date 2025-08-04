package ui

import (
	"gioui.org/widget/material"
)

type ChatScreen struct {
	th *material.Theme
}

func NewChatScreen(isLogin bool) *ChatScreen {
	ch := &ChatScreen{}
	return ch
}

func (ch *ChatScreen) Layout(screenPointer *Screen) {

}

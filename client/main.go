package main

import (
	"fmt"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"github.com/deoxyimran/mychat/client/ui/screens"
)

var screenPointer screens.Screen = screens.LOGIN_SCREEN
var prevScreenPointer screens.Screen = screens.CHAT_SCREEN

func main() {
	go func() {
		window := new(app.Window)
		window.Option(
			app.Title("MyChat"),
		)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	var ops op.Ops
	loginScreen := screens.NewLoginScreen(false)
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			fmt.Println("Exiting application...")
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			// Handle screen switching
			loginScreen.Layout(gtx, &screenPointer)
			e.Frame(gtx.Ops)
		}
	}
}

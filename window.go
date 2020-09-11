package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

import (
	"fmt"
	"os"
)

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {
	mw := &MyMainWindow {}

	MW := MainWindow{
		AssignTo: &mw.MainWindow,
		Title: "メインウィンドウ",
		MinSize: Size {150, 200},
		Size   : Size {300, 400},
		Layout: VBox {},
		Children: []Widget {
			PushButton {
				Text: "PB0",
			},
			HSplitter{
				StretchFactor: 20,
				Children: []Widget {
					PushButton {       
						Text: "PB1",   
					},
					HSpacer{},
					PushButton {
						Text: "PB2",
					},
				},
			},
			HSplitter{
				StretchFactor: 1,
				Children: []Widget {
					PushButton {
						Text: "PB3",
					},
					PushButton {
						Text: "PB4",
					},
				},
			},
			PushButton {
				Text: "PB5",
			},
		},
	}


	if _, err := MW.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

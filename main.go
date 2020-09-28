package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := &widget.Label{Text: "Hello Fyne!"}
	w.SetContent(&widget.Box{Children: []fyne.CanvasObject{
		hello,
		&widget.Button{Text: "Hi!", OnTapped: func() {
			hello.SetText("Welcome :)")
		}},
	}})

	w.ShowAndRun()
}

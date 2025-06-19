// main.go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
	"os"
	"os/user"
	"path/filepath"
	"fyne-hello/transpiler"
)

func main() {
	a := app.New()
	w := a.NewWindow("LiteJava IDE")
	w.Resize(fyne.NewSize(800, 600))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("LiteJava Code hier eingeben...")

	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Generierter Java-Code erscheint hier...")
	output.Disable()

	translateButton := widget.NewButton("Übersetzen", func() {
		javaCode := transpiler.Transpile(input.Text)
		output.SetText(javaCode)

		usr, _ := user.Current()
		filePath := filepath.Join(usr.HomeDir, "Main.java")
		os.WriteFile(filePath, []byte(javaCode), 0644)
	})

	w.SetContent(container.NewHSplit(
		container.NewBorder(nil, translateButton, nil, nil, input),
		output,
	))

	w.ShowAndRun()
}
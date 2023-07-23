package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// to keep track of number of files
var count int = 1

func showTextEditorApp() {

	w := myApp.NewWindow("Text Editor")

	w.Resize(fyne.NewSize(800, 600))

	//above text
	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Hydra Text Editor"),
		),
	)

	//textFile Creation
	content.Add(widget.NewButton("Create New File", func() {
		content.Add(widget.NewLabel("New File" + " " + (strconv.Itoa(count))))
		count++
	}))

	//inserting information to textFiles and savinf to os
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter Data to File")
	input.Resize(fyne.NewSize(400, 400))

	//to save files
	saveBtn := widget.NewButton("Save", func() {

		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)

				uc.Write(textData)
			}, w)

		saveFileDialog.SetFileName("New File " + strconv.Itoa(count-1) + ".txt")

		saveFileDialog.Show()
	})

	//to open button(to get content of file and display)
	openBtn := widget.NewButton("Open", func() {

		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				ReadData, _ := ioutil.ReadAll(r)

				Output := fyne.NewStaticResource("New File", ReadData)

				viewData := widget.NewMultiLineEntry()

				viewData.SetText(string(Output.StaticContent))

				w := fyne.CurrentApp().NewWindow(
					string(Output.StaticName))

				w.SetContent(container.NewScroll(viewData))

				w.Resize(fyne.NewSize(400, 400))

				w.Show()

			}, w)

		openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))

		openFileDialog.Show()
	})

	//displaying all contents
	editorContainer := container.NewVBox(
		content,
		input,
		container.NewHBox(
			saveBtn,
			openBtn,
		),
	)

	w.SetContent(editorContainer)
	w.Show()
}

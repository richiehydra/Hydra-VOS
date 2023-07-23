package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showGalleryApp() {

	w := myApp.NewWindow("Gallery App")

	w.Resize(fyne.NewSize(800, 600))

	//path of all files
	rootsrc := "C:\\Users\\hydrarichie\\Pictures\\Screenshots"

	files, err := ioutil.ReadDir(rootsrc)
	if err != nil {
		log.Fatal(err)
	}

	//pic array
	var picsArray []string

	for _, file := range files {
		//reading file names
		fmt.Println(file.Name(), file.IsDir())

		if !file.IsDir() {
			//chcking extsnions
			extension := strings.Split(file.Name(), ".")[1]

			//checking image or not

			if extension == "png" || extension == "jpeg" {
				picsArray = append(picsArray, rootsrc+"\\"+file.Name())

			}

		}
	}

	tabs := container.NewAppTabs(
		container.NewTabItem("Image 1", canvas.NewImageFromFile(picsArray[0])),
	)

	for i := 1; i < len(picsArray); i++ {
		tabs.Append(container.NewTabItem("Image"+strconv.Itoa(i+1), canvas.NewImageFromFile(picsArray[i])))
	}

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)

	w.Show()
}

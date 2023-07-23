package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalciApp() {

	output := ""

	var history []string
	//input
	input := widget.NewLabel(output)

	//All Buttons
	historyBtn := widget.NewButton("history", func() {

		for i := len(history) - 1; i >= 0; i-- {
			input.SetText(history[i])
		}

	})

	backbtn := widget.NewButton("back", func() {

		if len(output) > 0 {

			output = output[:len(output)-1]
			input.SetText(output)
		}

	})

	clearBtn := widget.NewButton("clear", func() {

		output = ""
		input.SetText(output)

	})

	openBtn := widget.NewButton("(", func() {

		output = output + "("
		input.SetText(output)

	})

	closeBtn := widget.NewButton(")", func() {

		output = output + ")"
		input.SetText(output)

	})

	divBtn := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)

	})

	sevenBtn := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)

	})

	eightBtn := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)

	})

	nineBtn := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)

	})

	mulBtn := widget.NewButton("*", func() {
		output = output + "*"
		input.SetText(output)

	})

	fourBtn := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)

	})

	fiveBtn := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)

	})

	sixBtn := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)

	})

	minusBtn := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)

	})

	oneBtn := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)

	})

	twoBtn := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)

	})

	threeBtn := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)

	})

	plusbtn := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)

	})

	zeroBtn := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)

	})

	dotBtn := widget.NewButton(".", func() {

		output = output + "."
		input.SetText(output)

	})

	eqBtn := widget.NewButton("=", func() {

		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				//to convert int to string
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strtoappend := output + "=" + ans
				history = append(history, strtoappend)
				output = ans
			} else {
				output = "Error:Invalid Expression"
			}
		} else {
			output = "Error:Invalid Expression"
		}

		input.SetText(output)

	})

	calcContainer := (container.NewVBox(
		input,

		container.NewGridWithColumns(1,

			container.NewGridWithColumns(2,
				historyBtn,
				backbtn,
			),

			container.NewGridWithColumns(4,
				clearBtn,
				openBtn,
				closeBtn,
				divBtn),

			container.NewGridWithColumns(4,
				nineBtn,
				eightBtn,
				sevenBtn,
				mulBtn),

			container.NewGridWithColumns(4,
				sixBtn,
				fiveBtn,
				fourBtn,
				plusbtn),

			container.NewGridWithColumns(4,

				threeBtn,
				twoBtn,
				oneBtn,
				minusBtn),

			container.NewGridWithColumns(2,

				container.NewGridWithColumns(2, zeroBtn,
					dotBtn),

				container.NewGridWithColumns(1,
					eqBtn),
			),
		),
	))

	w := myApp.NewWindow("Calculator")
	w.Resize(fyne.NewSize(500, 280))

	w.SetContent(
		calcContainer)
	w.Show()
}

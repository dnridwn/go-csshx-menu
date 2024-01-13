package main

import "github.com/rivo/tview"

func CreateModal(text string, buttons []string, buttonDoneHandler func(buttonIndex int, buttonLabel string)) *tview.Modal {
	modal := tview.NewModal()
	modal.SetText(text)
	modal.AddButtons(buttons)
	modal.SetDoneFunc(buttonDoneHandler)

	return modal
}

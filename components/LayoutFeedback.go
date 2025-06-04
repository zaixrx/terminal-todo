package components

import (
	"todo/shared"

	"github.com/rivo/tview"
)

func RenderFeedbackLayout(state *shared.AppState, feedback string, callback func()) {
	state.Pages.AddPage("feedback", tview.NewModal().
	SetText(feedback).
	AddButtons([]string{"OK"}).
	SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if callback != nil {
			callback()
		} else {
			state.Pages.RemovePage("feedback")
		}
	}), true, true)
}

func RenderErrorLayout(state *shared.AppState, feedback string) {
	state.Pages.AddPage("error", tview.NewModal().
	SetText(feedback).
	AddButtons([]string{"OK"}).
	SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		state.Pages.RemovePage("error")
	}), true, true)
}
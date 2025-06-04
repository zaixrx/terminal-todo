package components

import (
	"fmt"
	"todo/shared"
)

func RenderPage(name string, state *shared.AppState) {
	var (
		Page shared.Page
		present bool
	)

	Page, present = state.PagesMap[name]

	if !present {
		panic(fmt.Sprintf("%s doesn't exist in pages", name))
	}

	state.Pages.RemovePage(name)
	state.Pages.AddPage(name, Page(state), true, false)
	state.Pages.SwitchToPage(name)
}
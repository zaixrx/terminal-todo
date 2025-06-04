package shared

import (
	"github.com/rivo/tview"
)

const (
	PageViewTodos   = "viewTodo"
	PageCreateTodo  = "createTodo"
	PageSaveTodos   = "save"
)

type Page func(*AppState) tview.Primitive

type Todo struct {
	Text string
	Done bool
}

type AppState struct {
	App      *tview.Application
	Pages    *tview.Pages
	Todos    []*Todo
	PagesMap map[string]Page
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
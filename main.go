package main

import (
	"encoding/json"
	"os"
	"todo/components"
	"todo/shared"

	"github.com/rivo/tview"
)

func MustLoadTodos() []*shared.Todo  {
	var (
		path string = ""
		todos []*shared.Todo
	)

	if len(os.Args) >= 2 {
		path = os.Args[1]
	}

	if (path != "") {
		byt, err := os.ReadFile(path)
		shared.Check(err)

		err = json.Unmarshal(byt, &todos)
		shared.Check(err)
	} else {
		todos = make([]*shared.Todo, 0)
	}

	return todos
}

func MustInitializeState(state *shared.AppState) {
	state.App = tview.NewApplication()
	state.Pages = tview.NewPages()
	state.Todos = MustLoadTodos()
	state.PagesMap = map[string]shared.Page{
		shared.PageViewTodos: components.ViewTodos,
		shared.PageCreateTodo: components.CreateTodo,
		shared.PageSaveTodos: components.SaveTodos,
	}
}

func main() {
	var state shared.AppState;
 	MustInitializeState(&state)

	for name, Page := range state.PagesMap {
		state.Pages.AddPage(name, Page(&state), true, name == shared.PageViewTodos)
	}

	if err := state.App.SetRoot(state.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
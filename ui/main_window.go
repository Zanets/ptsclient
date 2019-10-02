package ui

import {
	"github.com/rivo/tview"
}

type main_window struct {
	app *tview.Application
	layout *tview.Grid
	sidebar *tview.Box
	app_content *tview.Box
	[] *app_ui
}

func NewMainWindow() *mainWindow {
	ins := main_window{}
	ins.sidebar := tview.NewBox().SetBorder(true).SetTitle("APP")
	ins.app_content := tview.NewBox().SetBorder(true).SetTitle("")
	ins.layout := tview.NewGrid().SetColumns(20, 0)
	ins.layout.AddItem(box1, 0, 0, 1, 1, 0, 0, false)
	ins.layout.AddItem(box2, 0, 1, 1, 1, 0, 0, false)
	ins.app := tview.NewApplication()
	ins.app.SetRoot(ins.layout, true)

	return &ins
}

func (this *main_window) Start() {
	if err := this.app.Run(); err != nil {
		panic(err)
	}
}

func (this *main_window) Stop() {
	if err := this.app.Stop(); err != nil {
		panic(err)
	}
}

func (this *main_window) AddApp() {


}


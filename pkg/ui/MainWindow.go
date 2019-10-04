package ui

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell"
)

type MainWindow struct {
	rootapp *tview.Application
	layout *tview.Grid
	sidebar *tview.List
	app *tview.TextView
	apps []APP_UI
}

func NewMainWindow() *MainWindow {
	ins := MainWindow{}
	ins.sidebar = tview.NewList()
	ins.app = tview.NewTextView()
	ins.layout = tview.NewGrid().SetColumns(20, 0)
	ins.layout.AddItem(ins.sidebar, 0, 0, 1, 1, 0, 0, true)
	ins.layout.AddItem(ins.app, 0, 1, 1, 1, 0, 0, false)
	ins.rootapp = tview.NewApplication()
	ins.rootapp.SetRoot(ins.layout, true)

	return &ins
}

func (this *MainWindow) StartSideBar() {

	this.sidebar.ShowSecondaryText(false)

	for i, app := range this.apps {
		this.sidebar.AddItem(app.GetName(), "", rune(i), nil)
	}
	
	this.SetSidebarEvent()
}

func (this *MainWindow) StartAPP() {
	this.SetAPPEvent()
}

func (this *MainWindow) SetSidebarEvent() {
	this.sidebar.SetSelectedFunc(func (i int, mainText string, secText string, c rune) {
		this.FocusOnAPP()
		// BUG: this line not work
		this.app.SetText("Loading...")
		this.app.SetText(this.apps[i].GetContent())
	})
}

func (this *MainWindow) SetAPPEvent() {
	this.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			this.FocusOnSidebar()
		}
		return event
	})
}

func (this *MainWindow) FocusOnSidebar() {
	this.rootapp.SetFocus(this.sidebar)
}

func (this *MainWindow) FocusOnAPP() {
	this.rootapp.SetFocus(this.app)
}

func (this *MainWindow) Start() {
	this.StartSideBar()
	this.StartAPP()
	if err := this.rootapp.Run(); err != nil {
		panic(err)
	}
}

func (this *MainWindow) Stop() {
	this.rootapp.Stop()
}

func (this *MainWindow) AddApp(app APP_UI) {
	this.apps = append(this.apps, app)
}


package main

import "github.com/Zanets/tclient/pkg/ui"
import "github.com/Zanets/tclient/pkg/app"

func main() {
	mainWindow := *ui.NewMainWindow()
	mainWindow.AddApp(app.APP_pts{})
	mainWindow.Start()
}

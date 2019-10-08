package ui
import (
	"github.com/rivo/tview"
)

type APP_UI interface {
	GetName() string
	GetContent() (*tview.Grid, error)
}

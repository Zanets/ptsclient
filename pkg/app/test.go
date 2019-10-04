package app

import "github.com/Zanets/tclient/pkg/ui"

type APP_test struct {
	ui.APP_UI
}

func (this APP_test) GetName() string {
	return "test"
} 

func (this APP_test) GetContent() string {
	return "test"
}


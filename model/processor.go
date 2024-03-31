package model

import "github.com/a-h/templ"

type Processor struct {
	Name         string
	Price        string
	Model        string
	SocketType   string
	Architecture string
	Power        string
	Frequency    string
	CoreNumber   string
	Link         templ.SafeURL
}

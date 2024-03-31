package model

import "github.com/a-h/templ"

type Accessory struct {
	Name  string
	Price string
	Link  templ.SafeURL
}

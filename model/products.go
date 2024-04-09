package model

import "github.com/a-h/templ"

type Products map[string][]Items

type Data struct {
	SourceList  []string
	ProductType []string
	Items       map[string]map[string]Items
}

type Items struct {
	Name       string
	FieldsList []string
	Items      []Item
}

type Item struct {
	Name   string
	Price  string
	URL    templ.SafeURL
	Fields map[string]string
}

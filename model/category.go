package model

type Field struct {
	Name string
}

type Category struct {
	Name   string
	Fields []Field
}

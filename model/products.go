package model

import (
	"github.com/a-h/templ"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Products map[string][]Items

type Data struct {
	SourceList  []string
	ProductType []string
	Items       map[string]map[string]Items
}

type Items struct {
	Name       string
	Items      []Item
	FieldsList []string
}

type Item struct {
	Name   string
	Price  currency.Amount
	URL    templ.SafeURL
	Fields map[string]string
}

func PriceToString(price currency.Amount) string {
	printer := message.NewPrinter(language.French)
	v := currency.NarrowSymbol(price)
	return printer.Sprint(v)
}

package model

import (
	"fmt"
	"testing"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func TestCurrency(t *testing.T) {
	c := currency.EUR.Amount(3.5)

	fmt.Println(c.Currency().String())
	mp := message.NewPrinter(language.French)
	format := currency.NarrowSymbol
	v := format(c)
	mp.Println(v)
}

// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "goth/model"

/*
	<option value="us">amazon.com</option>
	<option value="uk">amazon.co.uk</option>
	<option value="de">amazon.de</option>
	<option value="es">amazon.es</option>
	<option value="ca">amazon.ca</option>
	<option value="au">amazon.com.au</option>
	<option value="in">amazon.in</option>
	<option value="se">amazon.se</option>

*/

func sideBar(categories []model.Category) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"sidebar\"><fieldset id=\"locales\"><legend for=\"products-select\">Source</legend> <select name=\"data-source\" id=\"data-select\"><option value=\"fr\">amazon.fr</option></select></fieldset><fieldset id=\"locales\"><legend for=\"products-select\">Products</legend> <select name=\"products-source\" id=\"products-select\"><option value=\"accessories\">accessories</option></select></fieldset><ul><li><a href=\"/about\">About</a></li></ul></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
package mainpage

import (
	"bytes"
	"fmt"
	"html/template"
)

// ToHTML represents mainpage in HTML format for the handler
func ToHTML(bodyText string) string {
	tpl := template.Must(template.ParseFiles("mainpage/mainPage.html"))
	var b bytes.Buffer
	tpl.Execute(&b, template.HTML(bodyText))
	fmt.Println(b.String())
	return b.String()
}

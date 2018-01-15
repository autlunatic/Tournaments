package pairings

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

// ToHTML renders the Pairing List to a HTML Page
func ToHTML(p []P) string {
	fmt.Println("ASDF")
	tpl := template.Must(template.ParseFiles("PairingsList.html"))
	outP := []P{
		{1, 2, -4, -1, 0},
		{3, 4, -4, -2, 0},
		{5, 6, -4, -3, 0},
		{7, 8, -4, -4, 0},
	}
	fmt.Println("ASDF")
	var b bytes.Buffer
	tpl.Execute(&b, outP)
	tpl.Execute(os.Stdout, outP)
	return b.String()
}

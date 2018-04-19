package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type sage struct {
	Name  string
	Motto string
	Admin bool
}

func main() {

	g := sage{
		Name:  "Gandhi",
		Motto: "Peace",
		Admin: true,
	}

	b := sage{
		Name:  "Buddha",
		Motto: "Transcend",
		Admin: false,
	}

	sages := []sage{b, g}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}

}

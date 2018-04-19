package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	//you can access these by field name
	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Peace conquers all",
	}

	err := tpl.Execute(os.Stdout, gandhi)
	if err != nil {
		log.Fatalln(err)
	}
}

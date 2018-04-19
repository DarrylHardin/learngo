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

	buddha := sage{
		Name:  "Buddha",
		Motto: "Transcend this life",
	}

	mlk := sage{
		Name:  "MLK",
		Motto: "All life has value",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love your neighbor as yourself",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "Overcome evil with good, not evil with evil",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

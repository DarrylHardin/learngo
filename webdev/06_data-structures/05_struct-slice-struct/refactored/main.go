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

type car struct {
	Make  string
	Model string
	Year  int
}

// type items struct {
// 	Wisdom    []sage
// 	Transport []car
// }

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

	f := car{
		Make:  "Ford",
		Model: "Focus",
		Year:  2011,
	}

	c := car{
		Make:  "VW",
		Model: "Bug",
		Year:  1990,
	}

	cars := []car{f, c}
	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}

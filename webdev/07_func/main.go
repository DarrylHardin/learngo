package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacture string
	Model       string
	Year        int
}

var fm = template.FuncMap{
	"uc": strings.ToUpper, //all text uppercase
	"ft": firstThree,      //calls func firstThree
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s //returns string trimmed down to first 3 letters
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "Transcend",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Peace",
	}

	c := car{
		Manufacture: "Ford",
		Model:       "Focus",
		Year:        2015,
	}

	v := car{
		Manufacture: "VW",
		Model:       "Bug",
		Year:        1990,
	}

	sages := []sage{b, g}
	cars := []car{c, v}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

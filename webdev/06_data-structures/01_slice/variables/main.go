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

func main() {
	sages := []string{"Gandhi", "MLK", "Buhdda", "Jesus"}
	//code does not change here, but changes in html to specify key/value pairs
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

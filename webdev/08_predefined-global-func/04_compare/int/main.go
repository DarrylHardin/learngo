package main

import (
	"log"
	"os"
	"text/template"
)

var tlp *template.Template

func init() {
	tlp = template.Must(template.ParseFiles("tlp.gohtml"))
}

func main() {

	g1 := struct {
		Score1 int
		Score2 int
	}{
		1,
		2,
	}

	err := tlp.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}

}

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

type area struct {
	Hotels []hotel
	Region string
}

type hotel struct {
	Name   string
	Street string
	City   string
	State  string
	Zip    string
}

func main() {
	area := []area{
		{
			Region: "North",
			Hotels: []hotel{
				{"Mariat", "123", "Tallahassee", "FL", "32301"},
				{"Hotel 8", "555", "Leon", "FL", "33201"},
			},
		},
		{
			Region: "South",
			Hotels: []hotel{
				{"Sleepy Hotel", "667", "Miami", "FL", "34014"},
				{"Mainstay", "999", "Miami", "FL", "34014"},
			},
		},
	}
	err := tpl.Execute(os.Stdout, area)
	if err != nil {
		log.Fatalln(err)
	}
}

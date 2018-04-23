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

type mdish struct {
	Meat   string
	Veggie string
	Side   string
}

type dessert struct {
	Dairy string
}

type appetizer struct {
	Greasy string
}

type menu struct {
	Time      string
	Mdish     []mdish
	Dessert   []dessert
	Appetizer []appetizer
}

type restaurant struct {
	Name      string
	Breakfast menu
	Lunch     menu
	Dinner    menu
}

func main() {

	restaurants := []restaurant{
		{
			Name: "Happy Hammy",
			Breakfast: menu{
				Time:      "9 am",
				Mdish:     []mdish{{"Omelette", "", "Hasbrowns"}, {"Scrambled Eggs", "", "Hashbrowns"}},
				Dessert:   []dessert{{"Ice Cream"}, {"Cheese Cake"}},
				Appetizer: []appetizer{{"None"}},
			},
			Lunch: menu{
				Time:      "12 pm",
				Mdish:     []mdish{{"Steak", "", "Mashed Potatoes"}, {"Chicken Fried Steak", "", "Fries"}},
				Dessert:   []dessert{{"Cheese Cake"}},
				Appetizer: []appetizer{{"Bloomin Onion"}},
			},
			Dinner: menu{
				Time:      "4 pm",
				Mdish:     []mdish{{"Steak", "", "Mashed Potatoes"}},
				Dessert:   []dessert{{"Cheese Cake"}},
				Appetizer: []appetizer{{"Bloomin Onion"}},
			},
		},
		{
			Name: "Veggy Lover",
			Breakfast: menu{
				Time:      "9 am",
				Mdish:     []mdish{{"", "Veggie Eggs", "Hasbrowns"}, {"", "Veggie Scrambled Eggs", "Hashbrowns"}},
				Dessert:   []dessert{{"Ice Cream"}, {"Cheese Cake"}},
				Appetizer: []appetizer{{"None"}},
			},
			Lunch: menu{
				Time:      "12 pm",
				Mdish:     []mdish{{"", "Veggie Steak", "Mashed Potatoes"}, {"", "Veggie Chicken Fried Steak", "Fries"}},
				Dessert:   []dessert{{"Cheese Cake"}},
				Appetizer: []appetizer{{"Bloomin Onion"}},
			},
			Dinner: menu{
				Time:      "4 pm",
				Mdish:     []mdish{{"", "Steak", "Mashed Potatoes"}},
				Dessert:   []dessert{{"Cheese Cake"}},
				Appetizer: []appetizer{{"Bloomin Onion"}},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Fatalln(err)
	}
}

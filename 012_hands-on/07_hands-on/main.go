package main

import (
	"log"
	"os"
	"text/template"
)

type dish struct {
	Name  string
	Price float64
}

type meal struct {
	MealType string
	Dishes   []dish
}

type restaurant struct {
	RestaurantName string
	Meals          []meal
}

type menu []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := menu{
		restaurant{
			RestaurantName: "Cafe Alpha",
			Meals: []meal{
				meal{
					MealType: "Breakfast",
					Dishes: []dish{
						dish{
							"Bacon and Eggs",
							6.99,
						},
						dish{
							"Grilled Cheese",
							4.99,
						},
					},
				},
				meal{
					MealType: "Lunch",
					Dishes: []dish{
						dish{
							"Fish and Chips",
							12.99,
						},
						dish{
							"Fried Rice",
							11.99,
						},
					},
				},
				meal{
					MealType: "Dinner",
					Dishes: []dish{
						dish{
							"Pad Thai",
							16.99,
						},
						dish{
							"Black Pepper Steak",
							22.99,
						},
					},
				},
			},
		},
		restaurant{
			RestaurantName: "Cafe Beta",
			Meals: []meal{
				meal{
					MealType: "Breakfast",
					Dishes: []dish{
						dish{
							"Buttered Toast",
							3.99,
						},
					},
				},
				meal{
					MealType: "Lunch",
					Dishes: []dish{
						dish{
							"Greek Salad",
							8.99,
						},
					},
				},
				meal{
					MealType: "Dinner",
					Dishes: []dish{
						dish{
							"Mystery Meat",
							25.99,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}

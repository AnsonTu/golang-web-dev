package main

import (
	"log"
	"os"
	"text/template"
)

type hotels []hotel

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			"California Hotel", "123 Hotel Road", "Oakland", "12345", "Southern",
		},
		hotel{
			"California Hotel 2", "124 Hotel Road", "Oakland", "12345", "Northern",
		},
	}
	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}

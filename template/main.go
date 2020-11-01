package main

import (
	"log"
	"os"
	"time"
	"text/template"
)

type Computer struct {
	Brand     string
	OS        string
	Age       time.Time
	Color     string
	Memory    int
	Processor string
}

func main() {
	computers := fillComputers()

	var t = `{{- range . -}}
-------------------------------------
Brand: {{ .Brand }}
OS: {{ .OS }}
Age: {{ .Age }}
-------------------------------------
{{ end -}}
`

	c, err := template.New("computers").Parse(t)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Execute(os.Stdout, computers); err != nil {
		log.Fatal(err)
	}

}

func fillComputers() []Computer {
	var computers []Computer = []Computer{
		{
			Brand:     "Apple",
			OS:        "Catalina",
			Age:       time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC),
			Color:     "Silver",
			Memory:    16,
			Processor: "i7",
		},
		{
			Brand:     "Dell",
			OS:        "Windows 10",
			Age:       time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC),
			Color:     "Black",
			Memory:    16,
			Processor: "i5",
		},
	}
	return computers
}

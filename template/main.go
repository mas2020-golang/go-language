package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

type Computer struct {
	Brand       string
	OS          string
	Age         time.Time
	Color       string
	Memory      int
	Processor   string
	Accessories map[string]int
}

func main() {
	computers := fillComputers()
	fmt.Println("--Range example:")
	var t string
	//t = rangeExample()
	//t = indexExample()
	t = ifExample()
	execTemplate(&computers, t)
}

func execTemplate(computers *[]Computer, t string){
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
			Brand:       "Apple",
			OS:          "Catalina",
			Age:         time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC),
			Color:       "Silver",
			Memory:      16,
			Processor:   "i7",
			Accessories: map[string]int{"mouse": 12, "keyboard": 20},
		},
		{
			Brand:       "Dell",
			OS:          "Windows 10",
			Age:         time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC),
			Color:       "Black",
			Memory:      16,
			Processor:   "i5",
			Accessories: map[string]int{"mouse": 20, "screen": 130},
		},
		{
			Brand:       "Asus",
			OS:          "Windows 10",
			Age:         time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC),
			Color:       "Silver",
			Memory:      16,
			Processor:   "i7 6a",
			Accessories: map[string]int{"mouse": 30, "screen": 120, "audio": 200},
		},
	}
	return computers
}

/*
Example using the key of a map and a check to its value
 */
func indexExample() string {
	var t = `{{- range . -}}
-------------------------------------
Brand: {{ .Brand }}
OS: {{ .OS }}
Age: {{ .Age }}
Mouse: {{ .Accessories.mouse }}
	{{- if $screen := index .Accessories "screen" }}
Screen: {{ $screen }}
	{{- end }}
-------------------------------------
{{ end -}}
`
	return t
}

/*
Shows how to:
- define variables and use the values forward
- check conditions of two defined variables
*/
func ifExample() string {
	var t = `{{- range . -}}
-------------------------------------
Brand: {{ .Brand }}
OS: {{ .OS }}
Age: {{ .Age }}
Mouse: {{ .Accessories.mouse }}
	{{- /* definition of $screen and $mouse as variables */ -}}
  {{- $screen := index .Accessories "screen" }}
	{{- $mouse := index .Accessories "mouse" }}
	{{- if and $screen $mouse }}
Screen + mouse: {{ $screen }} + {{ $mouse }}
	{{- end }}
-------------------------------------
{{ end -}}
`
	return t
}

/*
Range example.
*/
func rangeExample() string {
	var t = `{{- range . -}}
-------------------------------------
Brand: {{ .Brand }}
OS: {{ .OS }}
Age: {{ .Age }}
-------------------------------------
{{ end -}}
`

	return t
}

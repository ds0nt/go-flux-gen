package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type collection []map[string]string

func parseCollection(name string, object map[string]string) {
	for v := range object {
		fmt.Printf("%s\n", v)
	}
}

func readDbJson() {
	content, err := ioutil.ReadFile("db.json")
	if err != nil {
		fmt.Errorf("Error:", err)
	}

	var data = map[string]collection{}

	err = json.Unmarshal(content, &data)
	if err != nil {
		fmt.Errorf("Json Parse Error: ", err)
	}
	for k, v := range data {
		fmt.Printf("Collection: %s\n", k)
		if len(v) != 0 {
			parseCollection(k, v[0])
		}
	}
}

func main() {
	readDbJson()
	//
	// 	// Define a template.
	// 	const letter = `
	// Dear {{.Name}},
	// {{if .Attended}}
	// It was a pleasure to see you at the wedding.{{else}}
	// It is a shame you couldn't make it to the wedding.{{end}}
	// {{with .Gift}}Thank you for the lovely {{.}}.
	// {{end}}
	// Best wishes,
	// Josie
	// `
	//
	// 	// Prepare some data to insert into the template.
	// 	type Recipient struct {
	// 		Name, Gift string
	// 		Attended   bool
	// 	}
	// 	var recipients = []Recipient{
	// 		{"Aunt Mildred", "bone china tea set", true},
	// 		{"Uncle John", "moleskin pants", false},
	// 		{"Cousin Rodney", "", false},
	// 	}
	//
	// 	// Create a new template and parse the letter into it.
	// 	t := template.Must(template.New("letter").Parse(letter))
	//
	// 	// Execute the template for each recipient.
	// 	for _, r := range recipients {
	// 		err := t.Execute(os.Stdout, r)
	// 		if err != nil {
	// 			log.Println("executing template:", err)
	// 		}
	// 	}
	//
}

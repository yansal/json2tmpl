package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"text/template"
)

func main() {
	log.SetFlags(0)
	if err := z(); err != nil {
		log.Fatal(err)
	}
}

func z() error {
	if len(os.Args) != 2 {
		return errors.New("usage: json2tmpl z.tmpl < z.json")
	}
	tmpl, err := template.ParseFiles(os.Args[1])
	if err != nil {
		return err
	}
	jsondec := json.NewDecoder(os.Stdin)
	jsondec.UseNumber()
	var v interface{}
	if err := jsondec.Decode(&v); err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, v)
}

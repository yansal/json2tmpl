package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/mgutz/ansi"
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

	tmplname := filepath.Base(os.Args[1])
	tmpl, err := template.New(tmplname).Funcs(template.FuncMap{
		"color": func(color string, arg interface{}) string {
			return ansi.Color(fmt.Sprint(arg), color)
		},
	}).ParseFiles(os.Args[1])
	if err != nil {
		return err
	}
	jsondec := json.NewDecoder(os.Stdin)
	jsondec.UseNumber()
	var v interface{}
	if err := jsondec.Decode(&v); err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(os.Stdout, tmplname, v)
}

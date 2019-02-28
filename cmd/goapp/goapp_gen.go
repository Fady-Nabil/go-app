// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Create("goapp_tmpl.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "package main")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "// Code generated by go generate; DO NOT EDIT.")

	gen := []struct {
		Var      string
		Filename string
	}{
		{Var: "goappOfflineJS", Filename: "goapp.js"},
	}

	for _, g := range gen {
		b, err := ioutil.ReadFile(g.Filename)
		if err != nil {
			panic(err)
		}

		fmt.Fprintln(f)
		fmt.Fprintln(f)
		fmt.Fprintf(f, "const %s = `%s`", g.Var, b)
	}
}
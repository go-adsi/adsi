package main

import (
	"flag"
	"log"
	"os"

	"github.com/go-adsi/adsi"
	"github.com/go-adsi/adsi/adspath"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	var q = flag.Arg(0)
	query, err := adspath.Parse(q) // This will correct namespace capitalization
	if err != nil {
		log.Fatalf("Invalid Path: %v\n", err)
	}

	obj, err := adsi.Open(query.String())
	if err != nil {
		log.Fatalf("Unable to open object: %v\n", err)
	}
	defer obj.Close()

	log.Printf("Query:  %v\n", query.String())
	log.Println("--------")

	show(obj)
}

func show(obj *adsi.Object) {
	name, _ := obj.Name()
	class, _ := obj.Class()
	guid, _ := obj.GUID()
	path, _ := obj.Path()
	parent, _ := obj.Parent()
	schema, _ := obj.Schema()
	log.Printf("Name:   %v\n", name)
	log.Printf("Class:  %v\n", class)
	log.Printf("GUID:   %v\n", guid)
	log.Printf("Path:   %v\n", path)
	log.Printf("Parent: %v\n", parent)
	log.Printf("Schema: %v\n", schema)
}

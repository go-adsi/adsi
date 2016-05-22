package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/adsi.v0"
	"gopkg.in/adsi.v0/api"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	var domain = flag.Arg(0)

	url := dfsrURL(domain)
	root := prepareObject(url)
	defer root.Close()

	rglist, err := fetchChildren(root)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(rglist); i++ {
		log.Printf("[%3d] Name: %v\n", i, rglist[i])
	}
}

func prepareObject(url string) *adsi.Object {
	ds, err := adsi.NewDirectoryService("")
	if err != nil {
		log.Fatal(err)
	}
	defer ds.Close()

	obj, err := ds.OpenObject(url, "", "", api.ADS_READONLY_SERVER|api.ADS_SECURE_AUTHENTICATION|api.ADS_USE_SEALING)
	if err != nil {
		log.Fatal(err)
	}

	return obj
}

func fetchChildren(parent *adsi.Object) (groups []string, err error) {
	c, err := parent.ToContainer()
	if err != nil {
		return
	}
	defer c.Close()

	iter, err := c.Children()
	if err != nil {
		log.Fatal(err)
	}
	defer iter.Close()

	i := 0
	for child, err := iter.Next(); err == nil; child, err = iter.Next() {
		defer child.Close()
		name, err := child.Name()
		if err != nil {
			log.Fatal(err)
		}
		guid, err := child.GUID()
		if err != nil {
			log.Fatal(err)
		}
		groups = append(groups, fmt.Sprintf("%-45s %s", name, guid))
		//show(child, fmt.Sprintf("%v: ", i))
		i++
	}

	return
}

func show(obj *adsi.Object, prefix string) {
	name, _ := obj.Name()
	class, _ := obj.Class()
	guid, _ := obj.GUID()
	path, _ := obj.Path()
	parent, _ := obj.Parent()
	schema, _ := obj.Schema()

	log.Println("--------")
	log.Printf("%sName: %v\n", prefix, name)
	log.Printf("%sClass: %v\n", prefix, class)
	log.Printf("%sGUID: %v\n", prefix, guid)
	log.Printf("%sPath: %v\n", prefix, path)
	log.Printf("%sParent: %v\n", prefix, parent)
	log.Printf("%sSchema: %v\n", prefix, schema)
}

func makeDN(attribute string, components ...string) string {
	if attribute != "" {
		for i := 0; i < len(components); i++ {
			components[i] = attribute + "=" + components[i]
		}
	}
	return strings.Join(components, ",")
}

func dfsrURL(domain string) string {
	cn := makeDN("cn", "DFSR-GlobalSettings", "System")
	dc := makeDN("dc", strings.Split(domain, ".")...)
	dn := strings.Join([]string{cn, dc}, ",")
	protocol := "LDAP"
	return protocol + "://" + dn
}

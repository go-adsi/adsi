package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/adsi.v0"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	var domain = flag.Arg(0)

	root, err := adsi.Open(dfsrPath(domain))
	if err != nil {
		log.Fatal(err)
	}
	defer root.Close()

	rglist, err := fetchChildren(root)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(rglist); i++ {
		log.Printf("[%3d] Name: %v\n", i, rglist[i])
	}
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
		i++
	}

	return
}

func dfsrPath(domain string) string {
	cn := makeDN("cn", "DFSR-GlobalSettings", "System")
	dc := makeDN("dc", strings.Split(domain, ".")...)
	dn := strings.Join([]string{cn, dc}, ",")
	protocol := "LDAP"
	return protocol + "://" + dn
}

func makeDN(attribute string, components ...string) string {
	if attribute != "" {
		for i := 0; i < len(components); i++ {
			components[i] = attribute + "=" + components[i]
		}
	}
	return strings.Join(components, ",")
}

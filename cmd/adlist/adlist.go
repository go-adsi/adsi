package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-adsi/adsi"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	var domain = flag.Arg(0)

	root, err := adsi.Open(domainPath(domain))
	if err != nil {
		log.Fatal(err)
	}
	defer root.Close()

	adlist, err := fetchChildren(root, 0)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(adlist); i++ {
		log.Printf("[%4d] %v\n", i, adlist[i])
	}
}

func fetchChildren(parent *adsi.Object, depth int) (groups []string, err error) {
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
		class, err := child.Class()
		if err != nil {
			log.Fatal(err)
		}

		cdepth := depth + 1

		plen := 55 - len(name) - depth*2
		if plen < 0 {
			plen = 0
		}
		padding := strings.Repeat(" ", plen)
		groups = append(groups, fmt.Sprintf("%s%s %s %s %s", strings.Repeat("  ", cdepth), name, padding, guid, class))

		sub, err := fetchChildren(child, cdepth)
		if err != nil {
			log.Fatal(err)
		}

		groups = append(groups, sub...)

		i++
	}

	return
}

func domainPath(domain string) string {
	dn := makeDN("dc", strings.Split(domain, ".")...)
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

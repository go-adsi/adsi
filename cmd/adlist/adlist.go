package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/go-ole/go-ole"

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

	stop := make(chan struct{})

	go comShim(stop)

	url := rootURL(domain)

	root := prepareRoot(url)
	defer root.Close()

	rglist, err := fetchChildren(root, 0)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(rglist); i++ {
		log.Printf("[%4d] %v\n", i, rglist[i])
	}

	root.Close()

	close(stop)
}

func prepareRoot(url string) *adsi.Object {
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

func rootURL(domain string) string {
	dn := makeDN("dc", strings.Split(domain, ".")...)
	protocol := "LDAP"
	return protocol + "://" + dn
}

func comShim(until chan struct{}) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED); err != nil {
		oleerr := err.(*ole.OleError)
		// S_FALSE           = 0x00000001 // CoInitializeEx was already called on this thread
		if oleerr.Code() != ole.S_OK && oleerr.Code() != 0x00000001 {
			return err
		}
	}

	defer ole.CoUninitialize()

	<-until

	return nil
}

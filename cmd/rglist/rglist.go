package main

import (
	"flag"
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

	root, err := adsi.Open(dfsrRootPath(domain))
	if err != nil {
		log.Fatal(err)
	}
	defer root.Close()

	groups, err := fetchGroups(root, domain)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(groups); i++ {
		group := groups[i]
		log.Printf("[%3d] Group: %-50s ID: %v\n", i, group.Name, group.ID)
		for _, folder := range group.Folders {
			log.Printf("        Folder: %-47s ID: %v\n", folder.Name, folder.ID)
		}
		for _, member := range group.Members {
			log.Printf("        Member: %-47s ID: %v   Computer: %v\n", member.Name, member.ID, member.Computer)
			for _, conn := range member.Connections {
				log.Printf("          Connection: %s\n", conn.From)
			}
		}
	}
}

func fetchGroups(root *adsi.Object, domain string) (groups []*ReplicationGroup, err error) {
	c, err := root.ToContainer()
	if err != nil {
		return
	}
	defer c.Close()

	iter, err := c.Children()
	if err != nil {
		log.Fatal(err)
	}
	defer iter.Close()

	for group, err := iter.Next(); err == nil; group, err = iter.Next() {
		defer group.Close()

		name, err := group.Name()
		if err != nil {
			log.Fatal(err)
		}
		guid, err := group.GUID()
		if err != nil {
			log.Fatal(err)
		}

		gc, err := group.ToContainer()
		if err != nil {
			log.Fatal(err)
		}
		defer gc.Close()

		topo, err := gc.Object("msDFSR-Topology", "cn=Topology")
		if err != nil {
			log.Fatal(err)
		}
		defer topo.Close()

		members, err := fetchMembers(topo)
		if err != nil {
			log.Fatal(err)
		}

		content, err := gc.Object("msDFSR-Content", "cn=Content")
		if err != nil {
			log.Fatal(err)
		}

		folders, err := fetchFolders(content)

		groups = append(groups, &ReplicationGroup{
			Name:    name,
			ID:      guid,
			Folders: folders,
			Members: members,
		})
	}

	return
}

func fetchMembers(topo *adsi.Object) (members []*Member, err error) {
	c, err := topo.ToContainer()
	if err != nil {
		return
	}
	defer c.Close()

	iter, err := c.Children()
	if err != nil {
		log.Fatal(err)
	}
	defer iter.Close()

	for member, err := iter.Next(); err == nil; member, err = iter.Next() {
		defer member.Close()

		name, err := member.Name()
		if err != nil {
			log.Fatal(err)
		}

		guid, err := member.GUID()
		if err != nil {
			log.Fatal(err)
		}

		comp, err := member.AttrString("msDFSR-ComputerReference")
		if err != nil {
			log.Fatal(err)
		}

		dn, err := member.Path()
		if err != nil {
			log.Fatal(err)
		}
		if proto := strings.Index(dn, "://"); proto >= 0 {
			if proto+3 < len(dn) {
				dn = dn[proto+3:]
			} else {
				dn = ""
			}
		}

		connections, err := fetchConnections(member)
		if err != nil {
			log.Fatal(err)
		}

		members = append(members, &Member{
			Name:        name,
			ID:          guid,
			Computer:    comp,
			DN:          dn,
			Connections: connections,
		})
	}

	return
}

func fetchFolders(content *adsi.Object) (folders []*ReplicationFolder, err error) {
	c, err := content.ToContainer()
	if err != nil {
		return
	}
	defer c.Close()

	iter, err := c.Children()
	if err != nil {
		log.Fatal(err)
	}
	defer iter.Close()

	for folder, err := iter.Next(); err == nil; folder, err = iter.Next() {
		defer folder.Close()

		name, err := folder.Name()
		if err != nil {
			log.Fatal(err)
		}

		guid, err := folder.GUID()
		if err != nil {
			log.Fatal(err)
		}

		folders = append(folders, &ReplicationFolder{
			Name: name,
			ID:   guid,
		})
	}

	return
}

func fetchConnections(member *adsi.Object) (connections []*Connection, err error) {
	c, err := member.ToContainer()
	if err != nil {
		return
	}
	defer c.Close()

	iter, err := c.Children()
	if err != nil {
		log.Fatal(err)
	}
	defer iter.Close()

	for conn, err := iter.Next(); err == nil; conn, err = iter.Next() {
		defer conn.Close()

		from, err := conn.AttrString("fromServer")
		if err != nil {
			log.Fatal(err)
		}
		if comma := strings.Index(from, ","); comma > 0 {
			from = from[0:comma]
		}

		guid, err := conn.GUID()
		if err != nil {
			log.Fatal(err)
		}

		connections = append(connections, &Connection{
			ID:   guid,
			From: from,
		})
	}

	return
}

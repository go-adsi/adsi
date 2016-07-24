package main

import "github.com/go-ole/go-ole"

type ReplicationGroup struct {
	Name    string
	ID      *ole.GUID
	Folders []*ReplicationFolder
	Members []*Member
}

type ReplicationFolder struct {
	Name string
	ID   *ole.GUID
}

type Member struct {
	Name        string
	ID          *ole.GUID
	Computer    string // Distinguished name of the computer
	DN          string
	Connections []*Connection
}

type Connection struct {
	ID   *ole.GUID
	From string // Distinguished name of source member in topology
}

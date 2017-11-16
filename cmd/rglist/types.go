package main

import "github.com/google/uuid"

type ReplicationGroup struct {
	Name    string
	ID      uuid.UUID
	Folders []*ReplicationFolder
	Members []*Member
}

type ReplicationFolder struct {
	Name string
	ID   uuid.UUID
}

type Member struct {
	Name        string
	ID          uuid.UUID
	Computer    string // Distinguished name of the computer
	DN          string
	Connections []*Connection
}

type Connection struct {
	ID   uuid.UUID
	From string // Distinguished name of source member in topology
}

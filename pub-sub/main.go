package main

import (
	"fmt"

	msg "sandbox/provider"
)

type Contacts struct {
	names []Name
}

type Name struct {
	FirstName string
	LastName  string
}

func (name *Name) FullName() string {
	return name.FirstName + " " + name.LastName
}

func (contacts *Contacts) ListAll() {
	for _, name := range contacts.names {
		fmt.Printf("%s \n", name.FullName())
	}
}

func main() {
	msg.SendMessage("test")

	name1 := Name{"Neo", "Sporin"}
	name2 := Name{"Theo", "Sporin"}

	contacts := Contacts{[]Name{name1, name2}}

	contacts.ListAll()
}

package main

import (
	"fmt"
)

// Person represents the generic parent type representing any kind of person
type Person struct {
	Name      string
	privateID string
}

// GetName is a public method that fetches a person's name
func (p Person) GetName() string {
	return p.Name
}

// GetPrivateID is a public method that sets a person's private ID
func (p *Person) SetPrivateID(id string) {
	p.privateID = id
}

// GetPrivateID is a public method that fetches a person's private ID
func (p Person) GetPrivateID() string {
	return p.privateID
}

// Student is a composite type that inherits Person fields and methods
type Student struct {
	Person
	Degrees map[string]float64
}

type nameGetter interface {
	GetName() string
}

type privateIDGetter interface {
	GetPrivateID() string
}

// People is an interface representing both Person & Student
type People interface {
	nameGetter
	privateIDGetter
}

func printPeople(people ...People) {
	for _, person := range people {
		fmt.Printf("My name is: %s\n", person.GetName())
		fmt.Printf("My private ID is: %s\n", person.GetPrivateID())
	}
}

func main() {
	p := Person{Name: "John"}
	p.SetPrivateID("id001")

	s := Student{
		Person: Person{Name: "Anonymous"},
		Degrees: map[string]float64{
			"Math": 8,
			"English": 9,
		},
	}
	s.Name = "Mike"
	s.SetPrivateID("id0002")

	printPeople(p, s)
}

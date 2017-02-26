package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
	Town string
}
type byName []Person
type byAge []Person
type byTown []Person

// sort need len, less and swap func
func (ps byName) Len() int {
	return len(ps)
}

func (ps byName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps byName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps byAge) Len() int {
	return len(ps)
}

func (ps byAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

func (ps byAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps byTown) Len() int {
	return len(ps)
}

func (ps byTown) Less(i, j int) bool {
	return ps[i].Town < ps[j].Town
}

func (ps byTown) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		{"Martin", 2, "Nantes"},
		{"Arthur", 9, "Toulouse"},
		{"Chloe", 11, "Lyon"},
		{"Luna", 8, "Lyon"},
	}
	sort.Sort(byName(kids))
	fmt.Println(kids)
	sort.Sort(byAge(kids))
	fmt.Println(kids)
	sort.Sort(byTown(kids))
	fmt.Println(kids)
}

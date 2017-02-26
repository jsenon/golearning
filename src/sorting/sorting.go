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

// for custom sort
// sort need len, less and swap func

type byName []Person
type byAge []Person
type byTown []Person

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

// All above could be deleted if we use new go 1.8

func main() {
	kids := []Person{
		{"Martin", 2, "Nantes"},
		{"Arthur", 9, "Toulouse"},
		{"Chloe", 11, "Lyon"},
		{"Luna", 8, "Lyon"},
	}

	// Bellow is for old sort

	sort.Sort(byName(kids))
	fmt.Println(kids)
	sort.Sort(byAge(kids))
	fmt.Println(kids)
	sort.Sort(byTown(kids))
	fmt.Println(kids)

	// Could be improve by sort slice but need go 1.8

	sort.Slice(kids, func(i, j int) bool { return kids[i].Name < kids[j].Name })
	fmt.Println("By Name quickly:", kids)
	sort.Slice(kids, func(i, j int) bool { return kids[i].Age < kids[j].Age })
	fmt.Println("By Age quickly:", kids)
	sort.Slice(kids, func(i, j int) bool { return kids[i].Town < kids[j].Town })
	fmt.Println("By Town quickly:", kids)

}

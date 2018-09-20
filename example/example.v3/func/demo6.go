package main

import "fmt"

type Cats struct {
	name           string
	scientificName string
	category       string
}

func New(name, scientificName, category string) Cats {
	return Cats{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func (cat *Cats) SetName (name string) {
	cat.name = name
}

func (cat Cats) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cats) Name () string {
	return cat.name
}

func (cat Cats) ScientifiName () string {
	return cat.scientificName
}

func (cat Cats) Category () string {
	return cat.category
}

func (cat Cats) String() string {
	return fmt.Sprintf("%s (category %s, name: %q)", cat.scientificName, cat.category, cat.name)
}

func main() {
	cat := New("litte dog", "American shorthair", "cat")
	cat.SetName("Moasor")
	fmt.Printf("The cat: %s\n", cat)

	cat.SetNameOfCopy("little pig")
	fmt.Printf("The cat: %s\n", cat)

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}

	_, ok := interface{}(cat).(Pet)
	fmt.Printf("Cat implements interface Pet: %v\n", ok)

	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*Cat implements interface Pet: %v\n", ok)
}

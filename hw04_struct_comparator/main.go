package main

import "fmt"

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   float32
	rate   float32
}

func (b *Book) SetFields(id int, title string, author string, year int, size float32, rate float32) {
	b.id = id
	b.title = title
	b.author = author
	b.year = year
	b.size = size
	b.rate = rate
}

func (b Book) ReadFields() string {
	return fmt.Sprintf("Book ID: %d; Title: %s; Author: %s; Year: %d; Size: %v; Rate: %v",
		b.id, b.title, b.author, b.year, b.size, b.rate)
}

type CompareBook struct {
	enum  string
	book1 Book
	book2 Book
}

func (cs CompareBook) Compare() bool {
	var result bool

	switch cs.enum {
	case "year":
		result = cs.book1.year > cs.book2.year
	case "size":
		result = cs.book1.size > cs.book2.size
	case "rate":
		result = cs.book1.rate > cs.book2.rate
	}

	return result
}

func main() {
	mybook1 := Book{}
	mybook1.SetFields(1, "super book", "David", 2020, 1.64, 5.1)

	strbook1 := mybook1.ReadFields()
	fmt.Println(strbook1)

	mybook2 := Book{}
	mybook2.SetFields(2, "great book", "Nikole", 2019, 2.72, 4.9)

	strbook2 := mybook2.ReadFields()
	fmt.Println(strbook2)

	cs1 := CompareBook{"year", mybook1, mybook2}
	cs2 := CompareBook{"size", mybook1, mybook2}
	cs3 := CompareBook{"rate", mybook1, mybook2}
	fmt.Println(`CompareBook "year":`, cs1.Compare())
	fmt.Println(`CompareBook "size":`, cs2.Compare())
	fmt.Println(`CompareBook "rate":`, cs3.Compare())
}

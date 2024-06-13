package hw04

import "fmt"

type fieldBook int

const (
	Year fieldBook = iota
	Size
	Rate
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   float32
	rate   float32
}

func (b Book) ID() int {
	return b.id
}

func (b Book) Title() string {
	return b.title
}

func (b Book) Author() string {
	return b.author
}

func (b Book) Year() int {
	return b.year
}

func (b Book) Size() float32 {
	return b.size
}

func (b Book) Rate() float32 {
	return b.rate
}

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size float32) {
	b.size = size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

type Comparator struct {
	fieldCompare fieldBook
}

func NewComparator(fieldCompare fieldBook) *Comparator {
	c := Comparator{fieldCompare}
	return &c
}

func (c Comparator) Compare(bookOne, bookTwo *Book) bool {
	var result bool

	switch c.fieldCompare {
	case Year:
		result = bookOne.year > bookTwo.year
	case Size:
		result = bookOne.size > bookTwo.size
	case Rate:
		result = bookOne.rate > bookTwo.rate
	}

	return result
}

func CompareBook() {
	mybook1 := Book{}
	mybook1.SetID(1)
	mybook1.SetTitle("super book")
	mybook1.SetAuthor("David")
	mybook1.SetYear(2020)
	mybook1.SetSize(1.64)
	mybook1.SetRate(5.1)

	fmt.Println(mybook1)
	fmt.Println("mybook1, title:", mybook1.title)

	mybook2 := Book{}
	mybook2.SetID(2)
	mybook2.SetTitle("great book")
	mybook2.SetAuthor("Nikole")
	mybook2.SetYear(2019)
	mybook2.SetSize(2.72)
	mybook2.SetRate(4.9)

	fmt.Println(mybook2)
	fmt.Println("mybook2, title:", mybook2.title)

	c1 := NewComparator(Year)
	c2 := NewComparator(Size)
	c3 := NewComparator(Rate)
	fmt.Println(`Compare book by field "year":`, c1.Compare(&mybook1, &mybook2))
	fmt.Println(`Compare book by field "size":`, c2.Compare(&mybook1, &mybook2))
	fmt.Println(`Compare book by field "rate":`, c3.Compare(&mybook1, &mybook2))
}

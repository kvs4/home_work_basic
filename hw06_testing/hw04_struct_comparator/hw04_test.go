package hw04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareEmpty(t *testing.T) {
	mybook1 := Book{}

	mybook2 := Book{}

	c1 := NewComparator(Year)
	assert.Equal(t, false, c1.Compare(&mybook1, &mybook2))
	c2 := NewComparator(Size)
	assert.Equal(t, false, c2.Compare(&mybook1, &mybook2))
	c3 := NewComparator(Rate)
	assert.Equal(t, false, c3.Compare(&mybook1, &mybook2))
}

func TestComparePositive(t *testing.T) {
	mybook1 := Book{}
	mybook1.SetID(1)
	mybook1.SetTitle("super book")
	mybook1.SetAuthor("David")
	mybook1.SetYear(2020)
	mybook1.SetSize(1.64)
	mybook1.SetRate(5.1)

	mybook2 := Book{}
	mybook2.SetID(2)
	mybook2.SetTitle("great book")
	mybook2.SetAuthor("Nikole")
	mybook2.SetYear(2019)
	mybook2.SetSize(2.72)
	mybook2.SetRate(4.9)

	c1 := NewComparator(Year)
	assert.Equal(t, true, c1.Compare(&mybook1, &mybook2))
	c2 := NewComparator(Size)
	assert.Equal(t, false, c2.Compare(&mybook1, &mybook2))
	c3 := NewComparator(Rate)
	assert.Equal(t, true, c3.Compare(&mybook1, &mybook2))
}

func TestStructBook(t *testing.T) {
	mybook := Book{}
	mybook.SetID(1)
	mybook.SetTitle("super book")
	mybook.SetAuthor("David")
	mybook.SetYear(2020)
	mybook.SetSize(1.64)
	mybook.SetRate(5.1)

	assert.Equal(t, 1, mybook.id)
	assert.Equal(t, "super book", mybook.title)
	assert.Equal(t, "David", mybook.author)
	assert.Equal(t, 2020, mybook.year)
	assert.Equal(t, float32(1.64), mybook.size)
	assert.Equal(t, float32(5.1), mybook.rate)

	assert.Equal(t, mybook.id, mybook.ID())
	assert.Equal(t, mybook.title, mybook.Title())
	assert.Equal(t, mybook.author, mybook.Author())
	assert.Equal(t, mybook.year, mybook.Year())
	assert.Equal(t, mybook.size, mybook.Size())
	assert.Equal(t, mybook.rate, mybook.Rate())
}

func TestNewCompator(t *testing.T) {
	c := Comparator{Year}
	assert.Equal(t, &c, NewComparator(Year))
}

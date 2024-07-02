package main

import (
	"encoding/json"
	"fmt"

	myproto "github.com/kvs4/home_work_basic/hw09_serialize/msg"
	"google.golang.org/protobuf/proto"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   float64
	Rate   float64
}

func (b *Book) MarshalJSON() ([]byte, error) {
	type Alias Book
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	})
}

func (b *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	return json.Unmarshal(data, &aux)
}

func SerializeBooksJSON(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

func DeserializeBooksJSON(data []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(data, &books)
	return books, err
}

func SerializeBooksProto(books *myproto.Books) ([]byte, error) {
	return proto.Marshal(books)
}

func DeserializeBooksProto(data []byte) (*myproto.Books, error) {
	var booksProto myproto.Books
	err := proto.Unmarshal(data, &booksProto)
	return &booksProto, err
}

func main() {
	book := Book{
		ID:     3,
		Title:  "my title3",
		Author: "my author3",
		Year:   2020,
		Size:   float64(4.04),
		Rate:   float64(2.2),
	}

	data, err := book.MarshalJSON()
	fmt.Println(data, err)

	books := []Book{
		{
			ID:     1,
			Title:  "my title",
			Author: "my author",
			Year:   2019,
			Size:   float64(1.05),
			Rate:   float64(5.3),
		},
		{
			ID:     2,
			Title:  "my title 2",
			Author: "my author 2",
			Year:   2018,
			Size:   float64(2.09),
			Rate:   float64(4.1),
		},
	}

	data, err = SerializeBooksJSON(books)
	fmt.Println(data, err)

	newbooks, err := DeserializeBooksJSON(data)
	fmt.Println(newbooks, err)

	// -----------------------------------
	booksProto := getWantedBooksProto1()

	dataProto, err := SerializeBooksProto(&booksProto)

	fmt.Println("This is books Proto:", dataProto, "error:", err)

	gotBooks, _ := DeserializeBooksProto(dataProto)

	fmt.Println("want value:", &booksProto)
	fmt.Println("got value:", gotBooks)
}

func getWantedBooksProto1() myproto.Books {
	book1 := myproto.Book{
		ID:     1,
		Title:  "my title",
		Author: "my author",
		Year:   2019,
		Size:   float64(1.05),
		Rate:   float64(5.3),
	}
	book2 := myproto.Book{
		ID:     2,
		Title:  "my title 2",
		Author: "my author 2",
		Year:   2018,
		Size:   float64(2.09),
		Rate:   float64(4.1),
	}

	return myproto.Books{Books: []*myproto.Book{&book1, &book2}}
}

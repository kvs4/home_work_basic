package main

import (
	"testing"

	myproto "github.com/kvs4/home_work_basic/hw09_serialize/msg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getWantedBooksJSON() []Book {
	return []Book{
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
}

func TestSerializeBookJSON(t *testing.T) {
	book := Book{
		ID:     3,
		Title:  "my title3",
		Author: "my author3",
		Year:   2020,
		Size:   float64(4.04),
		Rate:   float64(2.2),
	}

	data, err := book.MarshalJSON()
	require.NoError(t, err)

	gotBook := Book{}
	err = gotBook.UnmarshalJSON(data)

	require.NoError(t, err)
	assert.Equal(t, book, gotBook)
}

func TestSerializeBooks(t *testing.T) {
	books := getWantedBooksJSON()

	data, err := SerializeBooksJSON(books)
	require.NoError(t, err)

	gotBooks, err2 := DeserializeBooksJSON(data)

	require.NoError(t, err2)
	assert.Equal(t, books, gotBooks)
}

func getWantedBooksProto() myproto.Books {
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

func TestSerializeBooksProto(t *testing.T) {
	books := getWantedBooksProto()
	data, err := SerializeBooksProto(&books)
	require.NoError(t, err)

	gotBooks, err2 := DeserializeBooksProto(data)

	require.NoError(t, err2)

	// assert.Equal(t, books, *gotBooks)
	for i, book := range books.Books {
		assert.Equal(t, book.GetID(), gotBooks.Books[i].GetID())
		assert.Equal(t, book.GetAuthor(), gotBooks.Books[i].GetAuthor())
		assert.Equal(t, book.GetTitle(), gotBooks.Books[i].GetTitle())
		assert.Equal(t, book.GetYear(), gotBooks.Books[i].GetYear())
		assert.Equal(t, book.GetSize(), gotBooks.Books[i].GetSize())
		assert.Equal(t, book.GetRate(), gotBooks.Books[i].GetRate())
	}
}

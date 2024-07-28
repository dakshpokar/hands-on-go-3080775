// challenges/types-composite/begin/main.go
package main

import (
	"errors"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	auth  author
}

// define a library type to track a list of books
type library struct {
	books map[string][]book
}

// define addBook to add a book to the library
func (l *library) addBook(b book) {
	l.books[b.auth.name] = append(l.books[b.auth.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l *library) lookupByAuthor(auth string) ([]book, error) {
	b, ok := l.books[auth]
	if ok {
		return b, nil
	}
	return make([]book, 0), errors.New("Not Found!")
}

func main() {
	// create a new library
	l := library{books: map[string][]book{}}

	// add 2 books to the library by the same author
	b1 := book{title: "Golden Harbor", auth: author{name: "Daksh"}}
	l.addBook(b1)
	b2 := book{title: "Amazing Science", auth: author{name: "Daksh"}}
	l.addBook(b2)

	// add 1 book to the library by a different author
	b3 := book{title: "Awesome Science", auth: author{name: "Sanket"}}
	l.addBook(b3)

	// dump the library with spew
	spew.Dump(l)

	// lookup books by known author in the library
	books, err := l.lookupByAuthor("Sanket")
	if err == nil {
		spew.Dump(books)
	} else {
		fmt.Println("Error ala re!", err)
	}

	// print out the first book's title and its author's name
	if len(books) != 0 {
		fmt.Printf("title = %v, author = %v", books[0].title, books[0].auth.name)
	}
}

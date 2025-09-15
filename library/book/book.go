package book

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
}

func NewBook(title, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *Book) PrintInfo() {
	fmt.Printf("Titúlo: %s\n Autor: %s\n Páginas del Libro: %d\n", b.title, b.author, b.pages)
}

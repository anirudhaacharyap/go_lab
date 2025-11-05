package main

import "fmt"

type Book struct {
	title     string
	available bool
}

func (b *Book) borrow() {
	if b.available {
		b.available = false
		fmt.Println(b.title, "has been borrowed.")
	} else {
		fmt.Println(b.title, "is not available.")
	}
}

func (b *Book) returnBook() {
	if b.available {
		fmt.Println(b.title, "was not borrowed.")
	} else {
		b.available = true
		fmt.Println(b.title, "has been returned.")
	}
}

func (b *Book) checkAvailability() {
	status := "is not available."
	if b.available {
		status = "is available."
	}
	fmt.Println(b.title, status)
}

func main() {
	b := Book{"Go Programming", true}
	b.checkAvailability()
	b.borrow()
	b.checkAvailability()
	b.returnBook()
	b.checkAvailability()
}

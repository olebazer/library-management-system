package main

import "fmt"

type Author struct {
    id int
    firstName string
    lastName string
    birthday string
}

type Book struct {
    id int
    title string
    author Author
    release string
    available bool
}

func showBookInfo(book Book) {
    fmt.Println("ID:", book.id)
    fmt.Println("Title:", book.title)
    fmt.Printf("Author: %s ", book.author.firstName)
    fmt.Printf("%s\n", book.author.lastName)
    fmt.Println("Release:", book.release)
    fmt.Println("Available:", book.available)
}

var authors = []Author{
    {
        id: 1,
        firstName: "Harper",
        lastName: "Lee",
        birthday: "28.04.1926",
    },
    {
        id: 2,
        firstName: "Suzanne",
        lastName: "Collins",
        birthday: "11.08.1962",
    },
    {
        id: 3,
        firstName: "William",
        lastName: "Golding",
        birthday: "19.09.1911",
    },
}

var books = []Book{
    {
        id: 1,
        title: "To Kill a Mockingbird",
        author: authors[0],
        release: "11.07.1960",
        available: true,
    },
    {
        id: 2,
        title: "The Hunger Games",
        author: authors[1],
        release: "14.09.2008",
        available: true,
    },
    {
        id: 3,
        title: "Lord of the Flies",
        author: authors[2],
        release: "17.09.1954",
        available: false,
    },
}

func main() {
    for index, book := range books {
        fmt.Printf("Book No. %d:\n", index + 1)
        showBookInfo(book)
        fmt.Println()
    }
}

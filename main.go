package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Author struct {
	id        int
	firstName string
	lastName  string
	birthday  string
}

type Book struct {
	id        int
	title     string
	author    Author
	release   string
	available bool
}

var authors = []Author{
	{
		id:        1,
		firstName: "Harper",
		lastName:  "Lee",
		birthday:  "28.04.1926",
	},
	{
		id:        2,
		firstName: "Suzanne",
		lastName:  "Collins",
		birthday:  "11.08.1962",
	},
	{
		id:        3,
		firstName: "William",
		lastName:  "Golding",
		birthday:  "19.09.1911",
	},
}

var books = []Book{
	{
		id:        1,
		title:     "To Kill a Mockingbird",
		author:    authors[0],
		release:   "11.07.1960",
		available: true,
	},
	{
		id:        2,
		title:     "The Hunger Games",
		author:    authors[1],
		release:   "14.09.2008",
		available: true,
	},
	{
		id:        3,
		title:     "Lord of the Flies",
		author:    authors[2],
		release:   "17.09.1954",
		available: false,
	},
}

var scanner = bufio.NewScanner(os.Stdin)

func showBookInfo(book Book) {
	fmt.Println("ID:", book.id)
	fmt.Println("Title:", book.title)
	fmt.Printf("Author: %s ", book.author.firstName)
	fmt.Printf("%s\n", book.author.lastName)
	fmt.Println("Release:", book.release)
	fmt.Println("Available:", book.available)
}

func listBooks() {
	for index, book := range books {
		fmt.Printf("Book No. %d:\n", index+1)
		showBookInfo(book)
		fmt.Println()
	}
}

func createBook() {
	id := len(books) + 1
	fmt.Print("Enter book title: ")
	scanner.Scan()
	title := scanner.Text()
	fmt.Print("Enter author id: ")
	scanner.Scan()
	authorId, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Book creation failed")
		return
	}
	fmt.Print("Enter release date: ")
	scanner.Scan()
	release := scanner.Text()
	available := true
	books = append(books, Book{id, title, authors[authorId-1], release, available})
}

func showMenu() {
	fmt.Println("1 -> show menu")
	fmt.Println("2 -> list books")
	fmt.Println("3 -> create book")
	fmt.Println("4 -> list authors")
	fmt.Println("5 -> create author")
	fmt.Println("0 -> quit")
}

func shutdown() {
	fmt.Println("Good Bye")
}

func main() {
	fmt.Println("### LIBRARY MANAGEMENT SYSTEM ###")
	showMenu()

	commands := map[string]func(){
		"1": showMenu,
		"2": listBooks,
		"3": createBook,
		"4": listBooks,
		"5": createBook,
		"0": shutdown,
	}

	var input string
	for input != "0" {
		fmt.Print("(lms)> ")
		scanner.Scan()
		input = scanner.Text()
		if command, ok := commands[input]; ok {
			command()
		} else {
			fmt.Println("Invalid command")
		}
	}
}

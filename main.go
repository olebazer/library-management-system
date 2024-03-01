package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Author struct {
	Id        int
	FirstName string
	LastName  string
	Birthday  string
}

type Book struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	AuthorId  int    `json:"authorId"`
	Release   string `json:"release"`
	Available bool   `json:"available"`
}

var authors = []Author{}
var books = []Book{}
var scanner = bufio.NewScanner(os.Stdin)

func readJSONData() {
	file, err := os.ReadFile("./books.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	err = json.Unmarshal(file, &books)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	file, err = os.ReadFile("./authors.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	err = json.Unmarshal(file, &authors)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
}

func showBookInfo(book Book) {
	fmt.Println("ID:", book.Id)
	fmt.Println("Title:", book.Title)
	author := authors[book.AuthorId]
	fmt.Printf("Author: %s ", author.FirstName)
	fmt.Printf("%s\n", author.LastName)
	fmt.Println("Release:", book.Release)
	fmt.Printf("Available: %t\n\n", book.Available)
}

func showAuthorInfo(author Author) {
	fmt.Println("ID:", author.Id)
	fmt.Printf("Name: %s %s\n", author.FirstName, author.LastName)
	fmt.Printf("Birthday: %s\n\n", author.Birthday)
}

func listBooks() {
	fmt.Println("+++ LIST OF BOOKS +++")
	for index, book := range books {
		fmt.Printf("Book No. %d:\n", index+1)
		showBookInfo(book)
	}
}

func listAuthors() {
	fmt.Println("+++ LIST OF AUTHORS +++")
	for index, author := range authors {
		fmt.Printf("Author No. %d:\n", index+1)
		showAuthorInfo(author)
	}
}

func createBook() {
	fmt.Println("+++ CREATE NEW BOOK +++")
	id := len(books) + 1
	fmt.Print("Enter title: ")
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
	books = append(books, Book{id, title, authorId, release, available})
	fmt.Println()
}

func createAuthor() {
	fmt.Println("+++ CREATE NEW AUTHOR +++")
	id := len(authors) + 1
	fmt.Print("Enter first name: ")
	scanner.Scan()
	firstName := scanner.Text()
	fmt.Print("Enter last name: ")
	scanner.Scan()
	lastName := scanner.Text()
	fmt.Print("Enter birthday: ")
	scanner.Scan()
	birthday := scanner.Text()
	authors = append(authors, Author{id, firstName, lastName, birthday})
	fmt.Println()
}

func showMenu() {
	fmt.Println("+++ MENU +++")
	fmt.Println("1 -> show menu")
	fmt.Println("2 -> list books")
	fmt.Println("3 -> create book")
	fmt.Println("4 -> list authors")
	fmt.Println("5 -> create author")
	fmt.Printf("0 -> quit\n\n")
}

func shutdown() {
	fmt.Println("Good Bye")
}

func main() {
	readJSONData()
	fmt.Printf("### LIBRARY MANAGEMENT SYSTEM ###\n\n")
	showMenu()

	commands := map[string]func(){
		"1": showMenu,
		"2": listBooks,
		"3": createBook,
		"4": listAuthors,
		"5": createAuthor,
		"0": shutdown,
	}

	var input string
	for input != "0" {
		fmt.Print("(lms)> ")
		scanner.Scan()
		input = scanner.Text()
		fmt.Println()
		if command, ok := commands[input]; ok {
			command()
		} else {
			fmt.Printf("Invalid command\n\n")
		}
	}
}

// TODO: read books and authors from json file
// TODO: split project into multiple files
// TODO: intorduce customers
// TODO: try to use pointers

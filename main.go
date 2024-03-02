package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Author struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
}

type Book struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	AuthorId  int    `json:"authorId"`
	Release   string `json:"release"`
	Available bool   `json:"available"`
}

type Customer struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

var authors = []Author{}
var books = []Book{}
var customers = []Customer{}
var scanner = bufio.NewScanner(os.Stdin)

func readJSONData() {
	// reading books.json
	file, err := os.ReadFile("data/books.json")
	if err != nil {
		fmt.Println("Error reading book data:", err)
		return
	}
	// parsing books.json
	err = json.Unmarshal(file, &books)
	if err != nil {
		fmt.Println("Error parsing book data:", err)
		return
	}

	// reading authors.json
	file, err = os.ReadFile("data/authors.json")
	if err != nil {
		fmt.Println("Error reading author data:", err)
		return
	}
	// parsing authors.json
	err = json.Unmarshal(file, &authors)
	if err != nil {
		fmt.Println("Error parsing author data:", err)
		return
	}

	// reading customers.json
	file, err = os.ReadFile("data/customers.json")
	if err != nil {
		fmt.Println("Error reading customer data:", err)
		return
	}
	// parsing customers.json
	err = json.Unmarshal(file, &customers)
	if err != nil {
		fmt.Println("Error parsing customer data:", err)
	}
}

func showBookInfo(book *Book) {
	fmt.Println("ID:", book.Id)
	fmt.Println("Title:", book.Title)
	author := authors[book.AuthorId]
	fmt.Printf("Author: %s ", author.FirstName)
	fmt.Printf("%s\n", author.LastName)
	fmt.Println("Release:", book.Release)
	fmt.Printf("Available: %t\n\n", book.Available)
}

func showAuthorInfo(author *Author) {
	fmt.Println("ID:", author.Id)
	fmt.Printf("Name: %s %s\n", author.FirstName, author.LastName)
	fmt.Printf("Birthday: %s\n\n", author.Birthday)
}

func showCustomerInfo(customer *Customer) {
	fmt.Println("ID:", customer.Id)
	fmt.Println("Email:", customer.Email)
	fmt.Println("Username:", customer.Username)
}

func listBooks() {
	fmt.Println("+++ LIST OF BOOKS +++")
	for index, book := range books {
		fmt.Printf("Book No. %d:\n", index+1)
		showBookInfo(&book)
	}
}

func listAuthors() {
	fmt.Println("+++ LIST OF AUTHORS +++")
	for index, author := range authors {
		fmt.Printf("Author No. %d:\n", index+1)
		showAuthorInfo(&author)
	}
}

func listCustomers() {
	fmt.Println("+++ LIST OF CUSTOMERS +++")
	for index, customer := range customers {
		fmt.Printf("Customer No. %d:\n", index+1)
		showCustomerInfo(&customer)
	}
}

func createBook() {
	// create book from user input
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
	// add book to books array
	books = append(books, Book{id, title, authorId, release, available})
	fmt.Println()

	// write books to json file
	file, err := os.Create("data/books.json")
	if err != nil {
		fmt.Println("Error creating books file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(books)
	if err != nil {
		fmt.Println("Error encoding books JSON:", err)
		return
	}
}

func createAuthor() {
	// create author from user input
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
	// add author to authors array
	authors = append(authors, Author{id, firstName, lastName, birthday})
	fmt.Println()

	// write authors to json file
	file, err := os.Create("data/authors.json")
	if err != nil {
		fmt.Println("Error creating authors file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(authors)
	if err != nil {
		fmt.Println("Error encoding authors JSON:", err)
		return
	}
}

func createCustomer() {
	// create customer from user input
	fmt.Println("+++ CREATE NEW CUSTOMER +++")
	id := len(customers) + 1
	fmt.Print("Enter email adress: ")
	scanner.Scan()
	email := scanner.Text()
	fmt.Print("Enter username: ")
	scanner.Scan()
	username := scanner.Text()
	// add customer to customers array
	customers = append(customers, Customer{id, email, username})

	// write customers to json file
	file, err := os.Create("data/customers.json")
	if err != nil {
		fmt.Println("Error creating customers file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(customers)
	if err != nil {
		fmt.Println("Error encoding customers JSON:", err)
		return
	}
}

func showMenu() {
	fmt.Println("+++ MENU +++")
	fmt.Println("1 -> show menu")
	fmt.Println("2 -> list books")
	fmt.Println("3 -> create book")
	fmt.Println("4 -> list authors")
	fmt.Println("5 -> create author")
	fmt.Println("6 -> list customers")
	fmt.Println("7 -> create customer")
	fmt.Printf("0 -> quit\n\n")
}

func shutdown() {
	fmt.Println("Good Bye")
}

func main() {
	// fill authors, books and customers with json data
	readJSONData()
	fmt.Printf("### LIBRARY MANAGEMENT SYSTEM ###\n\n")
	showMenu()

	commands := map[string]func(){
		"1": showMenu,
		"2": listBooks,
		"3": createBook,
		"4": listAuthors,
		"5": createAuthor,
		"6": listCustomers,
		"7": createCustomer,
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

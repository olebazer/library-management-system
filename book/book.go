package book

import (
    "bufio"
    "encoding/json"
    "fmt"
    "library-management-system/author"
    "os"
    "strconv"
)

type Book struct {
    Id int `json:"id"`
    Title string `json:"title"`
    AuthorId int `json:"authorId"`
    Release string `json:"release"`
    Available bool `json:"available"`
}

var Books = []Book{}

func ReadBooksData() {
    // reading books.json
    file, err := os.ReadFile("data/books.json")
    if err != nil {
        fmt.Println("Error reading books data:", err)
        return
    }
    // parsing books.json
    err = json.Unmarshal(file, &Books)
    if err != nil {
        fmt.Println("Error parsing book data:", err)
    }
}

func showBookInfo(book *Book) {
    fmt.Println("ID:", book.Id)
    fmt.Println("Title:", book.Title)
    author := book.getAuthor()
    fmt.Printf("Author: %s ", author.FirstName)
    fmt.Printf("%s\n", author.LastName)
    fmt.Println("Release:", book.Release)
    fmt.Printf("Available: %t\n\n", book.Available)
}

func (book *Book) getAuthor() author.Author {
    for _, author := range author.Authors {
        if author.Id == book.AuthorId {
            return author
        }
    }
    return author.Author{}
}

func ListBooks() {
    fmt.Println("+++ LIST OF BOOKS +++")
    for index, book := range Books {
        fmt.Printf("Book No. %d:\n", index+1)
        showBookInfo(&book)
    }
}

func CreateBook() {
    // create book from user input
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("+++ CREATE NEW BOOK +++")
    id := len(Books) + 1
    fmt.Print("Enter title: ")
    scanner.Scan()
    title := scanner.Text()
    fmt.Print("Enter author id: ")
    scanner.Scan()
    authorId, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Printf("Book creation failed\n\n")
        return
    }
    fmt.Print("Enter release date: ")
    scanner.Scan()
    release := scanner.Text()
    available := true
    // add book to Books array
    Books = append(Books, Book{id, title, authorId, release, available})
    fmt.Println()

    // write Books to json file
    file, err := os.Create("data/Books.json")
    if err != nil {
        fmt.Println("Error creating Books file:", err)
        return
    }
    defer file.Close()
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "    ")
    err = encoder.Encode(Books)
    if err != nil {
        fmt.Println("Error encoding Books JSON:", err)
    }
}

package main

import (
    "bufio"
    "fmt"
    "library-management-system/author"
    "library-management-system/book"
    "library-management-system/customer"
    "os"
)

func readJSONData() {
    author.ReadAuthorsData()
    book.ReadBooksData()
    customer.ReadCustomersData()
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
    // fill Authors, Books and Customers arrays with json data
    readJSONData()
    fmt.Printf("### LIBRARY MANAGEMENT SYSTEM ###\n\n")
    showMenu()

    commands := map[string]func(){
        "1": showMenu,
        "2": book.ListBooks,
        "3": book.CreateBook,
        "4": author.ListAuthors,
        "5": author.CreateAuthor,
        "6": customer.ListCustomers,
        "7": customer.CreateCustomer,
        "0": shutdown,
    }

    scanner := bufio.NewScanner(os.Stdin)
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

package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "library-management-system/author"
    "library-management-system/book"
    "library-management-system/customer"
    "os"
)

func readJSONData() {
    // reading books.json
    file, err := os.ReadFile("data/books.json")
    if err != nil {
        fmt.Println("Error reading book data:", err)
        return
    }
    // parsing books.json
    err = json.Unmarshal(file, &book.Books)
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
    err = json.Unmarshal(file, &author.Authors)
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
    err = json.Unmarshal(file, &customer.Customers)
    if err != nil {
        fmt.Println("Error parsing customer data:", err)
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

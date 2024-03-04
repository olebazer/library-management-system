package customer

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
)

type Customer struct {
    Id int `json:"id"`
    Email string `json:"email"`
    Username string `json:"username"`
}

var Customers = []Customer{}

func showCustomerInfo(customer *Customer) {
    fmt.Println("ID:", customer.Id)
    fmt.Println("Email:", customer.Email)
    fmt.Printf("Username: %s\n\n", customer.Username)
}

func ListCustomers() {
    fmt.Println("+++ LIST OF CUSTOMERS +++")
    for index, customer := range Customers {
        fmt.Printf("Customer No. %d:\n", index+1)
        showCustomerInfo(&customer)
    }
}

func CreateCustomer() {
    // create customer from user input
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("+++ CREATE NEW CUSTOMER +++")
    id := len(Customers) + 1
    fmt.Print("Enter email adress: ")
    scanner.Scan()
    email := scanner.Text()
    fmt.Print("Enter username: ")
    scanner.Scan()
    username := scanner.Text()
    // add customer to Customers array
    Customers = append(Customers, Customer{id, email, username})
    fmt.Println()

    // write Customers to json file
    file, err := os.Create("data/customers.json")
    if err != nil {
        fmt.Println("Error creating Customers file:", err)
        return
    }
    defer file.Close()
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "    ")
    err = encoder.Encode(Customers)
    if err != nil {
        fmt.Println("Error encoding Customers JSON:", err)
        return
    }
}

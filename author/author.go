package author

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
)

type Author struct {
    Id int `json:"id"`
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Birthday string `json:"birthday"`
}

var Authors = []Author{}

func ReadAuthorsData() {
    // reading authors.json
    file, err := os.ReadFile("data/authors.json")
    if err != nil {
        fmt.Println("Error reading authors data:", err)
        return
    }
    // parsing authors.json
    err = json.Unmarshal(file, &Authors)
    if err != nil {
        fmt.Println("Error parsing author data:", err)
    }
}

func showAuthorInfo(author *Author) {
    fmt.Println("ID:", author.Id)
    fmt.Printf("Name: %s %s\n", author.FirstName, author.LastName)
    fmt.Printf("Birthday: %s\n\n", author.Birthday)
}

func ListAuthors() {
    fmt.Println("+++ LIST OF AUTHORS +++")
    for index, author := range Authors {
        fmt.Printf("Author No. %d:\n", index+1)
        showAuthorInfo(&author)
    }
}

func CreateAuthor() {
    // create author from user input
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("+++ CREATE NEW AUTHOR +++")
    id := len(Authors) + 1
    fmt.Print("Enter first name: ")
    scanner.Scan()
    firstName := scanner.Text()
    fmt.Print("Enter last name: ")
    scanner.Scan()
    lastName := scanner.Text()
    fmt.Print("Enter birthday: ")
    scanner.Scan()
    birthday := scanner.Text()
    // add author to Authors array
    Authors = append(Authors, Author{id, firstName, lastName, birthday})
    fmt.Println()

    // write Authors to json file
    file, err := os.Create("data/Authors.json")
    if err != nil {
        fmt.Println("Error creating Authors file:", err)
        return
    }
    defer file.Close()
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "    ")
    err = encoder.Encode(Authors)
    if err != nil {
        fmt.Println("Error encoding Authors JSON:", err)
        return
    }
}

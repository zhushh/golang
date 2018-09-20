package main

import (
    "fmt"
)

// 结构体定义
// type struct_type_name struct {
//      member definition;
//      member definition;
//      ...    
//      member definition;
// }
// variable_name := struct_type_name {val1, val2, val3 ...}
// variable_name := struct_type_name {key1:val1, key2: val2 ...}

type Books struct {
    title string
    author string
    subject string
    book_id int
}

func PrintBook(pBook *Books) {
    fmt.Printf("Book title: %s\n", pBook.title)
    fmt.Printf("Book author: %s\n", pBook.author)
    fmt.Printf("Book subject: %s\n", pBook.subject)
    fmt.Printf("Book id: %d\n", pBook.book_id)
}

func main() {
    go_book := Books{"Go language", "nickname", "Go basic grammer", 23312}

    fmt.Println(go_book)

    fmt.Println(Books{title: "C language", author: "nickname", subject: "Basic", book_id: 4232})

    fmt.Println(Books{title: "C++", author: "Mark"})

    var book1 Books
    book1.title = "Book1"
    book1.author = "Ding"
    book1.subject = "Ohter language"
    book1.book_id = 81273
    PrintBook(&book1)

    var pBook = &book1
    PrintBook(pBook)
}

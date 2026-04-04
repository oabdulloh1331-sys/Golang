package main

import "fmt"

var password string 

func main() {
    fmt.Println("Enter password:")
    fmt.Scanf("%s", &password)
    if password == "12346" {
        fmt.Println("Password correct.")
    } else {
        fmt.Println("Password incorrect.")
    }
}

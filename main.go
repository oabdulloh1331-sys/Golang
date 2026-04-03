package main

import "fmt"
var name string 
var age int8

func main() {
    fmt.Println("What is your name?")
    fmt.Scanf("%s", &name)
    fmt.Println("What is your age?")
    fmt.Scanf("%d", &age)
    fmt.Println("Hello, " + name)
    fmt.Println("You are " + fmt.Sprint(age) + " years old.")
}


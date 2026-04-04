package main

import "fmt"

func main() {
    var sum int
    var age int

    fmt.Println("Enter the total sum:")
    fmt.Scanf("%d", &sum)

    fmt.Println("Enter your age:")
    fmt.Scanf("%d", &age)

    if sum > 5000 || age > 60 {
        fmt.Println("You have 20% discount!")
    } else if sum > 2000 && age < 18 {
        fmt.Println("You have 10% discount!")
    } else {
        fmt.Println("No discount today.")
    }
}
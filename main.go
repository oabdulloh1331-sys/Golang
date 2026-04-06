package main

import "fmt"

func main() {
   password := ""
for password != "1234" {
    fmt.Println("Vvedite parol:")
    fmt.Scan(&password)
}
fmt.Println("Dostup otkrit!")



}
package main

import "fmt"

func main() {
    var name string = "Rashida" // full declaration
    var age = 25               // type inference
    city := "Accra"            // short declaration
    const PI = 3.14159         // constant

    fmt.Println(name)  // Rashida
    fmt.Println(age)   // 25
    fmt.Println(city)  // Accra
    fmt.Println(PI)    // 3.14159
}

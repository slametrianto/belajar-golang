package main

import "fmt"

func main(){
	fmt.Println("a...", 1 - 1 )
	fmt.Println("hasil =", 1 + 1)
	fmt.Println("total =", 10 * 2)
	fmt.Println("total =", 10 / 3)
	fmt.Println("total =", 10.0 / 3.0)
	fmt.Println("total =", 10 % 2)

    // booleans
	fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)

}
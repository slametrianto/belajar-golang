package main

import "fmt"

func main() {

	// if i % 3 = foo
	// if i % 5 = bar
	// if i % 3 && i % 5 = foobar
	
	for i := 1; i <= 30; i++ {

		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("foobar")

		} else if i % 3 == 0 {
			fmt.Println("foo")

		} else if i % 5 == 0 {
			fmt.Println("bar")
			
		} else {
			fmt.Println(i)
		}
	}
}
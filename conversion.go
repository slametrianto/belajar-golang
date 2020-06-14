package main

import "fmt"

func main(){
	var x int32 = 10
	var y int64 = int64(x)

	println(y)

	var z float64 = float64(y)
	fmt.Println(z)

	var nilai string = "100"
	
	nilaiInt, _ := strconv.Atoi(nilai)
	trconv.Atoi("-42")
	fmt.Println(nilaiInt)

	nilaiString := strconv.Itoa(-100)
	fmt.Println(nilaiString)

	// var nilai string = "100"
	// nilai i, err := strconv.Atoi("-42")
    // fmt.Println(nilai)

}
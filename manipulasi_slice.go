package main

import "fmt"

func main() {

	// slice1 := []string {
	// 	"slamet",
	// 	"rianto",
	// }
	
	// slice1 := make([]string, )
	// fmt.Println(slice1)


	// slice1 := make([]string, 3)

	// slice1[0] ="reni"
	// slice1[1] ="nur"
	// slice1[2] ="mamet"
	 
	// slice2 := slice1[:]

	// slice1[0] = "budi"

	// fmt.Println(slice1)
	// fmt.Println(slice2)



	slice1 := make([]string, 3)

	slice1[0] ="reni"
	slice1[1] ="nur"
	slice1[2] ="mamet"

	slice2 := append(slice1, "ervyna", "putri")

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[0] = "joko"
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := make([]string, 3)

	copy(slice3, slice1)

	fmt.Println(slice1)
	fmt.Println(slice3)

	slice1[0] = "budi"

	fmt.Println(slice1)
	fmt.Println(slice3)


}
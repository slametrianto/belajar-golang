package main

import "fmt"

func main(){

	var names [5]string
	names[0] = "reni"
	names[1] = "nur"
	names[2] = "mamet"
	names[3] = "rianto"
	names[4] = "slamet"

	// for i := 0; i < 3; i++ {
	
	//mengetahui panjang array
	for i := 0; i < len(names); i++ {
	 fmt.Println(names[i])

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])
     
}

//cara kedua

// for index, value := range names {
// 	fmt.Println("index", index, "=", value)



    //mengambil nilai value
	// for _, value := range names {
	// fmt.Println(value)

	//mengambil nilai index
	for index, _ := range names {
	fmt.Println(index)


    	}
	}





package main

import "fmt"

func main(){
	
	mahasiswa := make(map[string]string)

	mahasiswa["1001"] ="reni"
	mahasiswa["1002"] ="nur"
	mahasiswa["1003"] ="mamet"

	fmt.Println(mahasiswa)

	fmt.Println(mahasiswa["1001"])
	fmt.Println(mahasiswa["1002"])
	fmt.Println(mahasiswa["1003"])

	for nim, name := range mahasiswa {
		fmt.Println("nim", nim, "bernama", name)
		
	}

}
package main

import "fmt"

func main(){

	// mahasiswa := map[string]string {
		
	// 	"1001" : "eko",
	// 	"1002" : "slamet",
	// 	"1003" : "reni",

	// }

	// fmt.Println(mahasiswa)

	mahasiswa := map[string]map[string]string {

      "1001" : map[string]string {
		  "name" : "slamet",
		  "address" : "indonesia",
		  "gender" : "male",
	  },

	  "1002" : map[string]string {
		"name" : "slamet",
		"address" : "indonesia",
		"gender" : "male",

	},

	"1003" : map[string]string {
		"name" : "slamet",
		"address" : "indonesia",
		"gender" : "male",

},

	}

  	// fmt.Println(mahasiswa)
	// fmt.Println(mahasiswa["1001"]["name"])

	delete(mahasiswa,  "1003")
	fmt.Println(mahasiswa)


}
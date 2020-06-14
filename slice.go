package main

import "fmt"

func main(){

	names := [5]string {
	"reni",
	"nur",
	"slamet",
	"rianto",
	"mamet",
}

fmt.Println(names)

names[1] = "nugraha"
slice1 := names[1:4]
fmt.Println(slice1)

slice2 := names[:3]
fmt.Println(slice2)

slice3 :=names[1:]
fmt.Println(slice3)

slice4 := names[0:]
fmt.Println(slice4)

}
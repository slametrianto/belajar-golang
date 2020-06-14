package main

import "fmt"

func main(){
for i := 1; i <= 10; i++ {

	switch i {
	case 1:
		fmt.Println("satu") 
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("tiga") 
	case 4:
		fmt.Println("empat")  	
	case 5:
		fmt.Println("lima")
	case 6:
		fmt.Println("enam")
	case 7:
		fmt.Println("tujuh")
	case 8:
		fmt.Println("delapan")	
	default:
		fmt.Println("gak tau")	

	}

  }

  sistemOperasi := runtime.GOOS

   switch sistemOperasi{
   case "darwin":
	fmt.Println("mac")

   case "linux":
	fmt.Println("linux")

   default:
	fmt.Println("gak tau")

   }
	
}
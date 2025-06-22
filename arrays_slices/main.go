package main

import "fmt"

func main(){

	/* var names [3]string

	names[0] = "Mehmetcan"
	names[1] = "Ece"
	names[2] = "Ali" */

	var names =[3] string{"Mehmetcan", "Ece", "Ali","Jack"}	

	names[0] = "Serhat"
	names[1] = " Mehmetcan"

	names[3] = "Zeynep" // This will cause a compile-time error because the array size is fixed at 3.


	fmt.Println(names)
}
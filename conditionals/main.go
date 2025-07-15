package main

import "fmt"


func main() {

	/* var age = 17

	if age >= 18 {
		fmt.Println("you can vote.")
	} else { //else burada yazılmalı aşağı satırda yazarsan syntax error alıyorsun
		fmt.Println("you cannot vote.")
	} */

	a := 1
	b := 6
	c := 9


	if a>=b && a>=c {
		fmt.Println("a is the biggest one")
	} else if b >= a && b >= c {
		fmt.Println("b is the biggest one")
	} else {
		fmt.Println("c is the biggest one")
	}




	

}


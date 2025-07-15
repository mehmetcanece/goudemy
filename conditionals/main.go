package main

import "fmt"


func main() {

	/* var age = 17

	if age >= 18 {
		fmt.Println("you can vote.")
	} else { //else burada yazılmalı aşağı satırda yazarsan syntax error alıyorsun
		fmt.Println("you cannot vote.")
	} */

	a := 9
	b := 9
	c := 3


	if a>=b && a>=c {
		fmt.Println("a is the biggest one")
	} else if b >= a && b >= c {
		fmt.Println("b is the biggest one")
	} else if c>=a && c>=b {
		fmt.Println("c is the biggest one")
	}




	

}


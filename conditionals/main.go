package main

import "fmt"


func main() {

	var age = 17

	if age >= 18 {
		fmt.Println("you can vote.")
	} else { //else burada yazılmalı aşağı satırda yazarsan syntax error alıyorsun
		fmt.Println("you cannot vote.")
	}

}


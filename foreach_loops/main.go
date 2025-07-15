package main

import "fmt"

func main(){

/* 		var numbers = []int {1,2,3,4}
 */
		/* for index := 0 ; index < len(numbers) ; index++{
				fmt.Println(numbers[index])
		} */

		/* for _, value := range numbers { //index yazıp kafa karıştıracağına kullanmayacağın değer varsa _ ile atayabilirsin
			fmt.Println( value)
		} */

	/* 	var language = "Golang"


		for _, character := range language{
			fmt.Printf("Character %c\n", character)
		} */

		var names = map[string]int {
			"Mehmet":1,
			"Can":2,
			"Ece":3,
		}

		for key, value := range names{ // key ve value isimleri değiştirilebilir, önemli olan sıralama
			fmt.Println(key, value)
		}







	}
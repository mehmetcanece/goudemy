package main

import "fmt"

func main(){
	/* total := add(10,20)
	fmt.Println(total) */

	total, difference := calculation(10,20)
	fmt.Println(total,difference)
}

func add(x int, y int) int{
	//fmt.Println(x+y)
	return x + y
}


func calculation(x int, y int) (int, int){
	return x + y, x - y
} //iki farklı değer de dönebilir 

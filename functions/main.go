package main

import "fmt"

func main(){
	/* total := add(10,20)
	fmt.Println(total) */

	total, difference, multiple, divide := calculator(10,20)
	fmt.Println(total,difference,multiple,divide)
}

func add(x int, y int) int{
	//fmt.Println(x+y)
	return x + y
}


func calculator(x int, y int) (int, int, int, float64){
	return x + y, x - y , x * y, float64(x) / float64(y) // float64 dönüşümünü yapmazsak int olarak döner ve tam sayı olarak bölme işlemi yapar
} //iki farklı değer de dönebilir 

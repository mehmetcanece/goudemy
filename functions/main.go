package main

import "fmt"

func main(){
	/* total := add(10,20)
	fmt.Println(total) */

/* 	total, difference, multiple, divide := calculator(10,20)
	fmt.Println(total,difference,multiple,divide)
 */

	/* var numbers = []int{10,20,2,3,5}

	sum := slice_adder(numbers)
	fmt.Println(sum) */


/* 	fmt.Println(sum(3,4,5))
 */


 fmt.Println(sum(1,2,3,4,5,2,2,4,43))
 fmt.Println(sum(1,2,3,4,5,6)) //fonksiyonun içine kaç tane değer girersen o kadar alır ve toplar
}


func sum (numbers ...int)int { //bu numbers array gibi olucak kaç değer gelirse artık içinde tutucak
	sum := 0
	for _, value := range numbers{ // ne kadar değişken göndereceğin belli değilse bu şekilde kullan.
		sum += value
	}
	return sum
}

/* func sum (x int, y int , z int) int {

	return x + y + z

} */






/* func slice_adder(number []int)int{
	result := 0
	for _ , value := range number{
		result += value
	}
	return result
}

func add(x int, y int) int{
	//fmt.Println(x+y)
	return x + y
}


func calculator(x int, y int) (int, int, int, float64){
	return x + y, x - y , x * y, float64(x) / float64(y) // float64 dönüşümünü yapmazsak int olarak döner ve tam sayı olarak bölme işlemi yapar
} //iki farklı değer de dönebilir 
 */
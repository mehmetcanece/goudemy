package main

import "fmt"


func add12(x int){ // pass by value
	x += 12
}

func add12pointer(x *int){ // pass / call by reference
	*x += 12
}

func changeValue(numbers []int){ // slicelar referans tiplerdir, bu yüzden referans olarak geçer	
	numbers[0] = 1000
}

func main(){


	var numbers = []int{1,2,3}

	fmt.Println(numbers)
	changeValue(numbers)
	fmt.Println(numbers)
	/* var a = 10
	fmt.Println(a)
	add12(a)
	add12pointer(&a) //a nın adres bilgisini göndermemiz gerekiyo
	fmt.Println(a) */




	/* var a  = 10
	var b int
	var p *int
	
	b=a
	p= &a // a'nın adresini p'ye atadık

	*p=20

	fmt.Println(a,b) */



	/* var a int

	var p *int

	//var j *int

	a = 10

	p = &a // a'nın adresini p'ye atadık

	*p = 20 // p'nin işaret ettiği değeri 20 yaptık

	fmt.Println(a) */

	/* j = p // a'nın adresini j'ye de atadık

	fmt.Println(a)

	fmt.Println(p)

	fmt.Println(j)

	fmt.Println(*p)  */// p'nin işaret ettiği değeri yazdırdık. 
	//& işareti adresi veriyor, dereferencing
	
	// * işareti ise o adresin içindeki değeri veriyor.


}
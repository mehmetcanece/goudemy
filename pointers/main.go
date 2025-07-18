package main

import "fmt"

func main(){


	var a int

	var p *int

	var j *int

	a = 10

	p = &a // a'nın adresini p'ye atadık

	j = p // a'nın adresini j'ye de atadık

	fmt.Println(a)

	fmt.Println(p)

	fmt.Println(j)

	fmt.Println(*p) // p'nin işaret ettiği değeri yazdırdık. 
	//& işareti adresi veriyor, 
	
	// * işareti ise o adresin içindeki değeri veriyor.


}
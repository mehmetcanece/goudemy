package main

import "fmt"

func main(){


	var a int

	var p *int

	a = 10

	p = &a // a'nın adresini p'ye atadık

	fmt.Println(a)

	fmt.Println(p)

	fmt.Println(*p) // p'nin işaret ettiği değeri yazdırdık. 
	//& işareti adresi veriyor, 
	
	// * işareti ise o adresin içindeki değeri veriyor.


}
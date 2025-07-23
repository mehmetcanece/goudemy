package main

import "fmt"

func main(){


	var a  = 10
	var b int
	var p *int
	
	b=a
	p= &a // a'nın adresini p'ye atadık

	*p=20

	fmt.Println(a,b)



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
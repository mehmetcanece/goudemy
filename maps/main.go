package main

import "fmt"

func main(){

	/* var names map[string]int 

	names = make(map[string]int, 0)

	names["Mehmet"]=1
	names["Can"]=2
	names["Ece"]=3

	fmt.Println(names)
	fmt.Println(names["Mehmet"])
	fmt.Println(names["Ali"]) // 0 döner çünkü yok.  */


	names := map[string]int{
		"Mehmet":1,
		"Can":2,
		"Ece":3,	
	}


	delete(names,"Ece")

	fmt.Println(names)



}
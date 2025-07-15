package main

import "fmt"

func main(){

/* 	index := 1
	 

	for index <= 10 {
		fmt.Println(index)
		index++
	}
 */
/* total := 0
 for index := 1 ; index <= 3 ; index++ {
	fmt.Println(index)
	total += index
	fmt.Println("Total:", total)
 } */

/* index := 0

var names = [3]string{"Mehmet","Can","Ece"}

for index<len(names) {
	fmt.Println(names[index])
	index++
} */



/* for index := 0; index <=10; index++ {
	if index == 3 {
		break
	}
	fmt.Println(index)
} */


for index := 0; index <=10 ; index++{
	if index ==2 || index ==5 {
		continue
	}
	fmt.Println(index)
}






}
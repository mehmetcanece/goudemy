package main

import "fmt"



func main(){

	/* 
	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Golang Dersi"
	quantity = 10
	discount = 0.37
	isInStock = true

	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity))
	fmt.Println(discount, reflect.TypeOf(discount))
	fmt.Println(isInStock, reflect.TypeOf(isInStock)) */

/* 	var prodcutName = "Golang Dersi"


	fmt.Println(prodcutName, reflect.TypeOf(prodcutName))
 */

	/* productName := "Golang Dersi"

	fmt.Println(productName) */

/* 	var quantity int64 = 10

	fmt.Println(quantity, reflect.TypeOf(quantity)) */


	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Golang Dersi"
	quantity = 10
	discount = 0.37
	isInStock = true

	/* fmt.Println("Product name:", productName, "quantity:",quantity, "discount:", discount, "is in stock:", isInStock) //yanyana yazar bu */

 fmt.Printf("Product name: %s\nQuantity: %d\nDiscount: %f\nIs in Stock: %t\n",productName,quantity,discount,isInStock) // %v dersen default şekilde tanımlar 


}
/* This is a go program created to explore variables and data types in go language */
package main

import "fmt"

var age = 19
func main() {
	var smallVal int8 = 123
	fmt.Println(smallVal)
	fmt.Printf("The variable is of type %T\n",smallVal)
	
	var largeVal int16 = 1000
	fmt.Println(largeVal)
	fmt.Printf("The variable is of type %T\n",largeVal)
	
	var veryLargeVal int32 = 10000
	fmt.Println(veryLargeVal)
	fmt.Printf("The variable is of type %T\n",veryLargeVal)
	
	var smallFloat float32 = 212.14781945
	fmt.Println(smallFloat)
	fmt.Printf("The variable is of type %T\n",smallFloat)
	
	var largeFloat float64 = 292.1287651764
	fmt.Println(largeFloat)
	fmt.Printf("The variable is of type %T\n",largeFloat)
	
	var username string = "Arnab"
	fmt.Println(username)
	fmt.Printf("The variable is of type %T\n",username)
	
	var isAdult bool = true;
	fmt.Println(isAdult)
	fmt.Printf("The variable is of type %T\n",isAdult)
	
	// default values of variables
	
	// In go variables of type string are initialized with "" by default
	var name string;
	fmt.Println("The default value of name is",name)
	
	// In go variables of type int are initialized with the value zero by default
	var num1 int
	fmt.Println("The default value of num1 is",num1)
	
	var num2 float32
	fmt.Println("The default value of num2 is",num2)
	
	var num3 float64
	fmt.Println("The default value of num3 is",num3)
	
	var cnum1 complex64
	fmt.Println("The default value of cnum1 is",cnum1)
	var cnum2 complex128
	fmt.Println("The default value of cnum2 is",cnum2)
	
	// Implicit type casting in Go
	var friend = "Shubham"
	fmt.Println("My friend's name is",friend)
	
	favSub := "Science"
	fmt.Println(favSub)
	fmt.Println("This is Arnab Mitra")
	fmt.Println("I am",age,"years old")
	const PI float32 = 3.14
	fmt.Println("The value of pi is",PI)
}
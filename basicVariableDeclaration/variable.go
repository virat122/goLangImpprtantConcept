package main

import "fmt"

// constant named value
const pi = 3.14159

func main() {

	// variable named value
	var radius = 6
	fmt.Println(radius)

	var 

	name = "vikrant" //compiler automatically infers the variable data type 
	name = "vishal" // its ok
	//name=34  // {erorr} not okk once type is string in above it cnt change into int or other type
	fmt.Println("my name is ", name)

	//int8, uint8, int16, uint16, int32 , uint32, int64, uint64, int, uint, uintptr
	var rollno int = 23
	fmt.Printf("my roll no is  no is %d \n ", rollno)
	var rollno1 uint = 23
	fmt.Printf("my roll no is  no is %d \n", rollno1)

	var rollno2 int16 = 23
	fmt.Printf("my roll no is  no is %d \n", rollno2)

	var rollno3 int32 = 23
	fmt.Printf("my roll no is  no is %d\n ", rollno3)

	var rollno4 int64 = 23
	fmt.Printf("my roll no is  no is %d\n ", rollno4)

	var rollno5 uint16 = 26
	fmt.Printf("my roll no is  no is %d\n ", rollno5)




}

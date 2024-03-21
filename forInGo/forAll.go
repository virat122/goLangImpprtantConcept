package main

import "fmt"

func main() {
	//  1-> similar to while loop 
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//2-> 
	for j := 0; j < 5; j++ {
        fmt.Println(j)
    }

	for i := range 3 {
        fmt.Println("range", i)
    }



// -> for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.

    for {
        fmt.Println("loop")
        break
    }
	//You can also continue to the next iteration of the loop.

	
    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }




}

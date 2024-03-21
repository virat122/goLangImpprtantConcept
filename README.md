# goLangImpprtantConcept 
<h1> vikrant</h1>

<h2> ways to declare variable </h2>
var declares 1 or more variables.

    var a = "initial"
    fmt.Println(a)
You can declare multiple variables at once.

    var b, c int = 1, 2
    fmt.Println(b, c)
 Go (compiler) will infer the type of initialized variables.

    var d = true
    fmt.Println(d)
Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.

    var e int
    fmt.Println(e)





    <h2>switch</h2>
    
Here’s a basic switch.

    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
You can use commas to separate multiple expressions in the same case statement. We use the optional default case in this example as well.

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions can be non-constants.

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
A type switch compares types instead of values. You can use this to discover the type of an interface value. In this example, the variable t will have the type corresponding to its clause.

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")



<h2>Array in go  </h2>
In Go, an array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios.


package main
import "fmt"
func main() {
Here we create an array a that will hold exactly 5 ints. The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for ints means 0s.

    var a [5]int
    fmt.Println("emp:", a)
We can set a value at an index using the array[index] = value syntax, and get a value with array[index].

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
The builtin len returns the length of an array.

    fmt.Println("len:", len(a))
Use this syntax to declare and initialize an array in one line.

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}


<h1> road map for Basic  of go </h1>

![image](https://github.com/virat122/goLangImpprtantConcept/assets/121007372/96e83a60-65d9-41ba-a233-cae4b7bafc7a)





<h1> road map for advance  of go </h1>

![image](https://github.com/virat122/goLangImpprtantConcept/assets/121007372/78cb2855-9eee-4a3d-b894-4d2384278c40)




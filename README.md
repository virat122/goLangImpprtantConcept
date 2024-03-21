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




<h1> road map for Basic  of go </h1>

![image](https://github.com/virat122/goLangImpprtantConcept/assets/121007372/96e83a60-65d9-41ba-a233-cae4b7bafc7a)





<h1> road map for advance  of go </h1>

![image](https://github.com/virat122/goLangImpprtantConcept/assets/121007372/78cb2855-9eee-4a3d-b894-4d2384278c40)




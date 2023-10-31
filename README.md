# Go Course

# What is Go ?

1. Go is a fast, statically typed, compiled language that feels like a dynamically typed interpreted language

2. It is used to build high performance appications , cloud native apps etc

3. It is also used for server-side web development

In go we have package called `fmt` that is a part of the `go` standard library. It helps us in formatting ouput and has functions for printing output to the screen
## Go Syntax
Here is a sample Go code that prints Hello World to the screen

```Go
package main 
import "fmt"
func main() {
    fmt.Println("Hello World")  
}
```
Every program in `Go` is a part of a package. A package is defined by using the `package` keyword. Here the package is named as main.

`import fmt` imports everything from the the `fmt` package.

`func main() {}` This is a function. In go we define functions with the `func` keyword followed by the function name,opening and closing parentheses and the {}.

**Note**: The left curly brace { cannot appear at the beginning of the line. It throws an error.
The `Println()` method in `fmt` package prints a line on the screen and sends the cursor to the next line

# Declaring variables
we can declare variables in two ways
1. By using the `var` keyword
The syntax is as shown below
```Go
var variableName datatype = value
```
While declaring variables with the `var` keyword you have to specify either the datatype of the variable or the value.

**You can do like this**
```Go
var variableName datatype
```
**Or like this**
```Go
var variableName = value
```
**Note:** When the datatype of the variable is not specified then the go compiler decides the data type of the variable based on the value.

There is also a syntax where you don't need to use the `var` keyword as well

```Go
package main
import "fmt"
func main() {
userName := "Shubham"
fmt.Println(userName)
fmt.Println("The type of the variable is %T",userName)
}
```
**Note:** The `:=` operator can be used only inside a function.
```Go
package main
import "fmt"
userName := "Shubham"
func main() {
fmt.Println(userName)
fmt.Println("The type of the variable is %T",userName)
}
```
The above code results in an error !!

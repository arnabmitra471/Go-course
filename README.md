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
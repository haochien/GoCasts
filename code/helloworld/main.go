/*
package is similar to project or workspace
each go file should have this line to declair this file belongs to which package

main is used to build an executable package
if use any other kind of name, then when run build, it won't compile this package (reusable package)
*/
package main

// fmt is one of the standard libs in Go
import "fmt"

//main funtion will always exist if this is executable package
func main() {
	fmt.Println("Hi there!")
}

package main //programs will start running in package main https://golang.org/pkg/fmt/

import "fmt" // these are packages imported into this file. This one is from the import path fmt

func main() {
	//note the fmt at the start, this is the package.
	//the Println is the exported name of the package. Note the capital P.
	//This makes it accessable outside the package
	//fmt.Println can be seen as package.Type
	fmt.Println("Hello World") //note there are no semi-colons at the end of a statement
}

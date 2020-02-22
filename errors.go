package main

import "fmt"
import "github.com/andresorav/errors-go/package1"

func main() {
	err := package1.RunFunction()
	if err != nil {
		fmt.Printf("Got an error: %v", err)
	}
}

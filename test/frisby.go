package main

import (
	"fmt"

	"github.com/verdverm/frisby"
)

// Import frisby go get -u github.com/verdverm/frisby

func main() {
	fmt.Println("My first Frisby Test!")

	frisby.Create("Test Get Value from JSON").
		Get("http://localhost:2525").
		Send().
		ExpectStatus(200).
		ExpectJson("Stuff.Fruit.Apples", 25)

	frisby.Create("Test template").
		Get("http://localhost:2525/template/hola").
		Send().
		ExpectStatus(200).
		ExpectContent("Test, the parameter hola")

	frisby.Create("Test 200").
		Get("http://localhost:2525/params/1222").
		Send().
		ExpectStatus(200).
		ExpectContent("1222")

	frisby.Global.PrintReport()

}

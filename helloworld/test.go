package main

import "fmt"

func testMain() {
	errorsCounter := make(map[int]int)

	errorsCounter[200]++
	fmt.Println(errorsCounter)
	test()

}

func test() bool {
	if 1 == 1 {
		fmt.Println("teste1")
		return true
	}

	if 1 == 1 {
		fmt.Println("teste")
		return true
	}
	return false
}

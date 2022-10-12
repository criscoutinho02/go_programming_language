//Este programa imprime as variáveis do cli
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printByJoin(args []string) (r string) {
	return strings.Join(os.Args, " ")
}

func printByLoop(args []string) (r string) {

	var sep string
	for i := 1; i < len(os.Args); i++ {
		r += sep + os.Args[i]
		sep = " "
	}

	return

}

func printByRange(args []string) (r string) {

	for i, arg := range os.Args {
		r = r + strconv.Itoa(i) + " " + arg + "\n"
	}
	return

}

func main() {
	res1, res2, res3 := printByJoin(os.Args), printByLoop(os.Args), printByRange(os.Args)

	fmt.Println(res1, res2, res3)
}

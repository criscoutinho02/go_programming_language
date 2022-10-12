//Echo1 exibe seus argumentos em linha de comando.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(len(os.Args))
	fmt.Println(s)
}

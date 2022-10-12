//Este programa imprime as variáveis do cli
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println((strings.Join(os.Args, " ")))
	fmt.Println(os.Args)
}

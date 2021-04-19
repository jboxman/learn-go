package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sep := ""
	str := ""
	for idx, arg := range os.Args {
		fmt.Println(strings.Join([]string{fmt.Sprintf("%d", idx), arg}, sep))
		str = fmt.Sprintf("%d %s\n", idx, arg)
		fmt.Println(str)
	}
}

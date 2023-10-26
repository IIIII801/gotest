package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	str := scan.Text()
	str2 := strings.Split(str, " ")
	aa := str2[len(str2)-1]
	fmt.Println(len(aa))
}

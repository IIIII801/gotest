package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		str2 := ""
		for _, v := range str {
			str2 = string(v) + str2
		}
		fmt.Print(str2)
	}
}

//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//)
//
//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	for scanner.Scan() {
//		str := scanner.Text()
//		for i:=len(str)-1;i>=0 ;i--{
//			fmt.Print(string(str[i]))
//		}
//	}
//}

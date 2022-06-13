// mytest
package main

import (
	"fmt"
	"test/pack"
)

func main() {
	a := [4]int{10, 20, 30, 40}
	fmt.Println("Sum is:", pack.SumArr(a))
}

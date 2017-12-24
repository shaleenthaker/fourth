package main

import (
	"fmt"
)

func main() {
	//for k := 0; k < count; k++ {

		for j := 0; j < 100; j++ {
      fmt.Println("Multiplication table for: ", j)
			for i := 0; i <= 10; i++ {
				fmt.Printf("%d*%d=%d\n", j, i, i*j)
			}
		}
	//}
}

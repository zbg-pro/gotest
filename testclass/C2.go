package testclass

import "fmt"

var Total_sum1 int = 0

func Sum_test1(a int, b int) int {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	Total_sum += (a + b)
	fmt.Printf("Total_sum: %d\n", Total_sum)
	return a + b
}

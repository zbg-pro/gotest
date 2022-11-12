package main

import (
	"fmt"
)

func main1() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	sum = 0
	for {
		sum++
		if sum > 1000 {
			break
		}
	}
	fmt.Println(sum)

	//For-each range 循环
	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 4, 5}
	for i, x := range numbers {
		fmt.Println(i, x)
	}

	//for 循环的 range 格式可以省略 key 和 value，如下实例：
	map1 := make(map[int]float64)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	map1[5] = 5.0
	map1[5] = 6.0
	for key, value := range map1 {
		fmt.Printf("key=%d value=%f\n", key, value)
	}

	for _, value := range map1 {
		fmt.Printf("value=%f\n", value)
	}
	for key, _ := range map1 {
		fmt.Printf("key=%d\n", key)
	}

	// 不使用标记
	fmt.Println("---- continue ---- ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue
		}
	}

	// 使用标记
	fmt.Println("---- continue label ----")
aaa:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue aaa
		}
	}

	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}

	var C int
	for C < 100 {
		C++
		//isSuShu := true
		for ccc := 2; ccc < C; ccc++ {
			if C%ccc == 0 {
				goto BREAK
				//isSuShu = false
				//break
			}
		}
		//if isSuShu {
		fmt.Println(C, "是素数")
		//}
	BREAK:
	}

}

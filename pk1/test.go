package main

import (
	"fmt"
	"time"
)
import "unsafe"

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	const LENGTH int = 5
	const WIDTH int = 10
	var area int
	//const a, b, c = 1, false, ""

	area = LENGTH * WIDTH

	fmt.Println("aaa")

	_, numb, strs := numbers()
	fmt.Println(numb, strs, area)
	fmt.Println(a, b, c)

	const (
		a = iota
		b
		c
		d = "1133f"
		e
		f = false
		g = iota
		h
	)
	fmt.Println(a, b, c, d, e, f, g, h)

	const (
		i = 1 << iota
		j = 3 << iota
		k //<<n==*(2^n)。
		l
	)
	fmt.Println(i, j, k, l)

	//*************** switch使用
	var grade = "B"
	var marks int = 90
	switch marks {
	case 90:
		grade = "A"
	case 60, 70:
		grade = "C"
	case 80:
		grade = "B"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("优秀")
	case grade == "B":
		fmt.Println("良好")
	case grade == "C":
		fmt.Println("差")
	case grade == "D":
		fmt.Println("不及格")
	}
	fmt.Printf("您的等级是%s\n", grade)

	var x interface{}
	switch aa := x.(type) {
	case nil:
		fmt.Printf("x的类型是 %T\n", aa)
	case int:
		fmt.Println("x的类型是 int")
	case float32:
		fmt.Println("x的类型是 float32")
	case func(int) float64:
		fmt.Println("x的类型是 func(int)")
	case bool, string:
		fmt.Println("x的类型是 bool/string")
	default:
		fmt.Println("x的类型是 未知")
	}

	//fallthrough
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
	default:
		fmt.Println("6、case 条件语句为 默认")
	}

	//chan select test
	//这里体现的知识点：1 go并发执行管道，一个是压入int  最后压入boolean   2 select语句：会平均随机的进行执行case s/c，如果执行完成  就会进入boolean 然后完成
	ch := make(chan int)
	c1 := 0
	stopCh := make(chan bool)
	go Chann(ch, stopCh)

	for {
		select {
		case c1 = <-ch:
			fmt.Println("received C:", c1)
		case s := <-ch:
			fmt.Println("Receive S", s)
		case _ = <-stopCh:
			fmt.Println("go end")
			goto end
		}
	}
end:
}

// ?方法的第一个括号是传参使用，第二个括号是返回值使用，这是其中一种使用方式
func numbers() (int, int, string) {
	a, b, c := 1, 2, "aaaad"
	return a, b, c
}

func Chann(ch chan int, stopCh chan bool) {

	for j := 0; j < 10; j++ {
		ch <- j
		time.Sleep(time.Second)
	}
	stopCh <- false

}

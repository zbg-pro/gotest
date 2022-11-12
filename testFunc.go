package main

import (
	"awesomeProject/interfc"
	"awesomeProject/testclass"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	//1 Go 语言函数值传递值
	var a int = 100
	var b int = 200
	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)
	swap(a, b)
	fmt.Printf("交换后 a 的值为 : %d\n", a)
	fmt.Printf("交换后 b 的值为 : %d\n", b)

	//2 Go 语言函数引用传递值
	fmt.Println("Go 语言函数引用传递值")
	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)
	swap2(&a, &b) //这种用法实际上是通过&获取到变量的内存地址，例如：0xf840000040（每次的地址都可能不一样）。
	fmt.Printf("交换后 a 的值为 : %d\n", a)
	fmt.Printf("交换后 b 的值为 : %d\n", b)

	//3 Go 语言函数作为实参
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	//使用函数
	fmt.Println(getSquareRoot(2))

	var result1 int = testCallBack(1, callBack)
	fmt.Println("result1", result1)

	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})

	//4 Go 语言函数闭包
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence(0)
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	//5 Go 语言函数方法
	var c1 Circle
	c1.radius = 10.0
	fmt.Println("圆的面积：", c1.getArea()) //结构体包含的普通方法，直接返回float，是可以拿到c的属性，然后计算面积
	c1.changeRadiusBySelf(20)          //结构体包含方法：方法是指针类型的Circle，然后类似的拿到c，修改里面的值
	fmt.Println("c1的半径：", c1.radius)
	changeRadiusByNormalFunc(&c1, 30) //普通的把c1的内存地址交给方法，方法通过*c接收到，转换对象使用c，然后修改里面的值
	fmt.Println("c1的半径：", c1.radius)

	var totalSum int
	totalSum = testclass.Sum_test(2, 3)
	fmt.Println(totalSum, testclass.Total_sum)

	//6 Go 数组
	//几种申明方式：
	fmt.Println("几种申明方式")
	var arr1 []int
	fmt.Println(len(arr1))

	arr2 := []float64{1, 2}
	fmt.Println(len(arr2))

	arr3 := [2]int{}
	fmt.Println(len(arr3))

	arr4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr4))

	arr5 := [...]int{1: 3, 6: 4}
	fmt.Println(len(arr5))

	nums := []int{}
	var i int
	var length = 10
	for i = 0; i < length; i++ {
		for i := length - len(nums); i > 0; i-- {
			fmt.Print(" ")
		}
		nums = GetYangHuiTriangleNextLine(nums)
		fmt.Println(nums)
	}

	//7 Go 语言指针
	var ptest int = 10
	fmt.Println(&ptest)

	var ptest1 *int
	fmt.Printf("ptr 的值为 : %x\n", ptest1)

	ptest1 = &ptest
	fmt.Println(*ptest1)

	arrx := []int{10, 100, 1000}
	for i := 0; i < len(arrx); i++ {
		fmt.Printf("a[%d]=%d\n", i, arrx[i])
	}

	var ptr [MAX]*int
	for i := 0; i < MAX; i++ {
		ptr[i] = &arrx[i]
	}
	for i := 0; i < MAX; i++ {
		fmt.Printf("*ptr[%d]=%d\n", i, *ptr[i])
	}

	//8 Go 语言结构体
	fmt.Println(Books{bookId: 11, title: "三体", author: "刘", subject: "三体2"})
	fmt.Println(Books{"三体", "刘", "三体2", 11, "sdsdfsdf"})
	fmt.Println(Books{bookId: 11, author: "刘", subject: "三体2"})
	var book1 Books
	book1.bookId = 123
	book1.author = "刘"
	fmt.Println(book1)

	//作为函数参数
	printBook(book1)
	fmt.Println("参数化方法修改后：", book1)
	printBook2(&book1)
	fmt.Println("指针参数化方法修改后：", book1)

	//9 Go 语言切片(Slice)
	testSlice()

	//10 range
	testRange()

	//11 类型转换
	testCase()

	//12 Go 语言接口
	testInterface()

	//13 Go 错误处理
	testError(-1)

	testPanic()

	//14 GO 并发
	//go say("hello")
	//say("world")

	//15 通道（channel）
	//通道是用来传递数据的，如果没有指明方向 就是双向通道，如果<-是指明发送或者接收，可以结合goroutine 实现同步运行通讯
	testChan()
}

func testChan() {
	ch := make(chan int, 100)
	//ch <- 2 //如果这样写，go通道会认为没有准备接收者，然后就报错死锁
	go func() {
		ch <- 2
		ch <- 3
		ch <- 4
	}()
	a1 := <-ch
	fmt.Println("a", a1)
	b1 := <-ch
	fmt.Println("b", b1)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func testPanic() {
	//未使用 defer 调用匿名函数的 recover 捕获 panic：
	/*fmt.Println("this is a panic example")
	panic("this is a panic")
	r := recover()
	fmt.Printf("panic recover:%s", r)*/

	//使用defer调用匿名函数recover捕获panic异常
	//可以把defer看做是catch里面的方法包裹， 把panic看做是throw exception, recover里面收到的就是throw抛出后cach拿到的东西
	defer ccc()
	fmt.Println("this is a panic example")
	defer aaa(-1)
	panic("this is a panic")
	defer bbb()
}

func aaa(x int) {
	if r := recover(); r != nil {
		fmt.Printf("panic1 recover:%s\n", r)
		//panic("aaaaddd")
	}
}

func bbb() {
	if r := recover(); r != nil {
		fmt.Printf("panic2 recover:%s\n", r)
	}
}

func ccc() {
	if r := recover(); r != nil {
		fmt.Printf("panic3 recover:%s\n", r)
	}
}

func testError(x float64) {
	f, err := Sqrt(x)
	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println("error:", err)
	}

}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

func testInterface() {
	var phone interfc.Phone
	phone = new(interfc.NokiaPhone)

	phone.Call()
	rs, str := phone.Calc(1, 2)
	fmt.Println(rs, str)

	phone = new(interfc.IPhone)
	rs1, str1 := phone.Calc(1, 2)
	fmt.Println(rs1, str1)

	phone2 := interfc.IPhone{Version: "1.2.3.90"}

	var phone3 interfc.Phone
	phone3 = &interfc.NokiaPhone{}

	printReceive(phone2, "iphone")
	printReceive(phone3, "nokia")
	phone2.SetName("aaa")
	printReceive(phone2, "iphone")
	phone2.ChangeName("aa1")
	printReceive(phone2, "iphone")

	phone3.SetName("aaa")
	printReceive(phone3, "nokia")
	phone3.SetName("aa1")
	printReceive(phone3, "nokia")

}
func printReceive(p interfc.Phone, x string) {
	fmt.Println(x+" ", p)
}
func testCase() {
	var sum int = 17
	var count int = 5
	var mean float64

	mean = float64(count) / float64(sum)
	fmt.Println(mean)
}

func testRange() {
	fmt.Println(len(os.Args))
	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	nums := [3]int{5, 6, 7}
	for k, v := range nums {
		fmt.Println("源值地址：", &nums[k], " \t value的地址：", &v)
		fmt.Println("源值地址：", &nums[k], " \t value：", v)
	}
}

// 实际上就是个动态数组，抽象数组，数组不可用append增加，但是切片就可以
func testSlice() {
	var slice []int = make([]int, 2)
	fmt.Println(len(slice))
	slice1 := make([]int, 3)
	fmt.Println(len(slice1))

	arr := [4]int{1, 2, 3} //这是数组的申明，如果不指定大小，下面赋值会报错，数组长度申明时候不可变
	arr[3] = 123
	fmt.Println(arr)

	slice2 := arr[1:3]
	slice2 = append(slice2, 5)
	fmt.Println(slice2) //[2 3 5]

	var numbers = make([]int, 3, 5)
	numbers[0] = 0
	fmt.Println("添加前：", cap(numbers))
	//numbers[3] = 3               //错误的
	numbers = append(numbers, 4) //他们的增加是在原始数组的基础上追加元素
	numbers = append(numbers, 5)
	numbers = append(numbers, 6)
	numbers = append(numbers, 9)
	numbers = append(numbers, 11)
	numbers = append(numbers, 13)
	numbers = append(numbers, 15)
	fmt.Println("添加后：", cap(numbers))

	fmt.Println(numbers, len(numbers), cap(numbers))
	var numbers2 []int
	fmt.Println(numbers2)
	if numbers2 == nil {
		fmt.Printf("切片是空的")
	}

	number := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	number3 := number[2:5]
	printSlice(number3)

	var num []int
	printSlice(num)
	num = append(num, 0)
	printSlice(num)
	num = append(num, 1)
	printSlice(num)

	num = append(num, 2)
	printSlice(num)
	num = append(num, 3)
	printSlice(num)
	num = append(num, 4) // 可以看出，容量不够时，cap会自动扩容到2倍
	printSlice(num)

}
func printSlice(x []int) {
	fmt.Printf("len=%d  cap=%d   slice=%v\n", len(x), cap(x), x)
}
func printBook(book1 Books) {
	fmt.Println("参数化打印：", book1)
	book1.subject = "三体2"
}

func printBook2(book1 *Books) {
	fmt.Println("*book1指针参数化接收：", *book1)
	fmt.Println("book1指针参数化接收：", book1)
	fmt.Println("book1.author指针参数化接收：", book1.author)
	fmt.Println("(*book1).author指针参数化接收：", (*book1).author)
	book1.subject = "三体3" //&x给到-->申明为*的指针参数，指针参数操作和普通的一样，但是输出需要加*才能是实际值，否则就是地址

	var i int = 2
	var ptr *int
	ptr = &i
	fmt.Println("i=", i)
	fmt.Println("&i=", &i)
	fmt.Println("ptr=", ptr)
	fmt.Println("*ptr=", *ptr)

	book2 := Books{title: "三体。", subject: "三体3"}
	var ptrbook *Books
	ptrbook = &book2
	fmt.Println("&book2=", &book2)
	fmt.Println("ptrbook=", ptrbook) //对于结构化的内存地址打印出现的就是签名加个&的结构化值打印，奇怪？
	fmt.Println("&ptrbook=", *ptrbook)
	book2.PublicVar = "山西出版社1212dfsdf"
	result, _ := json.Marshal(book2)
	fmt.Println("json:", string(result))
	result2, _ := json.Marshal(ptrbook)
	fmt.Println("ptrbook json::", string(result2))
	fmt.Println("&ptrbook:", &ptrbook)
}

type Books struct {
	title     string
	author    string
	subject   string
	bookId    int
	PublicVar string
}

const MAX int = 3

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) changeRadiusBySelf(radius float64) {
	c.radius = radius
}

func changeRadiusByNormalFunc(c *Circle, radius float64) {
	c.radius = radius
}

// 定义相互交换值的函数
func swap(x, y int) (int, int) {
	var temp int
	temp = x
	x = y
	y = temp
	return x, y
}

func swap2(x, y *int) (*int, *int) { //*这种变量的申明类型 如果接住了内存地址，*x就是内存地址赋值结果，x就是实际值
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y
	*y = temp
	return x, y
}

func testCallBack(x int, callBack func(tt int) int) int { //实际这里的 函数参数的参数是没办法使用的，因为他没对函数参数进行实现，函数参数的实现在具体的callback方法里面
	fmt.Printf("我是回调2，x：%d\n", x)
	return callBack(x + 1)
}

func callBack(x int) int {
	fmt.Printf("callBack我是回调，x：%d\n", x)
	return x + 3
}

func getSequence(x int) func() int { //这里func()指的是返回一个函数，函数的返回值是int
	i := 0
	return func() int {
		i += 1
		return i + x
	}
}

func GetYangHuiTriangleNextLine(inArr []int) []int {
	var out []int
	arrLen := len(inArr)
	out = append(out, 1)
	if arrLen == 0 {
		return out
	}

	for i := 0; i < arrLen-1; i++ {
		out = append(out, inArr[i]+inArr[i+1])
	}
	out = append(out, 1)
	return out
}

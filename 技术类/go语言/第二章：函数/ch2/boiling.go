package main

import "fmt"

func sum(a, b int) (sum int) {
	sum = a + b
	return sum
}

//2.1.2多个返回值测试
func swap(a, b int) (int, int) {
	return a, b
}

//2.1.3测试参数为普通值的传递
func chvalue(a int) int {
	a = a + 1
	return a
}

//测试参数为指针的传递
func chpointer(a *int) {
	*a = *a + 1
	return
}

//2.1.4不定参数传递
func sum2(arr ...int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

func main() {
	//2.1.4fmt.Printf("不定参数传递%d\n", sum2(1, 2, 3, 4, 5))
	arr := []int{1, 2, 3, 4}
	slice := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("数组当作可变参数传入%d\n", sum2(arr...))
	fmt.Printf("可变数组当作可变参数传入%d\n", sum2(slice[:]...))
}

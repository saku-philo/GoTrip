package main

import "fmt"

func add(x, y int) int {
	return x * y
}

// 複数戻り値を返す事が可能
func swap(x, y string) (string, string) {
	return y, x
}

// 戻り値に名前をつける事が可能
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	a, b := swap("hello", "Go")
	fmt.Println(a, b)
	fmt.Println(split(17))
}

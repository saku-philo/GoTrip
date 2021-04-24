package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { // 「;」は省略可能, 条件省略で無限ループ
		sum += i
	}
	fmt.Println(sum)
}

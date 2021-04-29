package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// 文章をスペースで区切って配列に格納
	f := strings.Fields(s)
	m := make(map[string]int)

	for i := 0; i < len(f); i++ {
		key := f[i]

		// 同じkeyがあれば要素数をインクリメント
		_, exist := m[key]
		if exist {
			m[key] += 1
			// なければ要素、要素数を追加
		} else {
			m[f[i]] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

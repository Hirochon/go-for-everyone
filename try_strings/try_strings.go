package try_strings

import (
	"fmt"
	"strings"
)

func TryStrings() {
	// 文字列がある文字列で始まるかどうかを判定する
	fmt.Println("strings.HasPrefix")
	fmt.Println(strings.HasPrefix("hello", "hel"))
	fmt.Println(strings.HasPrefix("ello", "hel"))

	// 文字列がある文字列で終わるかどうかを判定する
	fmt.Println("strings.HasSuffix")
	fmt.Println(strings.HasSuffix("hello", "lo"))
	fmt.Println(strings.HasSuffix("hell", "lo"))

	// 文字列がある文字列を含んでいるかどうかを判定する
	fmt.Println("strings.Contains")
	fmt.Println(strings.Contains("hello", "el"))
	fmt.Println(strings.Contains("hello", "ho"))

	// 文字列を空白区切りで分割する
	fmt.Println("strings.Fields")
	fmt.Printf("%q\n", strings.Fields("hello world"))

	// セパレーターで文字列を分割する
	fmt.Println("strings.Split")
	fmt.Printf("%q\n", strings.Split("hello world", " "))
	fmt.Printf("%q\n", strings.Split("check,hoge,owowo", ","))

	// セパレーターで文字列を分割する。最大nの長さで返す。
	fmt.Println("strings.SplitN")
	fmt.Printf("%q\n", strings.SplitN("hello world", " ", 2))
	fmt.Printf("%q\n", strings.SplitN("check,hoge,owowo", ",", 2))

	// 文字列の前後の空白文字列を取り除く
	fmt.Println("strings.TrimSpace")
	fmt.Println(strings.TrimSpace("  hello world  "))

	// 文字列前後からcutsetに含まれる文字を取り除く
	fmt.Println("strings.Trim")
	fmt.Println(strings.Trim("ほええhello world ほええ ", "ほええ"))

	// 文字列の置換を行う
	fmt.Println("strings.Replace")
	fmt.Println(strings.Replace("hello world", "world", "everyone", 1))
}

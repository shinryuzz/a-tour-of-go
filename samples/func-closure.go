package main

import "fmt"

// 関数をラップしてその関数を返す
// func(int) int 型の関数 ... int を受け取り int を返す関数 = return 時に宣言してる関数のこと
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder() // クロージャをインスタンス化
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), // 
			neg(-2*i),
		)
	}
}

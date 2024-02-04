package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 構造体 Vertex に Abs() というメソッドを付与している と解釈できる
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale ... ポインタレシーバーを持つメソッド
// レシーバの実体を関数によって変更できる
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Scale2 ... ポインターレシーバーではなく普通のメソッドとする場合
// もしポインタとして受け取らずにScale()したい場合、更新した値を返す必要がある
// さらにその返り値をvに代入しなければ、v は更新されない
// また、メソッドを呼び出すたびに変数をコピーするのはメモリ効率が良くない。とくに v が大きな構造体の時など
func (v Vertex) Scale2(f float64) Vertex{
	v.X = v.X * f
	v.Y = v.Y * f
	return v
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	// v = v.Scale2(10)
	fmt.Println(v.Abs())
}

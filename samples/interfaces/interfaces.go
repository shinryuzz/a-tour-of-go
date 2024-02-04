package main

import (
	"fmt"
	"math"
)

/*
interface(インタフェース)型は、メソッドのシグニチャの集まりで定義します。

そのメソッドの集まりを実装した値を、interface型の変数へ持たせることができます。

注意: この例は、35行目でエラーになります。 Abs メソッドが、 Vertex ではなく *Vertex の定義であり、 Vertex が Abser インタフェースを実装していないということになるためエラーとなります。

*/

type Abser interface {
	Abs() float64 // float64を返す Abs() を持っているという意
}


func main() {
	var a Abser 
	f := MyFloat(-math.Sqrt2) // math.Sqrt2 ... 2 の平方根を返す
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	fmt.Println(a)

	a = &v // a *Vertex implements Abser
	fmt.Println(a)

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v // Error!

	fmt.Println(a.Abs())
}

type MyFloat float64

// MyFloat 型の変数に Abs() メソッドを定義
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// Vertex 構造体に Abs() メソッドを定義
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

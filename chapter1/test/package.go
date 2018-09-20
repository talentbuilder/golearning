package main //如果换成别的包名会报错

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	exportName()
	fmt.Println(add(5,4))
	fmt.Println(addBrief(6,4))
	a, b:= swap("3","8")
	fmt.Println(a, b)
	fmt.Println(split(30))
	fmt.Println(declareVariable)

}

//导出名函数命名
func exportName(){
	fmt.Println(math.Pi) //函数和变量大写开头的就是公有的，否则就是私有的
}

//参数和返回值
func add(x int, y int) int {
	return x + y
}

//x和y都是int型时，可以像下面这么写
func addBrief(x, y int) int {
	return x + y
}

//返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

//将返回值命名为x和y，return就是空的了
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func declareVariable(){
	var a, b, c bool
	var i int
	fmt.Println(a, b, c, i)
}




// 包前的注释

// 包前的注释2
package testpack

import (
	"fmt"
)

// 包后的注释

// 这是类型
// +proto
type Typ struct { // 类型左括号注释
	Nam string // 名字
} // 类型右括号注释

// 函数1
// +build
func (t *Typ) Func1(a, b int, c int, d *Typ) string { // 函数左括号注释
	return fmt.Sprint("call func1", a, b, c)
	// aaa
} // 函数右括号注释

// 函数2
// +proto
func (t *Typ) Func2(s string) {

}

/*
   函数3
   呵呵
*/
func (t *Typ) Func3() string {
	return "" + "22"
}

/*
   普通的
*/
func Func4() string {
	return "333" + "22"
}

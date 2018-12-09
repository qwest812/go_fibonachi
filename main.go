package main

import "fmt"

func fibinachi(el int) int {
	//fmt.Println(el)
	if el == 0 {
		return 0
	}
	if el == 1 {
		return 1
	}
	return fibinachi(el-1) + fibinachi(el-2)
}
func show(val int) int {

	if val >= 0 {
		//var result =

		fmt.Println(fibinachi(val))
		return show(val - 1)
	}
	return 0

}

func main() {
	show(7)

}

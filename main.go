package main

import "fmt"

var people string = "hello"

func main() {

	var start_num int = 23
	start_num = add_ten(start_num)
	fmt.Println("the number is ", start_num)

}

// return number plus 1)
func add_ten(base_num int) int {
	return base_num + 10
}

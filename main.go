package main

import "fmt"

func main() {
	var a int = 10 + 20
	var b float64 = 3.14 * 0.5
	var c complex128 = 5 + 7i
	var d uint = 50 * 100

	fmt.Println(a) // Output: 10
	fmt.Println(b) // Output: 3.14
	fmt.Println(c) // Output: (5+7i)
	fmt.Println(d)

	var isTrue bool = true
	var isFalse bool = false

	fmt.Println(isTrue)  // Output: true
	fmt.Println(isFalse) // Output: false

	var str string = "Hello, Go!"

	fmt.Println(str)

	var arr [10]string = [10]string{"Joseph", "Jesus", "John", "James", "Jeremiah", "Judah", "Joel", "Jonah", "Jude", "Jireh"}

	fmt.Println(arr)
}

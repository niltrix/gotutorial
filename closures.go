package main

import "fmt"

func intSeq() func() int {
	i := 0
	fmt.Println(&i)
	return func() int {
		i++
		fmt.Println(&i)
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt1 := intSeq()
	fmt.Println(nextInt1())
	nextInt2 := intSeq()
	fmt.Println(nextInt2())

	newInt := intSeq()
	fmt.Println(newInt())
}

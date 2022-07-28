package main

import (
	"errors"
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can NOT use 0 as denominator")
	}
	return a / b, nil
}

func divideWithNamedRParam(a, b int) (result int, err error) {
	if b == 0 {
		result = 0
		err = errors.New("Can NOT use 0 as denominator")
	} else {
		result = a / b
	}
	return
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Total : ", total)
}

func min(nums ...int) {
	fmt.Print(nums, " ")
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}
	fmt.Println("Min : ", min)
}

func Functest1() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)

	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)

	res, err := divide(4, 0)
	if err != nil {
		fmt.Println("ERROR : ", err.Error())
	} else {
		fmt.Println("4/2 = ", res)
	}
	res2, err2 := divideWithNamedRParam(4, 1)
	fmt.Println(res2, err2)

	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
	min(nums...)
}

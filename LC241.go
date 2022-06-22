package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(input string) []int {
	if isDigit(input) {
		temp, _ := strconv.Atoi(input)
		return []int{temp}
	}

	var res []int
	for index, c := range input {
		tmpC := string(c)
		if tmpC == "+" || tmpC == "-" || tmpC == "*" {
			left := diffWaysToCompute(input[:index])
			right := diffWaysToCompute(input[index+1:])
			for _, leftNum := range left {
				for _, rightNum := range right {
					var addNum int
					if tmpC == "+" {
						addNum = leftNum + rightNum
					} else if tmpC == "-" {
						addNum = leftNum - rightNum
					} else {
						addNum = leftNum * rightNum
					}
					res = append(res, addNum)
				}
			}
		}
	}
	return res
}

func isDigit(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return true
}

func mian() {
	fmt.Println(diffWaysToCompute("2-1-1"))
}

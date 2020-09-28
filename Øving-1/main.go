package main

import (
	"fmt"
	"os"
	"strconv"
)

func makeList(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Need arguments: <startValue>, <endValue>, <threadCount>")
		return
	}

	startValueString := os.Args[1]
	endValueString := os.Args[2]
	threadCountString := os.Args[3]

	startValue, err := strconv.Atoi(startValueString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	endValue, err := strconv.Atoi(endValueString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	threadCount, err := strconv.Atoi(threadCountString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	values := makeList(startValue, endValue)

}

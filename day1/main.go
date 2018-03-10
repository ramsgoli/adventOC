package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func byteArrayToIntArray(b []byte) []int {
	l := len(b)

	c := make([]int, 0)
	for i := 0; i < l; i++ {
		if s, err := strconv.Atoi(string(b[i])); err == nil {
			c = append(c, s)
		}
	}
	return c
}

func main() {
	var sum = 0
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ints := byteArrayToIntArray(content)
	length := len(ints)
	for i := 0; i < length-1; i++ {
		if ints[i] == ints[i+1] {
			sum += int(ints[i])
		}
	}

	// check if last element equals first
	if ints[length-1] == ints[0] {
		sum += int(ints[length-1])
	}

	fmt.Printf("The sum is: %d\n", sum)

}

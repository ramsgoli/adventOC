package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func stringToIntArray(s string) []int {
	nums := strings.Split(s, "\t")
	l := len(nums)

	c := make([]int, 0)
	for i := 0; i < l; i++ {
		if s, err := strconv.Atoi(string(nums[i])); err == nil {
			c = append(c, s)
		}
	}
	return c
}

func convertInputToMatrix(b []byte) [][]int {
	s := string(b)
	matrix := [][]int{}

	rows := strings.Split(s, "\n")

	for _, row := range rows {
		matrix = append(matrix, stringToIntArray(row))
	}

	return matrix
}

func main() {
	checksum := 0
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	nums := convertInputToMatrix(content)

	for _, row := range nums {
		min := row[0]
		max := row[0]
		for i := 0; i < len(row); i++ {
			if row[i] < min {
				min = row[i]
			}
			if row[i] > max {
				max = row[i]
			}
		}
		checksum += max - min
	}

	fmt.Printf("The checksum is: %d\n", checksum)
}

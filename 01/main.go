package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	data, err := readData("data.txt")
	if err != nil {
		panic(err)
	}
	processNums := []int{}

	for i := 0; i < len(data)-2; i++ {
		currentSum := data[i] + data[i+1] + data[i+2]
		processNums = append(processNums, currentSum)
	}

	prev := 0
	acc := -1

	for _, va := range processNums {
		if va > prev {
			acc += 1
		}
		prev = va
	}

	fmt.Println(acc)
}

func readData(fileName string) ([]int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int

	for {

		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			return nil, err
		}

		nums = append(nums, perline)
	}
	return nums, nil
}

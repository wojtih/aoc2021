package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := loadData()
	if err != nil {
		panic(err)
	}

	rowSize := len(data[0])

	invert := [][]string{}

	for i := 0; i < rowSize; i++ {
		inRow := []string{}
		for row := range data {
			inRow = append(inRow, data[row][i])
		}
		invert = append(invert, inRow)
	}
	var gamma []string

	rowSize = len(invert[0])
	for _, row := range invert {
		oneCount := 0
		zeroCount := 0
		for _, dig := range row {
			if dig == "1" {
				oneCount += 1
			} else {
				zeroCount += 1
			}

		}
		if oneCount > zeroCount {
			gamma = append(gamma, "1")
		} else {
			gamma = append(gamma, "0")
		}
	}
	var epsilon []string
	for _, c := range gamma {
		if c == "1" {
			epsilon = append(epsilon, "0")
		} else {
			epsilon = append(epsilon, "1")
		}
	}
	gammaBit, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	epsilonBit, _ := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)
	fmt.Println(gammaBit * epsilonBit)
}

func loadData() ([][]string, error) {
	file, err := os.Open("data.txt")
	if err != nil {
		return nil, err
	}

	nums := [][]string{}
	var row string
	for {
		_, err := fmt.Fscanf(file, "%s\n", &row) // give a patter to scan

		if err != nil {

			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, strings.Split(row, ""))

	}
	return nums, nil
}

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	data, err := readData("data_test.txt")
	if err != nil {
		panic(err)
	}

	c := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	for _, bing := range c {
		for _, card := range data {
			for _, raw := range card {
				fmt.Println(raw)
			}
			fmt.Println("")
		}
		time.Sleep(1 * time.Seconds)
		fmt.Println("-------------------------")
		for cardNum, x := range data {
			fmt.Printf("Bing: %d\n", bing)
			found := []int{}
		quit:
			for i, raw := range x {
				for j, elem := range raw {
					if elem == bing {
						found = []int{i, j}
						break quit
					}
				}
			}
			if len(found) > 0 {
				fmt.Println("Znalazłem cyferkę")
				fmt.Println(found)
				x[found[0]][found[1]] = 0
				for _, raw := range x {
					fmt.Println(raw)
				}

				result := false
				for i, raw := range x {
					if all(raw, 0) {
						result = true
						fmt.Printf("KArta %d wygrywa \n", cardNum)
						fmt.Printf("MAMAY BINGO W RZEDZIE %d\n", i)
						break
					}
				}
				for i := 0; i < len(x); i++ {
					result = true
					for j := 0; j < len(x); j++ {
						if x[j][i] != 0 {
							result = false
							break
						}
					}
					if result == true {
						fmt.Printf("KArta %d wygrywa \n", cardNum)
						fmt.Printf("MAMAY BINGO W KOLUMNIE %d\n", i)
						break
					}
				}
				for i := 0; i < len(x); i++ {
					result = true
					if x[i][i] != 0 {
						result = false
						break
					}
				}
				if result == true {
					fmt.Printf("KArta %d wygrywa \n", cardNum)
					fmt.Println("MAMAY BINGO PO PRZEKONTNEJ")
					break
				}
				if result {

					os.Exit(0)
				}
			}
			found = []int{}
		}
	}
}

func all(s []int, e int) bool {
	res := true
	for _, a := range s {
		if a != e {
			res = false
			return res
		}
	}
	return res
}

func readData(fileName string) ([][][]int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	nums := [][][]int{}
	card := [][]int{}
	raw := []int{}
	var n int
	for {
		co, err := fmt.Fscanf(file, "%d", &n) // give a patter to scan
		if co == 0 {
			break
		}
		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				break // stop reading the file
			}
			return nil, err
		}

		fmt.Println(n)
		raw = append(raw, n)
		if len(raw) == 5 {
			card = append(card, raw)
			raw = []int{}
			if len(card) == 5 {
				nums = append(nums, card)
				card = [][]int{}
			}
		}
	}
	return nums, nil
}

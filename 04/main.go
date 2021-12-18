package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fullData, err := readData("data_test.txt")
	if err != nil {
		panic(err)
	}

	c := fullData.numbers
	fmt.Println(c)
	data := fullData.data
	for _, bing := range c {
		printData(data)
		time.Sleep(1000)
		fmt.Println("************************************")
		fmt.Println(len(data))
		fmt.Println("************************************")
		fmt.Printf("Bing: %d\n", bing)

		for cardNum, x := range data {
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
				fmt.Printf("Znalazłem cyferkę. x: %d; y: %d\n", found[0], found[1])
				x[found[0]][found[1]] = bing * -1

				result := false
				for _, raw := range x {
					if all(raw) {
						result = true
						fmt.Printf("Karta %d wygrywa. Bing: %d\n", cardNum, bing)
						break
					}
				}
				if !result {
					for i := 0; i < len(x); i++ {
						result = true
						for j := 0; j < len(x); j++ {
							if x[j][i] >= 0 {
								result = false
								break
							}
						}
						if result {
							fmt.Printf("Karta %d wygrywa. Bing: %d\n", cardNum, bing)
							break
						}
					}
				}
				if !result {
					for i := 0; i < len(x); i++ {
						result = true
						if x[i][i] >= 0 {
							result = false
							break
						}
					}
				}
				if result {
					fmt.Println("------------------------------")
					printData(data)

					if len(x) == 1 {
						sum := sumCard(data[cardNum])
						fmt.Printf("Karta %d wygrywa. Suma: %d. Bing: %d\n", cardNum, sum, bing)
						fmt.Println(sum * bing)
						os.Exit(0)
					} else {
						fmt.Println(len(x))
						RemoveIndex(x, cardNum)
					}
				}
			}
			fmt.Printf("Nie znalazłem cyferki %d\n", bing)
			found = []int{}
		}
	}
}
func RemoveIndex(s [][]int, index int) []int {
	ret := make([][]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
func all(s []int) bool {
	res := true
	for _, a := range s {
		if a >= 0 {
			res = false
			return res
		}
	}
	return res
}

func sumCard(card [][]int) int {
	sum := 0
	for _, raw := range card {
		for _, i := range raw {
			if i >= 0 {
				sum += i
			}
		}
	}
	return sum
}

type data struct {
	numbers []int
	data    [][][]int
}

func readData(fileName string) (data, error) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	nums := [][][]int{}
	card := [][]int{}

	scanner := bufio.NewScanner(file)
	res := data{}
	scanner.Scan()
	raw := []int{}

	for _, f := range strings.Split(scanner.Text(), ",") {
		i, err := strconv.Atoi(f)
		if err == nil {
			raw = append(raw, i)
		}
	}
	res.numbers = raw

	for scanner.Scan() {
		raw := numbers(scanner.Text())
		// co, err := fmt.Fscanf(line, "%d %d %d %d %d\n", &n1, &n2, &n3, &n4, &n5) // give a patter to scan
		fmt.Println(raw)

		if len(raw) == 5 {
			card = append(card, raw)
		}

		if len(card) == 5 {
			nums = append(nums, card)
			card = [][]int{}
		}

	}
	res.data = nums
	return res, nil
}

func printData(data [][][]int) {
	for _, card := range data {
		for _, raw := range card {

			fmt.Print("|")
			fmt.Print(raw)
			fmt.Println("|")
		}
		fmt.Println("-------------------------")
	}
}

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

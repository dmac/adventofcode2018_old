package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(one(input))
	fmt.Println(two(input))
}

func readInput(path string) ([]int, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	if err != nil {
		return nil, err
	}
	var input []int
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		input = append(input, n)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return input, nil
}

func one(input []int) int {
	sum := 0
	for _, n := range input {
		sum += n
	}
	return sum
}

func two(input []int) int {
	m := make(map[int]struct{})
	i := 0
	sum := 0
	for {
		sum += input[i]
		if _, ok := m[sum]; ok {
			return sum
		}
		m[sum] = struct{}{}
		i = (i + 1) % len(input)
	}
}

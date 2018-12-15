package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	n := count(input, 2)
	m := count(input, 3)
	fmt.Println(n * m)
	fmt.Println(findSimilar(input))
}

func readInput(path string) ([]string, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	if err != nil {
		return nil, err
	}
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return input, nil
}

func count(input []string, repeats int) int {
	count := 0
	for _, s := range input {
		m := make(map[rune]int)
		for _, r := range s {
			m[r]++
		}
		for _, n := range m {
			if n == repeats {
				count++
				break
			}
		}
	}
	return count
}

func findSimilar(input []string) string {
	for _, s0 := range input {
		for _, s1 := range input {
			if s0 == s1 {
				continue
			}
			if len(s0) != len(s1) {
				continue
			}
			idx := -1
			found := false
			for i := 0; i < len(s0); i++ {
				if s0[i] != s1[i] {
					if idx >= 0 {
						found = false
						break
					}
					idx = i
					found = true
				}
			}
			if found {
				return s0[:idx] + s0[idx+1:]
			}
		}
	}
	return ""
}

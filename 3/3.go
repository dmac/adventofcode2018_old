package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	id     int
	row    int
	col    int
	width  int
	height int
}

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(one(input))
	fmt.Println(two(input))
}

func readInput(path string) ([]claim, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	if err != nil {
		return nil, err
	}
	var input []claim
	for scanner.Scan() {
		var clm claim
		var err error
		fields := strings.Fields(scanner.Text())
		clm.id, err = strconv.Atoi(fields[0][1:])
		if err != nil {
			return nil, err
		}
		colrow := strings.Split(fields[2], ",")
		clm.col, err = strconv.Atoi(colrow[0])
		if err != nil {
			return nil, err
		}
		clm.row, err = strconv.Atoi(colrow[1][:len(colrow[1])-1])
		if err != nil {
			return nil, err
		}
		wh := strings.Split(fields[3], "x")
		clm.width, err = strconv.Atoi(wh[0])
		if err != nil {
			return nil, err
		}
		clm.height, err = strconv.Atoi(wh[1])
		if err != nil {
			return nil, err
		}
		input = append(input, clm)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return input, nil
}

func one(input []claim) int {
	maxrow := 0
	maxcol := 0
	for _, clm := range input {
		row := clm.row + clm.height
		if row > maxrow {
			maxrow = row
		}
		col := clm.col + clm.width
		if col > maxcol {
			maxcol = col
		}
	}
	cloth := make([]int, maxrow*maxcol)
	for _, clm := range input {
		for r := clm.row; r < clm.row+clm.height; r++ {
			for c := clm.col; c < clm.col+clm.width; c++ {
				i := r*maxcol + c
				cloth[i]++
			}
		}
	}
	count := 0
	for _, n := range cloth {
		if n > 1 {
			count++
		}
	}
	return count
}

func two(input []claim) int {
	maxrow := 0
	maxcol := 0
	for _, clm := range input {
		row := clm.row + clm.height
		if row > maxrow {
			maxrow = row
		}
		col := clm.col + clm.width
		if col > maxcol {
			maxcol = col
		}
	}
	cloth := make([][]int, maxrow*maxcol)
	overlapIDs := make(map[int]struct{})
	for _, clm := range input {
		for r := clm.row; r < clm.row+clm.height; r++ {
			for c := clm.col; c < clm.col+clm.width; c++ {
				i := r*maxcol + c
				if len(cloth[i]) > 0 {
					overlapIDs[clm.id] = struct{}{}
					for _, id := range cloth[i] {
						overlapIDs[id] = struct{}{}
					}
				}
				cloth[i] = append(cloth[i], clm.id)
			}
		}
	}
	for _, clm := range input {
		if _, ok := overlapIDs[clm.id]; !ok {
			return clm.id
		}
	}
	return 0
}

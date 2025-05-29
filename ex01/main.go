package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Общая сложность алгоритма O(n log n)
func main() {
	sections, err := parseSections("data_prog_contest_problem_1.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	points := 0
	
	if len(sections) >= 1 {
		sort.Slice(sections, func(i, j int) bool {
			return sections[i][1] < sections[j][1]
		})

		currentPoint := sections[0][1]
		points = 1

		for _, seg := range sections {
			if seg[0] > currentPoint {
				currentPoint = seg[1]
				points++
			}
		}
	}

	fmt.Println(points)
}

func parseSections(filename string) ([][2]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file: %w", err)
	}
	defer file.Close()

	var sections [][2]int
	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return nil, fmt.Errorf("Empty file")
	}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("Incorrect input data: %w", err)
		}

		start, err1 := strconv.Atoi(fields[0])
		end, err2 := strconv.Atoi(fields[1])
		if err1 != nil || err2 != nil || start < 0 || end < 0 {
			return nil, fmt.Errorf("Incorrect input data: %w", err)
		}
		if start > end {
			return nil, fmt.Errorf("Incorrect input data: %w", err)
		}
		sections = append(sections, [2]int{start, end})
	}

	return sections, nil
}

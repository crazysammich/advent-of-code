package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	file, err := os.Open("day-1/day-1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	left_list := []int{}
	right_list := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		left, _ := strconv.Atoi(words[0])
		right, _ := strconv.Atoi(words[len(words)-1])

		left_list = append(left_list, left)
		right_list = append(right_list, right)
	}
	sort.Ints(left_list)
	sort.Ints(right_list)

	// partOne(left_list, right_list)
	partTwo(left_list, right_list)

}

func partOne(left_list []int, right_list []int) {
	total_distance := 0
	for i, n := range left_list {
		distance := math.Abs(float64(n - right_list[i]))
		total_distance += int(distance)
	}
	fmt.Println(total_distance)
}

func partTwo(left_list []int, right_list []int) {
	counts := map[int]int{}

	for _, n := range left_list {
		if _, ok := counts[n]; ok {
			continue
		}
		counts[n] = 0
	}

	for _, n := range right_list {
		if _, ok := counts[n]; ok {
			counts[n]++
		}
	}

	sim_score := 0
	for _, n := range left_list {
		count := counts[n]
		sim_score += n * count
	}
	fmt.Println(sim_score)
}

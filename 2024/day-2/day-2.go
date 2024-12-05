package day2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	file, err := os.Open("day-2/day-2.test")
	if err != nil {
		log.Fatal(err)
	}

	totalSafe := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")

		for i := 0; i < len(levels); i++ {
			newLevels := append([]string{}, levels[:i]...)
			newLevels = append(newLevels, levels[i+1:]...)
			fmt.Println(newLevels)

			if areLevelsValid(newLevels) {
				totalSafe++
				break
			}
		}
	}
	fmt.Println(totalSafe)
}

func areLevelsValid(levels []string) bool {
	isIncreasing := false
	isDecreasing := false

	for i := 0; i < len(levels)-1; i++ {
		n1, _ := strconv.Atoi(levels[i])
		n2, _ := strconv.Atoi(levels[i+1])
		diff := n1 - n2

		if diff < -3 || diff == 0 || diff > 3 {
			return false
		}
		if diff > 0 {
			isIncreasing = true
		} else if diff < 0 {
			isDecreasing = true
		}
	}
	if isDecreasing && isIncreasing {
		return false
	}
	return true
}

func isLevelInSafetyRange(a int, b int) bool {
	diff := math.Abs(float64(a) - float64(b))
	return diff >= 1 && diff <= 3
}

// func isDecreasing(a int, b int) int {
// 	diff := a - b
// 	if diff >= 1 && diff <= 3 {
// 		return -1
// 	}
// 	return 0
// }

// func isIncreasing(a int, b int) int {
// 	diff := a - b
// 	if diff <= -1 && diff >= -3 {
// 		return 1
// 	}
// 	return 0
// }

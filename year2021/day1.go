package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//sample := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	cwd, _ := os.Getwd()
	path := cwd + "/year2021/day1.list"

	nums, err := readLines(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	increasedCount := getIncreasedCount(nums)
	fmt.Println(increasedCount)

	windowSums := getWindowSums(nums)
	increasedWindowCount := getIncreasedCount(windowSums)
	fmt.Println(increasedWindowCount)
}

func getIncreasedCount(nums []int) int {
	increasedCount := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			increasedCount += 1
		}
	}

	return increasedCount
}

func getWindowSums(nums []int) []int {
	length := len(nums)
	windowSums := make([]int, 0, length - 2)

	for i := 0; i < length - 2; i++ {
		sum := nums[i] + nums[i+1] + nums[i+2]
		windowSums = append(windowSums, sum)
	}

	return windowSums
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		lines = append(lines, num)
	}
	return lines, scanner.Err()
}

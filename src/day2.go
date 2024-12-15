package src

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day2 struct{}

func (d *Day2) Loaded() {
	fmt.Println("⭐️")
}

func (d *Day2) PartOne() {
	file, err := os.Open("inputs/day_2.txt")
	if err != nil {
		fmt.Println("Erro openning:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeReports := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		lineArr := []int{}

		for _, str := range parts {
			n, _ := strconv.Atoi(str)
			lineArr = append(lineArr, n)

		}

		if d.isSafe(lineArr) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}

func (d *Day2) PartTwo() {
	file, err := os.Open("inputs/day_2.txt")
	if err != nil {
		fmt.Println("Erro openning:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeReports := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		lineArr := []int{}

		for _, str := range parts {
			n, _ := strconv.Atoi(str)
			lineArr = append(lineArr, n)

		}

		if d.isSafe(lineArr) {
			safeReports++
		} else {
			if d.isSafeWithDeletion(lineArr) {
				safeReports++
			}
		}
	}

	fmt.Println(safeReports)
}

func (d *Day2) isSafeWithDeletion(list []int) bool {
	for i := 0; i < len(list); i++ {

		// copy of the original list
		listCopy := make([]int, len(list))
		copy(listCopy, list)

		// remove current element from list
		tempList := removeElement(listCopy, i)

		if d.isSafe(tempList) {
			return true
		}

	}
	return false
}

func (d *Day2) isSafe(list []int) bool {
	direction := list[1] > list[0]
	limitIndex := len(list)

	if direction {
		// incresing

		for i := 1; i < limitIndex; i++ {
			diff := list[i] - list[i-1]
			if diff > 3 || diff < 1 {
				return false
			}
		}
		return true

	} else {
		// decreasing

		for i := 1; i < limitIndex; i++ {
			diff := list[i] - list[i-1]
			if diff < -3 || diff > -1 {
				return false
			}
		}
		return true
	}
}

package src

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Day3 struct{}

func (d *Day3) Loaded() {
	fmt.Println("⭐️")
}

func (d *Day3) PartOne() {
	file, err := os.Open("inputs/day_3.txt")
	if err != nil {
		fmt.Println("Erro openning:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	results := 0

	for scanner.Scan() {
		line := scanner.Text()

		regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
		matches := regex.FindAllString(line, -1)

		for _, match := range matches {
			results += d.mul(match)
		}
	}

	fmt.Println(results)
}

func (d *Day3) PartTwo() {
	file, err := os.Open("inputs/day_3.txt")
	if err != nil {
		fmt.Println("Erro openning:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	results := 0
	matchList := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		regex := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
		matches := regex.FindAllString(line, -1)

		matchList = append(matchList, matches...)
	}

	cleanedList := d.digest(matchList)

	for _, match := range cleanedList {
		results += d.mul(match)
	}

	fmt.Println(results)
}

func (d *Day3) digest(list []string) []string {
	disabled := false
	newList := []string{}

	for i := 0; i < len(list); i++ {

		e := list[i]

		if e == "do()" {
			disabled = false
		} else if e == "don't()" {
			disabled = true
		} else {
			if !disabled {
				newList = append(newList, e)
			}
		}
	}

	return newList
}

func (d *Day3) mul(input string) int {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := regex.FindStringSubmatch(input)

	num1, _ := strconv.Atoi(matches[1])
	num2, _ := strconv.Atoi(matches[2])

	return num1 * num2
}

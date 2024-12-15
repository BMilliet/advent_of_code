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

		// up, down or empty
		limitIndex := len(lineArr)
		direction := ""
		previous := 0
		isUnsafe := false

		for i := 0; i < limitIndex; i++ {

			e := lineArr[i]

			if i == 0 {
				previous = e
				continue
			}

			eval := e - previous

			if eval > 0 {
				if eval > 3 || direction == "down" {
					// unsafe
					isUnsafe = true
					break
				}

				direction = "up"
				previous = e
			} else if eval < 0 {
				if eval < -3 || direction == "up" {
					// unsafe
					isUnsafe = true
					break
				}

				direction = "down"
				previous = e
			} else {
				if direction != "flat" || e != previous {
					// unsafe
					isUnsafe = true
					break
				}
				direction = "flat"
				previous = e
			}

		}

		if !isUnsafe {
			safeReports += 1
		}
	}

	fmt.Println(safeReports)
}

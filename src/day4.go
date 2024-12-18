package src

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Day4 struct{}

func (d *Day4) Loaded() {
	fmt.Println("⭐️")
}

func (d *Day4) PartOne() {
	file, err := os.Open("inputs/day_4.txt")
	if err != nil {
		fmt.Println("Erro openning:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	grid := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		lineArr := []string{}

		for _, letter := range line {
			lineArr = append(lineArr, string(letter))
		}

		grid = append(grid, lineArr)
	}

	sum += d.findHorizontal(grid)
	sum += d.findVertical(grid)
	sum += d.findDiagonalTop(grid)
	sum += d.findDiagonalBottom(grid)

	fmt.Println(sum)
}

func (d *Day4) PartTwo() {
	file, err := os.Open("inputs/day_4.txt")
	if err != nil {
		fmt.Println("Erro openning:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	grid := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		lineArr := []string{}

		for _, letter := range line {
			lineArr = append(lineArr, string(letter))
		}

		grid = append(grid, lineArr)
	}

	sum = d.countXMas(grid)

	fmt.Println(sum)
}

func (d *Day4) countXMas(grid [][]string) int {
	count := 0

	rows := len(grid)
	cols := len(grid[0])

	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {

			pivot := grid[row][col]

			if pivot == "A" {
				topLeft := grid[row-1][col-1]
				topRight := grid[row-1][col+1]
				bottomLeft := grid[row+1][col-1]
				bottomRight := grid[row+1][col+1]

				if d.shouldCountX(topLeft, topRight, bottomLeft, bottomRight) {
					count++
				}
			}
		}
	}

	return count
}

func (d *Day4) shouldCountX(topL, topR, bottomL, bottomR string) bool {
	if topL == "M" && bottomR == "S" && bottomL == "M" && topR == "S" {
		return true
	}

	if topL == "S" && bottomR == "M" && bottomL == "S" && topR == "M" {
		return true
	}

	if topL == "M" && bottomR == "S" && bottomL == "S" && topR == "M" {
		return true
	}

	if topL == "S" && bottomR == "M" && bottomL == "M" && topR == "S" {
		return true
	}

	return false
}

func (d *Day4) findDiagonalTop(grid [][]string) int {
	rows := len(grid)
	cols := len(grid[0])

	diagonals := make([][]string, rows+cols-1)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			diagonalIndex := i + j
			diagonals[diagonalIndex] = append(diagonals[diagonalIndex], grid[i][j])
		}
	}

	return d.findHorizontal(diagonals)
}

func (d *Day4) findDiagonalBottom(grid [][]string) int {
	rows := len(grid)
	cols := len(grid[0])

	totalDiagonals := rows + cols - 1

	diagonals := make([][]string, totalDiagonals)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			diagonalIndex := (i - j) + (cols - 1)
			diagonals[diagonalIndex] = append(diagonals[diagonalIndex], grid[i][j])
		}
	}

	return d.findHorizontal(diagonals)
}

func (d *Day4) findVertical(grid [][]string) int {
	newGrid := [][]string{}

	limitCol := len(grid[0])
	limitRow := len(grid)

	for col := 0; col < limitCol; col++ {
		temp := []string{}

		for row := 0; row < limitRow; row++ {
			e := grid[row][col]
			temp = append(temp, e)
		}

		newGrid = append(newGrid, temp)
	}

	return d.findHorizontal(newGrid)
}

func (d *Day4) findHorizontal(grid [][]string) int {
	sum := 0

	for _, line := range grid {
		stringLine := strings.Join(line, "")

		regex := regexp.MustCompile(`XMAS`)
		matches := regex.FindAllString(stringLine, -1)

		sum += len(matches)

		regex = regexp.MustCompile(`SAMX`)
		matches = regex.FindAllString(stringLine, -1)

		sum += len(matches)
	}

	return sum
}

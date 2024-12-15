package src

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day1 struct{}

func (d *Day1) Loaded() {
	fmt.Println("⭐️")
}

func (d *Day1) PartOne() {
	file, err := os.Open("inputs/day_1.txt")
	if err != nil {
		fmt.Println("Erro openning:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	list1 := []int{}
	list2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		e1, _ := strconv.Atoi(parts[0])
		e2, _ := strconv.Atoi(parts[1])
		list1 = append(list1, e1)
		list2 = append(list2, e2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err)
	}

	list1 = quickSort(list1)
	list2 = quickSort(list2)
	limitIndex := len(list1)
	resp := 0

	for i := 0; i < limitIndex; i++ {
		e1 := list1[i]
		e2 := list2[i]
		dist := e1 - e2
		dist = absInt(dist)
		resp += dist
	}

	fmt.Println(resp)
}

func (d *Day1) PartTwo() {
	file, err := os.Open("inputs/day_1.txt")
	if err != nil {
		fmt.Println("Error openning:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	list1 := []int{}
	list2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		e1, _ := strconv.Atoi(parts[0])
		e2, _ := strconv.Atoi(parts[1])
		list1 = append(list1, e1)
		list2 = append(list2, e2)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err)
	}

	frequency := createFrequencyMap(list2)
	limitIndex := len(list1)
	resp := 0

	fmt.Println(frequency)

	for i := 0; i < limitIndex; i++ {

		e := list1[i]

		if f, ok := frequency[e]; ok {
			resp += e * f
		}

	}

	fmt.Println(resp)
}

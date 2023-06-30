package main

import (
	"bufio"
	"fmt"
	"os"
)

//read in a line (rucksack)
//split the rucksack into two compartments
//compare each item in compartment 1 to each item in compartment 2
//if two items are the same, find the priority of the item
//provide sum of priorities for all rucksacks
// Goal is to provide a single line to a function and receive the priority of the matching items. Main goroutine will keep tally
//Simple rucksack type with two compartments that contain slices of runes
type rucksack struct {
	compartment1 []rune
	compartment2 []rune
	commonItem   rune
}

func (r *rucksack) BuildCompartments(line string) {
	for i, item := range line {
		if i < len(line)/2 {
			r.compartment1 = append(r.compartment1, item)
		}
		if i >= len(line)/2 {
			r.compartment2 = append(r.compartment2, item)
		}
	}
}

func (r *rucksack) FindCommonItem() {
	for _, v := range r.compartment1 {
		for _, k := range r.compartment2 {
			if v == k {
				r.commonItem = v
			}
		}
	}
}

func main() {
	input, _ := os.Open("aoc22_day3_input")
	defer input.Close()
	scanner := bufio.NewScanner(input)
	tally := 0
	for scanner.Scan() {
		var rs rucksack
		rs.BuildCompartments(scanner.Text())
		rs.FindCommonItem()
		tally += priority(rs.commonItem)
		fmt.Println(string(rs.commonItem))
	}
	fmt.Println(tally)
}

func priority(item rune) int {

	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	}
	if item >= 'A' && item <= 'Z' {
		return int(item - 'A' + 27)
	}
	return 0
}

// Goal is to determine the Elf with the food that adds up to the most calories.  Data contains blank-line delimited entries. Each entry represents the number of calories for a
// food item.  Each collection of these values represents the elf they belong to.
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//build regex to check for a blank line
	blankLine := regexp.MustCompile(`(?m)^\s*$`)

	//read file and panic if we can't open it for some reason
	rawdata, err := ioutil.ReadFile("./input_data/day1_challenge1_input")
	if err != nil {
		panic(0)
	}

	//split the string returned from ioutil.ReadFile into a []string at each blankline
	//Index of this []string is the elf and the remaining strings will need to be split in to a separate []string. Creating a [][]string
	splitData := blankLine.Split(string(rawdata), -1)

	//Data structure used to represent an Elf and its food items. Index is the elf. Inside []int represent the calories of each food item
	var elfData [][]int

	// As we had to play with Strings, it is now time to convert the Strings in the data to Integers. Nested for loops to catch the 2-dimensional list
	for _, v := range splitData {
		tempSliceString := strings.Fields(v)
		var tempSliceInt []int
		for _, i := range tempSliceString {
			b, err := strconv.Atoi(i)
			if err != nil {
				panic(0)
			}
			tempSliceInt = append(tempSliceInt, b)
		}
		elfData = append(elfData, tempSliceInt)
	}

	//Now we can tally up the elements per Elf and store them in a new data structure. We can keep the Elf as the index and a simple int as the total of the calories

	var totalPerElf []int
	for _, v := range elfData {
		tally := 0
		for _, i := range v {
			tally += i
		}
		totalPerElf = append(totalPerElf, tally)
	}

	//Brute force Min/Max check
	max := totalPerElf[0]
	for _, v := range totalPerElf {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)

	//We've solved the first part of this puzzle.
	//The second part of the puzzel wants the total value of the top-3 elves' calories.

	sort.Ints(totalPerElf)

	topThreeElves := totalPerElf[len(totalPerElf)-3 : len(totalPerElf)]
	fmt.Println(topThreeElves)
	total := 0
	for _, v := range topThreeElves {
		total += v
	}
	fmt.Println(total)
}

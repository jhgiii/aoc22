package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Determine Move Needed given Outcome requirements
func DetermineMove(outcome int, player1 string) int {
	if outcome == 0 {
		switch {
		case player1 == "A":
			return 3
		case player1 == "B":
			return 1
		case player1 == "C":
			return 2
		}
	}
	if outcome == 3 {
		switch {
		case player1 == "A":
			return 1
		case player1 == "B":
			return 2
		case player1 == "C":
			return 3
		}
	}
	if outcome == 6 {
		switch {
		case player1 == "A":
			return 2
		case player1 == "B":
			return 3
		case player1 == "C":
			return 1
		}
	}
	return 0
}

//Determine outcome for player 2
func DetermineOutcome(game []string) int {
	//Determine if I should Win / Lose / Draw by looking at game[1]
	switch {
	case game[1] == "X":
		return 0
	case game[1] == "Y":
		return 3
	case game[1] == "Z":
		return 6
	}
	return 0
}

//helper function to convert the string form of input from AOC22 Day 2 to an Integer for easier comparisons
// Return functions represent the "score" for each option
// Rock: 1
// Paper: 2
// Scissors: 3
func ConvertToInt(letter string) int {
	switch {
	case letter == "A" || letter == "X":
		return 1
	case letter == "B" || letter == "Y":
		return 2
	case letter == "C" || letter == "Z":
		return 3
	default:
		return 0
	}
}

// Compares two integers where 1 = rock, 2 = paper, 3 = scissors and returns the value associated with Player2's result.
// Return: 0 -> Player 2 lose
// Return: 3 -> Player 2 Draw
// Return: 6 -> Player 2 Win
func CheckGameResult(player1, player2 int) int {
	switch {
	case player1 == player2:
		return 3
	case player1 == 1 && player2 == 2:
		return 6
	case player1 == 1 && player2 == 3:
		return 0
	case player1 == 2 && player2 == 1:
		return 0
	case player1 == 2 && player2 == 3:
		return 6
	case player1 == 3 && player2 == 1:
		return 6
	case player1 == 3 && player2 == 2:
		return 0
	}
	return 0
}

func main() {
	//Read input
	rawdata, err := ioutil.ReadFile("aoc22_day2_input")
	if err != nil {
		panic(0)
	}
	//Split string on newlines to build a []string where the index is the line number in the input file and the elements are the  games
	parsedData := strings.Split(string(rawdata), "\n")

	//Iterate through the list and build a [][]string where the outside index is the game, and the inside represents the moves played in the game
	var games [][]string
	for _, v := range parsedData {
		tempGame := strings.Fields(v)
		games = append(games, tempGame)
	}
	//iterate through the [][]string and convert to [][]int for easier comparisons to determine win/lose/draw
	var gamesInt [][]int
	for _, v := range games {
		var tempSliceInt []int
		tempSliceInt = append(tempSliceInt, ConvertToInt(v[0]))
		tempSliceInt = append(tempSliceInt, ConvertToInt(v[1]))

		gamesInt = append(gamesInt, tempSliceInt)
	}

	//iterate through the games and tally the score
	tally := 0
	for _, v := range gamesInt {
		//Add the value of the Selection (R=1, P=2, S=3)
		tally += v[1]
		//Add the result of the game (Win/Lose/Draw)
		tally += CheckGameResult(v[0], v[1])
	}
	fmt.Printf("Value for Day 2, puzzle 1: %v\n", tally)

	//Day 2 Puzzle 2 adds the requirement that the second value in the input indicates whether I should win/lose/draw where:
	//X -> I Lose
	//Y -> I Draw
	//Z -> I Win
	p2Tally := 0
	for _, v := range games {
		//Get the points for the outcome and add to the tally
		outcome := DetermineOutcome(v)
		p2Tally += outcome
		//Get the points for the required move and add to the tally
		move := DetermineMove(outcome, v[0])
		p2Tally += move
	}
	fmt.Printf("Value for Day2, Puzzle 2 Result: %v\n", p2Tally)

}

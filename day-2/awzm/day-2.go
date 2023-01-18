package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Shape string

const (
	Rock     Shape = "rock"
	Paper    Shape = "paper"
	Scissors Shape = "scissors"
)

func parseOpponentShape(s string) (Shape, error) {
	switch s {
	case "A":
		return Rock, nil
	case "B":
		return Paper, nil
	case "C":
		return Scissors, nil
	default:
		return "", errors.New("Invalid opponent shape")
	}
}

func parseYourShape(s string) (Shape, error) {
	switch s {
	case "X":
		return Rock, nil
	case "Y":
		return Paper, nil
	case "Z":
		return Scissors, nil
	default:
		return "", errors.New("Invalid your shape")
	}
}

func shapeScore(shape Shape) (int, error) {
	switch shape {
	case Rock:
		return 1, nil
	case Paper:
		return 2, nil
	case Scissors:
		return 3, nil
	}
	return 0, errors.New("Invalid shape")
}

func roundOutcome(you, opponent Shape) int {
	switch {
	case you == opponent:
		return 3 // draw
	case you == Rock && opponent == Scissors:
		return 6 // win
	case you == Paper && opponent == Rock:
		return 6 // win
	case you == Scissors && opponent == Paper:
		return 6 // win
	}
	return 0 // loss
}

func RockPaperScissors(you, opponent Shape) (int, error) {
	score, err := shapeScore(you)
	if err != nil {
		return 0, err
	}
	return score + roundOutcome(you, opponent), nil
}

func main() {
	file, err := os.Open("games.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		opponentShape, err := parseOpponentShape(line[0:1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		yourShape, err := parseYourShape(line[2:3])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		score, err := RockPaperScissors(yourShape, opponentShape)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		totalScore += score
	}

	fmt.Println("Total Score:", totalScore)
}

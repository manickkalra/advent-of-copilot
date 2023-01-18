package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Shape string
type Outcome string

const (
	Rock     Shape = "rock"
	Paper    Shape = "paper"
	Scissors Shape = "scissors"

	Win  Outcome = "win"
	Draw Outcome = "draw"
	Lose Outcome = "lose"
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

func parseYourOutcome(s string) (Outcome, error) {
	switch s {
	case "X":
		return Lose, nil
	case "Y":
		return Draw, nil
	case "Z":
		return Win, nil
	default:
		return "", errors.New("Invalid your outcome")
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

func inferYourShape(opponentShape Shape, yourOutcome Outcome) (Shape, error) {
	switch {
	case opponentShape == Rock && yourOutcome == Win:
		return Paper, nil
	case opponentShape == Rock && yourOutcome == Draw:
		return Rock, nil
	case opponentShape == Rock && yourOutcome == Lose:
		return Scissors, nil
	case opponentShape == Paper && yourOutcome == Win:
		return Scissors, nil
	case opponentShape == Paper && yourOutcome == Draw:
		return Paper, nil
	case opponentShape == Paper && yourOutcome == Lose:
		return Rock, nil
	case opponentShape == Scissors && yourOutcome == Win:
		return Rock, nil
	case opponentShape == Scissors && yourOutcome == Draw:
		return Scissors, nil
	case opponentShape == Scissors && yourOutcome == Lose:
		return Paper, nil
	}
	return "", errors.New("Invalid opponent shape or your outcome")
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
		yourOutcome, err := parseYourOutcome(line[2:3])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		yourShape, err := inferYourShape(opponentShape, yourOutcome)
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

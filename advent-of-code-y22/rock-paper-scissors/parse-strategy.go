package main

import (
	"bufio"
	"fmt"
	game "github.com/d-sauer/exploring-go/advent-of-code-y22/rock-paper-scissors/game"
	"io/fs"
	"regexp"
)

type PlayerMove struct {
	Player1Move string
	Player2Move string
}

type StrategyGuide struct {
	PlayerMove []PlayerMove
}

var regex = regexp.MustCompile(".*(?P<P1>[ABC]).*(?P<P2>[XYZ]).*")

var strategyMoves = map[string]game.Move{
	"A": game.Rock,
	"B": game.Paper,
	"C": game.Scissors,
	"X": game.Rock,
	"Y": game.Paper,
	"Z": game.Scissors,
}

var MoveScore = map[game.Move]int{
	game.Rock:     1,
	game.Paper:    2,
	game.Scissors: 3,
}

var OutcomeScore = map[game.Outcome]int{
	game.Win:  6,
	game.Draw: 3,
	game.Loss: 0,
}

func LoadStrategy(file fs.File, strategyGuide *StrategyGuide) error {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var line = 0
	player2Score := 0
	for fileScanner.Scan() {
		line++
		content := fileScanner.Text()

		match := regex.FindStringSubmatch(content)
		player2Score += strategyGuide.PlayRound(match[regex.SubexpIndex("P1")], match[regex.SubexpIndex("P2")])

	}
	fmt.Printf("Player 2 score: %d\n", player2Score)

	return nil
}

func (sg *StrategyGuide) PlayRound(player1Move string, player2Move string) (player2Score int) {
	player1 := strategyMoves[player1Move]

	player2 := planMove(player1, player2Move)
	_, p2o := game.Play(player1, player2)
	return MoveScore[player2] + OutcomeScore[p2o]
}

// planMove based on strategy
// X - you need to lose
// Y - you need to draw
// Z - you need to win
func planMove(player1Move game.Move, strategy string) game.Move {
	move := player1Move
	if strategy == "Z" {
		move = game.Move(mod(int(player1Move+1), 3))
	} else if strategy == "X" {
		move = game.Move(mod(int(player1Move-1), 3))
	}

	fmt.Printf("Player1: %d, player2 (%s): %d\n", player1Move, strategy, move)
	return move
}

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

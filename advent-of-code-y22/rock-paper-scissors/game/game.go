// Package game for Rock-Paper-Scissors game: https://en.wikipedia.org/wiki/Rock_paper_scissors
// With basic rules: Rock -> beat -> Scissors -> beat -> Paper -> beat -> Rock ...
package game

type Move int

const (
	Rock Move = iota
	Paper
	Scissors
)

type Outcome int

const (
	Draw Outcome = iota
	Win
	Loss
)

// Play function receives player1 and player 2 Move and return
func Play(player1 Move, player2 Move) (Outcome, Outcome) {
	diff := player1 - player2

	if diff == -1 || diff == 2 {
		return Loss, Win
	}
	if diff == 1 || diff == -2 {
		return Win, Loss
	}
	return Draw, Draw
}

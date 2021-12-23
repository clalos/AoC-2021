package main

import "fmt"

func main() {
	p1StartingPosition := 1
	p2StartingPosition := 5

	// Run game.
	g := &game{
		player1Position: p1StartingPosition,
		player2Position: p2StartingPosition,
		player1Score:    0,
		player2Score:    0,
		isPlayer1Turn:   true,
	}
	result := playGame(g)

	fmt.Println("P1 wins:", result.player1Wins)
	fmt.Println("P2 wins:", result.player2Wins)
}

// game holds a single game counters and positions.
type game struct {
	player1Position int
	player2Position int
	player1Score    int
	player2Score    int
	isPlayer1Turn   bool
}

// player1Wins returns true if player 1 won.
func (g *game) player1Wins() bool {
	return g.player1Score >= 21
}

// player2Wins returns true if player 2 won.
func (g *game) player2Wins() bool {
	return g.player2Score >= 21
}

// playTurn plays a single turn.
// diceSum is the sum of 3 rolls, e.g: 111 = 3, 213 = 6.
// It must not alter the current game state.
func (g *game) playTurn(diceSum int) *game {
	if g.player1Wins() || g.player2Wins() {
		return g
	}

	if g.isPlayer1Turn {
		newPosition := (((g.player1Position + diceSum) - 1) % 10) + 1
		newScore := g.player1Score + newPosition
		return &game{
			player1Position: newPosition,
			player2Position: g.player2Position,
			player1Score:    newScore,
			player2Score:    g.player2Score,
			isPlayer1Turn:   false,
		}
	}
	newPosition := (((g.player2Position + diceSum) - 1) % 10) + 1
	newScore := g.player2Score + newPosition
	return &game{
		player1Position: g.player1Position,
		player2Position: newPosition,
		player1Score:    g.player1Score,
		player2Score:    newScore,
		isPlayer1Turn:   true,
	}
}

// result holds the total wins for each player.
type result struct {
	player1Wins uint
	player2Wins uint
}

// add sums the s wins to the r wins.
func (s *result) add(r *result) *result {
	return &result{s.player1Wins + r.player1Wins, s.player2Wins + r.player2Wins}
}

// times multiplies the wins by the factor.
func (s *result) times(factor uint) *result {
	return &result{s.player1Wins * factor, s.player2Wins * factor}
}

// playGame plays the game in all universes and return the total wins.
func playGame(g *game) *result {
	if g.player1Wins() {
		return &result{1, 0}
	}

	if g.player2Wins() {
		return &result{0, 1}
	}

	// Run the game for all possible dice sums.
	// The result is multiplied  by the time that sum occurs.
	// For example: the sum 4 occurs with 3 combinations: 112,121,211
	return  playGame(g.playTurn(3)).
		add(playGame(g.playTurn(4)).times(3)).
		add(playGame(g.playTurn(5)).times(6)).
		add(playGame(g.playTurn(6)).times(7)).
		add(playGame(g.playTurn(7)).times(6)).
		add(playGame(g.playTurn(8)).times(3)).
		add(playGame(g.playTurn(9)));
}

package tennis

import "testing"

func TestNewGameShouldReturnLoveAll(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(0)
	game.SetScorePlayerTwo(0)

	score := game.GetScore()

	if score != "Love all" {
		t.Errorf("Error")
	}
}

func TestPlayerOneWinsFirstBall(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(1)
	game.SetScorePlayerTwo(0)

	score := game.GetScore()

	if score != "Fifteen,Love" {
		t.Errorf("Expect 'Fifteen,Love' but got '%s'", score)
	}
}

func TestFifteenAll(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(1)
	game.SetScorePlayerTwo(1)

	score := game.GetScore()

	if score != "Fifteen all" {
		t.Errorf("Expect 'Fifteen all' but got '%s'", score)
	}
}

func TestPlayerOneWinsFirstOneBalls(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(0)
	game.SetScorePlayerTwo(1)

	score := game.GetScore()

	if score != "Love,Fifteen" {
		t.Errorf("Expect 'Love,Fifteen' but got '%s'", score)
	}
}

func TestPlayerTwoWinsFirstTwoBalls(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(0)
	game.SetScorePlayerTwo(2)

	score := game.GetScore()

	if score != "Love,Thirty" {
		t.Errorf("Expect 'Love,Thirty' but got '%s'", score)
	}
}

func TestPlayerOneWinsFirstThreeBalls(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(3)
	game.SetScorePlayerTwo(0)

	score := game.GetScore()

	if score != "Forty,Love" {
		t.Errorf("Expect 'Forty,Love' but got '%s'", score)
	}
}

func TestPlayersAreDeuce(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(3)
	game.SetScorePlayerTwo(3)

	score := game.GetScore()

	if score != "Deuce" {
		t.Errorf("Expect 'Deuce' but got '%s'", score)
	}
}

func TestPlayerOneWinsGame(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(4)
	game.SetScorePlayerTwo(0)

	score := game.GetScore()

	if score != "May wins" {
		t.Errorf("Expect 'May wins' but got '%s'", score)
	}
}

func TestPlayerTwoWinsGame(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(1)
	game.SetScorePlayerTwo(4)

	score := game.GetScore()

	if score != "Ying wins" {
		t.Errorf("Expect 'Ying wins' but got '%s'", score)
	}
}

func TestPlayersAreDuce4(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(4)
	game.SetScorePlayerTwo(4)

	score := game.GetScore()

	if score != "Deuce" {
		t.Errorf("Expect 'Deuce' but got '%s'", score)
	}
}

func TestPlayerTwoAdvantage(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(4)
	game.SetScorePlayerTwo(5)

	score := game.GetScore()

	if score != "Advantage Ying" {
		t.Errorf("Expect 'Advantage Ying' but got '%s'", score)
	}
}

func TestPlayerOneAdvantage(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(5)
	game.SetScorePlayerTwo(4)

	score := game.GetScore()

	if score != "Advantage May" {
		t.Errorf("Expect 'Advantage May' but got '%s'", score)
	}
}

func TestPlayerTwoWins(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(2)
	game.SetScorePlayerTwo(4)

	score := game.GetScore()

	if score != "Ying wins" {
		t.Errorf("Expect 'Ying wins' but got '%s'", score)
	}
}

func TestPlayerTwoWinsAfterAdvantage(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(6)
	game.SetScorePlayerTwo(8)

	score := game.GetScore()

	if score != "Ying wins" {
		t.Errorf("Expect 'Ying wins' but got '%s'", score)
	}
}

func TestPlayerOneWinsAfterAdvantage(t *testing.T) {
	game := newGame()
	game.SetScorePlayerOne(8)
	game.SetScorePlayerTwo(6)

	score := game.GetScore()

	if score != "May wins" {
		t.Errorf("Expect 'May wins' but got '%s'", score)
	}
}

func newGame() *Game {
	return &Game{
		PlayerOneName: "May",
		PlayerTwoName: "Ying",
	}
}

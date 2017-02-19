package tennis

type Game struct {
	PlayerOneScore int
	PlayerTwoScore int
	PlayerOneName  string
	PlayerTwoName  string
}

func (g *Game) SetScorePlayerOne(score int) {
	g.PlayerOneScore += score
}

func (g *Game) SetScorePlayerTwo(score int) {
	g.PlayerTwoScore += score
}

func (g *Game) GetScore() string {
	if g.PlayerOneScore > 2 && g.PlayerTwoScore > 2 {
		if g.PlayerOneScore == g.PlayerTwoScore {
			return "Deuce"
		}

		if g.PlayerOneScore-g.PlayerTwoScore == 1 {
			return "Advantage " + g.PlayerOneName
		}

		if g.PlayerTwoScore-g.PlayerOneScore == 1 {
			return "Advantage " + g.PlayerTwoName
		}

		if g.PlayerOneScore-g.PlayerTwoScore == 2 {
			return g.PlayerOneName + " wins"
		}

		if g.PlayerTwoScore-g.PlayerOneScore == 2 {
			return g.PlayerTwoName + " wins"
		}
	}

	if g.PlayerOneScore == g.PlayerTwoScore {
		return mapsScore(g.PlayerOneScore) + " all"
	}

	if g.PlayerOneScore == 4 {
		return g.PlayerOneName + " wins"
	}

	if g.PlayerTwoScore == 4 {
		return g.PlayerTwoName + " wins"
	}

	return mapsScore(g.PlayerOneScore) + "," + mapsScore(g.PlayerTwoScore)
}

func mapsScore(scoreDiff int) string {
	switch scoreDiff {
	case 0:
		return "Love"
	case 1:
		return "Fifteen"
	case 2:
		return "Thirty"
	case 3:
		return "Forty"
	default:
		return ""
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

package tennis

import (
	"fmt"
	"log"
)

type Game struct {
	score1 int
	score2 int
}

func (g *Game) Player1Win() {
	g.score1++
}

func (g *Game) Player2Win() {
	g.score2++
}

func (g Game) Score() string {
	log.Println(g.score2 - g.score1)
	score := []string{"Love", "15", "30", "40"}
	if g.score1 >= 4 || g.score2 >= 4 {
		if g.score2-g.score1 == 2 {
			return "Player 2 Win"
		}
		if g.score1-g.score2 == 2 {
			return "Player 1 Win"
		}
		return fmt.Sprintf("%s - %s", score[g.score1], score[g.score2])
	}

	return fmt.Sprintf("%s - %s", score[g.score1], score[g.score2])
}

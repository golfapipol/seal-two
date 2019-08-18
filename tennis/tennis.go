package tennis

import (
	"fmt"
)

func Tennis(score1, score2 int) string {
	if score2-score1 == 2 {
		return "Player 2 Win"
	}
	if score1-score2 == 2 {
		return "Player 1 Win"
	}
	score := []string{"Love", "15", "30", "40"}
	return fmt.Sprintf("%s - %s", score[score1], score[score2])
}

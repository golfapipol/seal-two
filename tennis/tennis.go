package tennis

import (
	"fmt"
)

func Tennis(score1, score2 int) string {
	score := []string{"Love", "15", "30", "40"}
	if score1 >= 4 || score2 >= 4 {
		if score2-score1 == 2 {
			return "Player 2 Win"
		}
		if score1-score2 == 2 {
			return "Player 1 Win"
		}
		return fmt.Sprintf("%s - %s", score[score1], score[score2])
	}
	return fmt.Sprintf("%s - %s", score[score1], score[score2])
}

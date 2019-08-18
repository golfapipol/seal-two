package tennis_test

import (
	. "tennis"
	"testing"
)

func Test_TennisGame_Input_1_0_Should_Be_15_Love(t *testing.T) {
	expected := "15 - Love"

	game := Game{}
	game.Player1Win()
	actual := game.Score()

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_TennisGame_Input_1_1_Should_Be_15_15(t *testing.T) {
	expected := "15 - 15"

	game := Game{}
	game.Player1Win()
	game.Player2Win()
	actual := game.Score()

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_TennisGame_Input_1_2_Should_Be_15_30(t *testing.T) {
	expected := "15 - 30"

	game := Game{}
	game.Player1Win()
	game.Player2Win()
	game.Player2Win()
	actual := game.Score()

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_TennisGame_Input_1_3_Should_Be_15_40(t *testing.T) {
	expected := "15 - 40"

	game := Game{}
	game.Player1Win()
	game.Player2Win()
	game.Player2Win()
	game.Player2Win()
	actual := game.Score()

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_TennisGame_Input_2_3_Should_Be_30_40(t *testing.T) {
	expected := "30 - 40"

	game := Game{}
	game.Player1Win()
	game.Player2Win()
	game.Player2Win()
	game.Player1Win()
	game.Player2Win()
	actual := game.Score()

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_TennisGame_Input_2_4_Should_Be_Player2_Win(t *testing.T) {
	expected := "Player 2 Win"

	game := Game{}
	game.Player1Win()
	game.Player2Win()
	game.Player2Win()
	game.Player1Win()
	game.Player2Win()
	game.Player2Win()
	actual := game.Score()

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_TennisGame_Input_8_6_Should_Be_Player1_Win(t *testing.T) {
	expected := "Player 1 Win"

	game := Game{}
	game.Player1Win()
	game.Player2Win()
	game.Player2Win()
	game.Player1Win()
	game.Player2Win()
	game.Player2Win()
	game.Player1Win()
	game.Player1Win()
	game.Player2Win()
	game.Player1Win()
	game.Player2Win()
	game.Player1Win()
	game.Player1Win()
	game.Player1Win()
	actual := game.Score()

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

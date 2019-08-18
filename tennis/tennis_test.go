package tennis_test

import (
	. "tennis"
	"testing"
)

func Test_Tennis_Input_1_0_Should_Be_15_Love(t *testing.T) {
	expected := "15 - Love"

	actual := Tennis(1, 0)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_Tennis_Input_1_1_Should_Be_15_15(t *testing.T) {
	expected := "15 - 15"

	actual := Tennis(1, 1)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_Tennis_Input_1_2_Should_Be_15_30(t *testing.T) {
	expected := "15 - 30"

	actual := Tennis(1, 2)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_Tennis_Input_1_3_Should_Be_15_40(t *testing.T) {
	expected := "15 - 40"

	actual := Tennis(1, 3)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_Tennis_Input_2_3_Should_Be_30_40(t *testing.T) {
	expected := "30 - 40"

	actual := Tennis(2, 3)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_Tennis_Input_2_4_Should_Be_Player2_Win(t *testing.T) {
	expected := "Player 2 Win"

	actual := Tennis(2, 4)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_Tennis_Input_8_6_Should_Be_Player1_Win(t *testing.T) {
	expected := "Player 1 Win"

	actual := Tennis(8, 6)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

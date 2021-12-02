package main

import (
	"testing"
)

func TestMoveForward(t *testing.T) {
	want := position{15, 60, 10}
	got := position{10, 10, 10}
	got.MoveForward(5)

	if (got.x != want.x) ||
		(got.y != want.y) ||
		(got.aim != want.aim) {
		t.Errorf(
			"MoveForward expected to move to position %d, %d, %d "+
				"- got %d, %d, %d \n",
			want.x, want.y, want.aim, got.x, got.y, got.aim)
	}
}

func TestMoveUp(t *testing.T) {
	want := position{10, 10, 5}
	got := position{10, 10, 10}
	got.MoveUp(5)

	if (got.x != want.x) ||
		(got.y != want.y) ||
		(got.aim != want.aim) {
		t.Errorf(
			"MoveUp expected to move to position %d, %d, %d "+
				"- got %d, %d, %d \n",
			want.x, want.y, want.aim, got.x, got.y, got.aim)
	}

}

func TestMoveDown(t *testing.T) {
	want := position{10, 10, 15}
	got := position{10, 10, 10}
	got.MoveDown(5)

	if (got.x != want.x) ||
		(got.y != want.y) ||
		(got.aim != want.aim) {
		t.Errorf(
			"MoveDown expected to move to position %d, %d, %d "+
				"- got %d, %d, %d \n",
			want.x, want.y, want.aim, got.x, got.y, got.aim)
	}

}

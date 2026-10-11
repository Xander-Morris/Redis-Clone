package commands

import (
	"fmt"
	"testing"
)

func TestSetThenGet(t *testing.T) {
	val, err := ExecuteCommand([]string{"SET", "xander", "5"})

	if val != nil || err != nil {
		t.Errorf("Valid SET should return nil, nil")
	}

	val, err = ExecuteCommand([]string{"GET", "xander"})

	fmt.Println(val, err)

	if val != "5" || err != nil {
		t.Errorf("GET should have returned 5, but received %s and %s", val, err)
	}
}

func TestSetThenIncrement(t *testing.T) {
	val, err := ExecuteCommand([]string{"SET", "xander", "5"})

	if val != nil || err != nil {
		t.Errorf("Valid SET should return nil, nil")
	}

	val, err = ExecuteCommand([]string{"INCR", "xander"})

	fmt.Println(val, err)

	if val != 6 || err != nil {
		t.Errorf("INCR should have returned 6, but received %s and %s", val, err)
	}
}

func TestExists(t *testing.T) {
	val, err := ExecuteCommand([]string{"EXISTS", "xander"})

	if val != false || err != nil {
		t.Errorf("Should have returned false!")
	}

	ExecuteCommand([]string{"SET", "xander", "5"})
	val, err = ExecuteCommand([]string{"EXISTS", "xander"})

	if val != true || err != nil {
		t.Errorf("Should have returned true!");
	}
}
package commands

import (
	"fmt"
	"testing"
)

func TestExecuteCommand(t *testing.T) {
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

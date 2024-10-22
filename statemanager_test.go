package statemanager_test

import (
	"os"
	"testing"

	"github.com/alessandrobessi/statemanager"
)

type TestState struct {
	Count int `json:"count"`
}

func TestStateManager(t *testing.T) {
	// Define the initial state
	initialState := TestState{Count: 0}
	sm := statemanager.NewStateManager[TestState]("test_state.json", initialState)

	// Update and save state
	sm.State.Count++
	err := sm.SaveState()
	if err != nil {
		t.Fatalf("failed to save state: %v", err)
	}

	// Reload state and verify
	smReloaded := statemanager.NewStateManager[TestState]("test_state.json", initialState)
	if smReloaded.State.Count != 1 {
		t.Fatalf("expected count 1, got %d", smReloaded.State.Count)
	}

	// Clean up
	os.Remove("test_state.json")
}

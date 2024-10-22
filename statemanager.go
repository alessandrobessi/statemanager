package statemanager

import (
	"encoding/json"
	"fmt"
	"os"
)

// StateManager structure to encapsulate state management
type StateManager struct {
	FilePath string
	State    interface{}
}

// NewStateManager creates a new StateManager instance
func NewStateManager(filePath string, initialState interface{}) *StateManager {
	sm := &StateManager{
		FilePath: filePath,
		State:    initialState,
	}

	// Try loading the state from file
	err := sm.LoadState()
	if err != nil {
		fmt.Println("Starting with initial state:", initialState)
	}

	return sm
}

// LoadState reads the state from a file and updates the State field
func (sm *StateManager) LoadState() error {
	file, err := os.ReadFile(sm.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File doesn't exist, start with initial state
		}
		return err // Other errors (e.g., permission issues)
	}

	// Unmarshal the JSON into the State
	err = json.Unmarshal(file, sm.State)
	if err != nil {
		return fmt.Errorf("error parsing state file: %w", err)
	}

	return nil
}

// SaveState writes the current state to the file
func (sm *StateManager) SaveState() error {
	data, err := json.MarshalIndent(sm.State, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling state: %w", err)
	}

	err = os.WriteFile(sm.FilePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing state file: %w", err)
	}

	return nil
}

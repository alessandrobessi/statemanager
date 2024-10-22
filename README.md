# statemanager
**Lightweight, Flexible State Management for Go Applications!** 🚀

This is a Go module that provides easy file-based state management with support for user-defined state structures. It’s designed for applications that need to persist state across multiple invocations by writing the state to a file.

With this module, users can define their own state struct, and the state manager will handle loading and saving the state to a file in JSON format.

 ## Features
    - Generic state management: Define your own state structure using Go's generics.
    - File-based persistence: State is saved to and loaded from a file in JSON format.
    - Lightweight: No external dependencies are required.

## Installation

To install the module, use the go get command:
```bash

go get github.com/alessandrobessi/statemanager
```
Then, import it in your Go code:

```go

import "github.com/alessandrobessi/statemanager"
```

## Quick Start
1. Define Your State Struct

You can define your own state structure based on your application’s needs.

```go

type MyState struct {
    Count int    `json:"count"`
    Name  string `json:"name"`
}
```
2. Create and Use the State Manager

You can then create an instance of the StateManager, providing the path to the state file and an initial state. The module will load the state from the file or initialize it with the provided state if the file doesn’t exist.

```go

package main

import (
    "fmt"
    "github.com/alessandrobessi/statemanager"
)

type MyState struct {
    Count int    `json:"count"`
    Name  string `json:"name"`
}

func main() {
    initialState := MyState{Count: 0, Name: "My App"}
    sm := statemanager.NewStateManager("state.json", initialState)

    // Access and update the state
    myState := sm.State
    myState.Count++
    fmt.Printf("Name: %s, Invocation Count: %d\n", myState.Name, myState.Count)

    // Save the updated state
    if err := sm.SaveState(); err != nil {
        fmt.Println("Error saving state:", err)
    }
}
```

3. Running the Program

Each time the program is invoked, it will:

    - Load the state from state.json.
    - Increment the Count field.
    - Save the updated state back to state.json.

4. Customizing the State Structure

The StateManager uses Go generics, so you can define any struct to hold your state and the module will work seamlessly with it.

```go

type AnotherState struct {
    LastRun string `json:"last_run"`
    Count   int    `json:"count"`
}
```
You can create a new StateManager with this state structure just as easily:

```go

sm := statemanager.NewStateManager("another_state.json", AnotherState{Count: 0, LastRun: "never"})
```

## API Reference
- `NewStateManager(filePath string, initialState T)`

Creates a new StateManager instance.

    - filePath (string): The path to the file where the state is stored.
    - initialState (T): The initial state of your application. This can be any struct.

- `LoadState() error`

Loads the state from the file. If the file doesn’t exist, it will continue using the provided initial state.

- `SaveState() error`

Saves the current state to the file. The state is serialized into JSON format.

## Testing

To run the tests for the module, use the following command:

```bash

go test ./...
```

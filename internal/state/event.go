package state

import (
	"fmt"
	"sync"
)

// GameState represents the state of a game for a single user.
type GameState struct {
	UserID     int64
	Score      int
	Level      int
	IsGameOver bool
}

// NewGameState initializes a new GameState with default values for a user.
func NewGameState(userID int64) *GameState {
	return &GameState{
		UserID:     userID,
		Score:      0,
		Level:      1,
		IsGameOver: false,
	}
}

// AddScore increases the score by a given value.
func (gs *GameState) AddScore(points int) {
	if !gs.IsGameOver {
		gs.Score += points
		fmt.Printf("User %d: Added %d points, new score: %d\n", gs.UserID, points, gs.Score)
	} else {
		fmt.Printf("User %d: Cannot add score, the game is over.\n", gs.UserID)
	}
}

// AdvanceLevel increments the level of the game.
func (gs *GameState) AdvanceLevel() {
	if !gs.IsGameOver {
		gs.Level++
		fmt.Printf("User %d: Advanced to level %d\n", gs.UserID, gs.Level)
	} else {
		fmt.Printf("User %d: Cannot advance level, the game is over.\n", gs.UserID)
	}
}

// EndGame sets the game state to over.
func (gs *GameState) EndGame() {
	gs.IsGameOver = true
	fmt.Printf("User %d: Game Over!\n", gs.UserID)
}

// GetState returns the current state as a formatted string.
func (gs *GameState) GetState() string {
	return fmt.Sprintf("User %d - Score: %d, Level: %d, Game Over: %v",
		gs.UserID, gs.Score, gs.Level, gs.IsGameOver)
}

// GameManager manages the game states for multiple users.
type GameManager struct {
	mu     sync.RWMutex
	states map[int64]*GameState
}

// NewGameManager initializes a new GameManager.
func NewGameManager() *GameManager {
	return &GameManager{
		states: make(map[int64]*GameState),
	}
}

// GetOrCreateGameState retrieves an existing GameState or creates a new one for the user.
func (gm *GameManager) GetOrCreateGameState(userID int64) *GameState {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	// Check if the game state exists for this user
	if state, exists := gm.states[userID]; exists {
		return state
	}

	// Create a new game state for the user if not found
	newState := NewGameState(userID)
	gm.states[userID] = newState
	return newState
}

// DeleteGameState removes the game state for a user.
func (gm *GameManager) DeleteGameState(userID int64) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	delete(gm.states, userID)
}

// GetAllStates returns the state of all users (useful for debugging or admin features).
func (gm *GameManager) GetAllStates() string {
	gm.mu.RLock()
	defer gm.mu.RUnlock()

	var result string
	for userID, state := range gm.states {
		result += fmt.Sprintf("User %d: %s\n", userID, state.GetState())
	}
	return result
}

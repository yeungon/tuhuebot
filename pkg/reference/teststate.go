package reference

import (
	"fmt"

	"github.com/yeungon/tuhuebot/internal/state"
)

func Test() { // Create a new GameManager to handle states for all users.
	gm := state.NewGameManager()

	// Simulate interactions with two users.
	user1ID := int64(12345)
	user2ID := int64(67890)

	// User 1's game actions.
	user1State := gm.GetOrCreateGameState(user1ID)
	user1State.AddScore(10)
	user1State.AdvanceLevel()

	// User 2's game actions.
	user2State := gm.GetOrCreateGameState(user2ID)
	user2State.AddScore(5)
	user2State.AdvanceLevel()

	// User 1's game actions again.
	user1State.AddScore(15)
	user1State.EndGame()

	// Display the states of both users.
	fmt.Println("User 1 State:", user1State.GetState())
	fmt.Println("User 2 State:", user2State.GetState())

	// Clean up user1's state.
	gm.DeleteGameState(user1ID)

	// Display all remaining states.
	fmt.Println("All States:", gm.GetAllStates())
}

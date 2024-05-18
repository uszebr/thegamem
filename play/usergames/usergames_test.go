package usergames

import (
	"testing"

	"github.com/uszebr/thegamem/play/game"
)

func TestGetUserGames(t *testing.T) {
	// Test that GetUserGames returns a singleton instance
	ug1 := GetUserGames()
	ug2 := GetUserGames()
	if ug1 != ug2 {
		t.Errorf("Expected same instance, got different instances")
	}
}

func TestAddGameForUser(t *testing.T) {
	// Initialize UserGames instance
	userGames := GetUserGames()

	// Create a mock game
	mockGame := &game.Game{}

	// Add game for user
	userID := "user123"
	userGames.AddGameForUser(userID, mockGame)

	// Verify that the game was added
	retrievedGame, ok := userGames.GetGameForUser(userID)
	if !ok {
		t.Errorf("Expected to find game for user %s, but did not", userID)
	}
	if retrievedGame != mockGame {
		t.Errorf("Expected game %v, got %v", mockGame, retrievedGame)
	}
}

func TestGetGameForUser(t *testing.T) {
	// Initialize UserGames instance
	userGames := GetUserGames()

	// Test retrieving a game for a non-existent user
	nonExistentUserID := "nonExistentUser"
	_, ok := userGames.GetGameForUser(nonExistentUserID)
	if ok {
		t.Errorf("Expected not to find game for user %s, but found one", nonExistentUserID)
	}

	// Create and add a mock game
	mockGame := &game.Game{}
	userID := "user123"
	userGames.AddGameForUser(userID, mockGame)

	// Test retrieving the game for an existing user
	retrievedGame, ok := userGames.GetGameForUser(userID)
	if !ok {
		t.Errorf("Expected to find game for user %s, but did not", userID)
	}
	if retrievedGame != mockGame {
		t.Errorf("Expected game %v, got %v", mockGame, retrievedGame)
	}
}

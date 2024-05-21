package rounder

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uszebr/thegamem/play/model/modelfactory"
	"github.com/uszebr/thegamem/play/player"
)

func TestNewRounder(t *testing.T) {
	mockPlayer := player.New(modelfactory.GetModelFactory().MustCreateModel("alwaysgreen"))
	roundUUID := uuid.New()
	rounder := New(LEFT, mockPlayer, roundUUID)
	assert.Equal(t, string(LEFT), rounder.GetLocation(), "Expected location to be LEFT")
	assert.Equal(t, mockPlayer, rounder.player, "Expected player to be mockPlayer")
	assert.Equal(t, roundUUID, rounder.roundUUID, "Expected roundUUID to match")
	assert.Empty(t, rounder.Signals, "Expected Signals to be empty")
	assert.Empty(t, rounder.Scores, "Expected Scores to be empty")
	assert.Empty(t, rounder.ScoreSums, "Expected ScoreSums to be empty")
	assert.Equal(t, 0, rounder.RoundScoreSum, "Expected RoundScoreSum to be 0")
}

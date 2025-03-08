package gamekit

import (
	"testing"

	"github.com/progrium/darwinkit/internal/assert"
)
	
func TestGameKitValid(t *testing.T) {
	// Test that the classes can be accessed
	assert.NotNil(t, GameCenterViewControllerClass)
	assert.NotNil(t, PlayerClass)
	assert.NotNil(t, LeaderboardClass)
	assert.NotNil(t, AchievementClass)
	assert.NotNil(t, AchievementDescriptionClass)
	assert.NotNil(t, ScoreClass)
	assert.NotNil(t, MatchClass)
	assert.NotNil(t, MatchRequestClass)

	// Test the constants
	assert.Equal(t, ViewStateDefault, 0)
	assert.Equal(t, ViewStateLeaderboards, 1)
	assert.Equal(t, ViewStateAchievements, 2)
	assert.Equal(t, TimeScopeToday, 0)
	assert.Equal(t, TimeScopeWeek, 1)
	assert.Equal(t, TimeScopeAllTime, 2)
	assert.Equal(t, PlayerScopeGlobal, 0)
	assert.Equal(t, PlayerScopeFriendsOnly, 1)
	assert.Equal(t, DataModeReliable, 0)
	assert.Equal(t, DataModeUnreliable, 1)
}
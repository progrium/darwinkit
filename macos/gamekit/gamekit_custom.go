package gamekit

import (
	"github.com/progrium/darwinkit/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// GameCenterViewController provides a view controller for Game Center features
type GameCenterViewController struct {
	objc.Object
}

// GameCenterViewControllerClass is the class instance for GameCenterViewController
var GameCenterViewControllerClass = objc.GetClass("GKGameCenterViewController")

// NewGameCenterViewController creates a new game center view controller
func NewGameCenterViewController() GameCenterViewController {
	return GameCenterViewController{objc.Call[objc.Object](GameCenterViewControllerClass, objc.Sel("alloc")).Send(objc.Sel("init"))}
}

// SetViewState sets the view state of the game center view controller
func (g GameCenterViewController) SetViewState(viewState int) {
	g.Send(objc.Sel("setViewState:"), viewState)
}

// SetGameCenterDelegate sets the delegate for the game center view controller
func (g GameCenterViewController) SetGameCenterDelegate(delegate objc.Object) {
	g.Send(objc.Sel("setGameCenterDelegate:"), delegate)
}

// Local player methods
func LocalPlayer() Player {
	return Player{objc.Call[objc.Object](PlayerClass, objc.Sel("localPlayer"))}
}

// Player represents a Game Center player
type Player struct {
	objc.Object
}

// PlayerClass is the class instance for Player
var PlayerClass = objc.GetClass("GKPlayer")

// PlayerID returns the player's ID
func (p Player) PlayerID() foundation.String {
	return objc.Call[foundation.String](p, objc.Sel("playerID"))
}

// DisplayName returns the player's display name
func (p Player) DisplayName() foundation.String {
	return objc.Call[foundation.String](p, objc.Sel("displayName"))
}

// Alias returns the player's alias
func (p Player) Alias() foundation.String {
	return objc.Call[foundation.String](p, objc.Sel("alias"))
}

// IsAuthenticated returns whether the player is authenticated
func (p Player) IsAuthenticated() bool {
	return objc.Call[bool](p, objc.Sel("isAuthenticated"))
}

// Authenticate authenticates the player
func (p Player) AuthenticateWithCompletionHandler(completion foundation.CompletionHandler) {
	p.Send(objc.Sel("authenticateWithCompletionHandler:"), completion)
}

// Leaderboard represents a Game Center leaderboard
type Leaderboard struct {
	objc.Object
}

// LeaderboardClass is the class instance for Leaderboard
var LeaderboardClass = objc.GetClass("GKLeaderboard")

// NewLeaderboard creates a new leaderboard
func NewLeaderboard() Leaderboard {
	return Leaderboard{objc.Call[objc.Object](LeaderboardClass, objc.Sel("alloc")).Send(objc.Sel("init"))}
}

// SetIdentifier sets the identifier of the leaderboard
func (l Leaderboard) SetIdentifier(identifier foundation.String) {
	l.Send(objc.Sel("setIdentifier:"), identifier)
}

// SetTimeScope sets the time scope of the leaderboard
func (l Leaderboard) SetTimeScope(timeScope int) {
	l.Send(objc.Sel("setTimeScope:"), timeScope)
}

// SetPlayerScope sets the player scope of the leaderboard
func (l Leaderboard) SetPlayerScope(playerScope int) {
	l.Send(objc.Sel("setPlayerScope:"), playerScope)
}

// LoadScoresWithCompletionHandler loads scores for the leaderboard
func (l Leaderboard) LoadScoresWithCompletionHandler(completion foundation.CompletionHandler) {
	l.Send(objc.Sel("loadScoresWithCompletionHandler:"), completion)
}

// Achievement represents a Game Center achievement
type Achievement struct {
	objc.Object
}

// AchievementClass is the class instance for Achievement
var AchievementClass = objc.GetClass("GKAchievement")

// NewAchievementWithIdentifier creates a new achievement with the specified identifier
func NewAchievementWithIdentifier(identifier foundation.String) Achievement {
	return Achievement{objc.Call[objc.Object](AchievementClass, objc.Sel("achievementWithIdentifier:"), identifier)}
}

// SetPercentComplete sets the percent complete for the achievement
func (a Achievement) SetPercentComplete(percentComplete float64) {
	a.Send(objc.Sel("setPercentComplete:"), percentComplete)
}

// SetShowsCompletionBanner sets whether to show a completion banner for the achievement
func (a Achievement) SetShowsCompletionBanner(showsCompletionBanner bool) {
	a.Send(objc.Sel("setShowsCompletionBanner:"), showsCompletionBanner)
}

// ReportAchievements reports achievements to Game Center
func ReportAchievements(achievements foundation.Array, completion foundation.CompletionHandler) {
	objc.Call[objc.Void](AchievementClass, objc.Sel("reportAchievements:withCompletionHandler:"), achievements, completion)
}

// AchievementDescription represents the description of a Game Center achievement
type AchievementDescription struct {
	objc.Object
}

// AchievementDescriptionClass is the class instance for AchievementDescription
var AchievementDescriptionClass = objc.GetClass("GKAchievementDescription")

// LoadAchievementDescriptionsWithCompletionHandler loads achievement descriptions
func LoadAchievementDescriptionsWithCompletionHandler(completion foundation.CompletionHandler) {
	objc.Call[objc.Void](AchievementDescriptionClass, objc.Sel("loadAchievementDescriptionsWithCompletionHandler:"), completion)
}

// Identifier returns the identifier of the achievement description
func (a AchievementDescription) Identifier() foundation.String {
	return objc.Call[foundation.String](a, objc.Sel("identifier"))
}

// Title returns the title of the achievement description
func (a AchievementDescription) Title() foundation.String {
	return objc.Call[foundation.String](a, objc.Sel("title"))
}

// AchievedDescription returns the achieved description of the achievement description
func (a AchievementDescription) AchievedDescription() foundation.String {
	return objc.Call[foundation.String](a, objc.Sel("achievedDescription"))
}

// UnachievedDescription returns the unachieved description of the achievement description
func (a AchievementDescription) UnachievedDescription() foundation.String {
	return objc.Call[foundation.String](a, objc.Sel("unachievedDescription"))
}

// Score represents a Game Center score
type Score struct {
	objc.Object
}

// ScoreClass is the class instance for Score
var ScoreClass = objc.GetClass("GKScore")

// NewScoreWithLeaderboardIdentifier creates a new score with the specified leaderboard identifier
func NewScoreWithLeaderboardIdentifier(identifier foundation.String) Score {
	return Score{objc.Call[objc.Object](ScoreClass, objc.Sel("scoreWithLeaderboardIdentifier:"), identifier)}
}

// SetValue sets the value of the score
func (s Score) SetValue(value int64) {
	s.Send(objc.Sel("setValue:"), value)
}

// ReportScores reports scores to Game Center
func ReportScores(scores foundation.Array, completion foundation.CompletionHandler) {
	objc.Call[objc.Void](ScoreClass, objc.Sel("reportScores:withCompletionHandler:"), scores, completion)
}

// Match represents a Game Center match
type Match struct {
	objc.Object
}

// MatchClass is the class instance for Match
var MatchClass = objc.GetClass("GKMatch")

// SendDataToAllPlayersWithDataMode sends data to all players in the match
func (m Match) SendDataToAllPlayersWithDataMode(data foundation.Data, dataMode int) bool {
	return objc.Call[bool](m, objc.Sel("sendDataToAllPlayers:withDataMode:"), data, dataMode)
}

// SetDelegate sets the delegate for the match
func (m Match) SetDelegate(delegate objc.Object) {
	m.Send(objc.Sel("setDelegate:"), delegate)
}

// Disconnect disconnects from the match
func (m Match) Disconnect() {
	m.Send(objc.Sel("disconnect"))
}

// MatchRequest represents a request for a Game Center match
type MatchRequest struct {
	objc.Object
}

// MatchRequestClass is the class instance for MatchRequest
var MatchRequestClass = objc.GetClass("GKMatchRequest")

// NewMatchRequest creates a new match request
func NewMatchRequest() MatchRequest {
	return MatchRequest{objc.Call[objc.Object](MatchRequestClass, objc.Sel("alloc")).Send(objc.Sel("init"))}
}

// SetMaxPlayers sets the maximum number of players for the match
func (m MatchRequest) SetMaxPlayers(maxPlayers int) {
	m.Send(objc.Sel("setMaxPlayers:"), maxPlayers)
}

// SetMinPlayers sets the minimum number of players for the match
func (m MatchRequest) SetMinPlayers(minPlayers int) {
	m.Send(objc.Sel("setMinPlayers:"), minPlayers)
}

// SetPlayerGroup sets the player group for the match
func (m MatchRequest) SetPlayerGroup(playerGroup int) {
	m.Send(objc.Sel("setPlayerGroup:"), playerGroup)
}

// SetPlayerAttributes sets the player attributes for the match
func (m MatchRequest) SetPlayerAttributes(playerAttributes int) {
	m.Send(objc.Sel("setPlayerAttributes:"), playerAttributes)
}

// Constants for Game Center view controller view states
const (
	ViewStateDefault      = 0
	ViewStateLeaderboards = 1
	ViewStateAchievements = 2
	ViewStateChallenges   = 3
)

// Constants for leaderboard time scopes
const (
	TimeScopeToday     = 0
	TimeScopeWeek      = 1
	TimeScopeAllTime   = 2
)

// Constants for leaderboard player scopes
const (
	PlayerScopeGlobal   = 0
	PlayerScopeFriendsOnly = 1
)

// Constants for match data modes
const (
	DataModeReliable    = 0
	DataModeUnreliable  = 1
)
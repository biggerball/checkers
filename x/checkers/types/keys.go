package types

import "time"

const (
	// ModuleName defines the module name
	ModuleName = "checkers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkers"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SystemInfoKey = "SystemInfo-value-"
)

const (
	GameCreatedEventType      = "new-game-created" // Indicates what event type to listen to
	GameCreatedEventCreator   = "creator"          // Subsidiary information
	GameCreatedEventGameIndex = "game-index"       // What game is relevant
	GameCreatedEventBlack     = "black"            // Is it relevant to me?
	GameCreatedEventRed       = "red"              // Is it relevant to me?
)

const (
	MovePlayedEventType      = "move-played"
	MovePlayedEventBoard     = "board"
	MovePlayedEventCreator   = "creator"
	MovePlayedEventGameIndex = "game-index"
	MovePlayedEventCapturedX = "captured-x"
	MovePlayedEventCapturedY = "captured-y"
	MovePlayedEventWinner    = "winner"
)

const (
	GameRejectedEventType      = "game-rejected"
	GameRejectedEventCreator   = "creator"
	GameRejectedEventGameIndex = "game-index"
)

const (
	NoFifoIndex = "-1"
)

const (
	//MaxTurnDuration = time.Duration(1 * 60 * 1000_000_000)
	MaxTurnDuration = time.Duration(24 * 3_600 * 1000_000_000) // 1 day
	// DeadlineLayout 日期格式
	DeadlineLayout = "2006-01-02 15:04:05.999999999 +0000 UTC"
)

const (
	GameForfeitedEventType      = "game-forfeited"
	GameForfeitedEventGameIndex = "game-index"
	GameForfeitedEventWinner    = "winner"
	GameForfeitedEventBoard     = "board"
	GameCreatedEventWager       = "wager"
	GameCreatedEventDenom       = "denom"
)

const (
	CreateGameGas       = 15000
	PlayMoveGas         = 1000
	RejectGameRefundGas = 14000
)

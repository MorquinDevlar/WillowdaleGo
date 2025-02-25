package events

import (
	"time"

	"github.com/volte6/gomud/internal/configs"
	"github.com/volte6/gomud/internal/items"
)

// Used to apply or remove buffs
type Buff struct {
	UserId        int
	MobInstanceId int
	BuffId        int
}

func (b Buff) Type() string { return `Buff` }

// Used for giving/taking quest progress
type Quest struct {
	UserId     int
	QuestToken string
}

func (q Quest) Type() string { return `Quest` }

// For special room-targetting actions
type RoomAction struct {
	RoomId       int
	SourceUserId int
	SourceMobId  int
	Action       string
	Details      any
	ReadyTurn    uint64
}

func (r RoomAction) Type() string { return `RoomAction` }

// Used for Input from players/mobs
type Input struct {
	UserId        int
	MobInstanceId int
	InputText     string
	ReadyTurn     uint64
	Flags         EventFlag
}

func (i Input) Type() string { return `Input` }

// Messages that are intended to reach all users on the system
type Broadcast struct {
	Text            string
	SkipLineRefresh bool
}

func (b Broadcast) Type() string { return `Broadcast` }

type Message struct {
	UserId          int
	ExcludeUserIds  []int
	RoomId          int
	Text            string
	IsQuiet         bool // whether it can only be heard by superior "hearing"
	IsCommunication bool // If true, this is a communication such as "say" or "emote"
}

func (m Message) Type() string { return `Message` }

// Special commands that only the webclient is equipped to handle
type WebClientCommand struct {
	ConnectionId uint64
	Text         string
}

func (b WebClientCommand) Type() string { return `WebClientCommand` }

// GMCP Commands from clients to server
type GMCPIn struct {
	ConnectionId uint64
	Command      string
	Json         []byte
}

func (b GMCPIn) Type() string { return `GMCPIn` }

// GMCP Commands from server to client
type GMCPOut struct {
	UserId  int
	Payload any
}

func (b GMCPOut) Type() string { return `GMCPOut` }

// Messages that are intended to reach all users on the system
type System struct {
	Command string
	Data    any
}

func (s System) Type() string { return `System` }

// Payloads describing sound/music to play
type MSP struct {
	UserId    int
	SoundType string // SOUND or MUSIC
	SoundFile string
	Volume    int    // 1-100
	Category  string // special category/type for MSP string
}

func (m MSP) Type() string { return `MSP` }

// Fired whenever a mob or player changes rooms
type RoomChange struct {
	UserId        int
	MobInstanceId int
	FromRoomId    int
	ToRoomId      int
}

func (r RoomChange) Type() string { return `RoomChange` }

// Fired every new round
type NewRound struct {
	RoundNumber uint64
	TimeNow     time.Time
	Config      configs.Config
}

func (r NewRound) Type() string { return `NewRound` }

type NewTurn struct {
	TurnNumber uint64
	TimeNow    time.Time
	Config     configs.Config
}

func (r NewTurn) Type() string { return `NewTurn` }

type ItemOwnership struct {
	UserId        int
	MobInstanceId int
	Item          items.Item
	Gained        bool
}

func (r ItemOwnership) Type() string { return `ItemOwnership` }

package model

import (
	"time"
)

type BasePlayerInfo struct {
	Name      string
	GuildCard string
	Level     uint16
	Class     string
}

type QuestRun struct {
	Server              string
	PlayerName          string
	PlayerClass         string
	GuildCard           string
	UserName            string
	AllPlayers          []BasePlayerInfo
	Id                  string
	Difficulty          string
	Episode             uint16
	QuestName           string
	QuestComplete       bool
	QuestStartTime      time.Time
	QuestStartDate      string
	QuestEndTime        time.Time
	QuestDuration       string
	DeathCount          int
	HP                  []uint16
	TP                  []uint16
	PB                  []float32
	Meseta              []uint32
	MesetaCharged       []int
	Room                []uint16
	IllegalShifta       bool
	PbCategory          bool
	ShiftaLvl           []int16
	DebandLvl           []int16
	Invincible          []bool
	Events              []Event
	Monsters            map[int]Monster
	Bosses              map[string]BossData
	MonsterCount        []int
	MonstersKilledCount []int
	MonsterHpPool       []int
	MonstersDead        int
	WeaponsUsed         map[string]string
	EquipmentUsedTime   map[string]map[string]int
	FreezeTraps         []uint16
	FTUsed              uint16
	DTUsed              uint16
	CTUsed              uint16
	TPUsed              uint16
}

type Monster struct {
	Name       string
	Id         uint16
	UnitxtId   uint32
	SpawnTime  time.Time
	KilledTime time.Time
	Alive      bool
	Frame1     bool
}

type BossData struct {
	Name       string
	Id         uint16
	UnitxtId   uint32
	SpawnTime  time.Time
	KilledTime time.Time
	FirstFrame int
	Hp         []int
}

type Event struct {
	Second      int
	Description string
}

type Game struct {
	Id               string
	Player           string
	PlayerNames      []string
	PlayerClasses    []string
	PlayerGcs        []string
	Category         string
	Episode          int
	Quest            string
	QuestAndCategory string
	Time             time.Duration
	Timestamp        time.Time
	FormattedTime    string
	FormattedDate    string
	GameGzip         []byte
	P1Gzip           []byte
	P1HasStats       bool
	P2Gzip           []byte
	P2HasStats       bool
	P3Gzip           []byte
	P3HasStats       bool
	P4Gzip           []byte
	P4HasStats       bool
}

type FormattedPlayerInfo struct {
	Name      string
	GuildCard string
	HasPov    bool
	Class     string
}

type FormattedGame struct {
	Id           string
	Players      []FormattedPlayerInfo
	PbRun        bool
	NumPlayers   int
	Episode      int
	Quest        string
	Time         string
	Date         string
	RelativeDate string
}

type ClientInfo struct {
	VersionMajor int
	VersionMinor int
	VersionPatch int
}

type MessageOfTheDay struct {
	Authorized bool
	Message    string
}

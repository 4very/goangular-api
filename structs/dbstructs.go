package structs

import "time"

type Player struct {
	PUUID     string    `json:"pUUID"`
	PID       int64     `json:"PID"`
	Name      string    `json:"Name"`
	Server    string    `json:"Server"`
	Class     int       `json:"Class"`
	GUUID     string    `json:"GUUID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Guild struct {
	GUUID     string    `json:"GUUID"`
	GID       int64     `json:"GID"`
	Name      string    `json:"Name"`
	Server    string    `json:"Server"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Leaderboard struct {
	LUUID     string    `json:"LUUID"`
	PUUID     string    `json:"PUUID"`
	DPS       float64   `json:"DPS"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

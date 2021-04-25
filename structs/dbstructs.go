package structs

import "time"

type Player struct {
	PUUID     string    `json:"PUUID"`
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

type Report struct {
	RUUID     string    `json:"RUUID"`
	GUUID     string    `json:"GUUID"`
	RID       string    `json:"RID"`
	Name      string    `json:"Name"`
	NumFights int       `json:"NumFights"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Fight struct {
	FUUID     string `json:"FUUID"`
	RUUID     string `json:"RUUID"`
	Fnum      int    `json:"Fnum"`
	Eid       int64  `json:"Eid"`
	ComParsed bool   `json:"ComParsed"`
}

type ComData struct {
	COMUUID string  `json:"COMUUID"`
	FUUID   string  `json:"FUUID"`
	PUUID   string  `json:"pUUID"`
	DPS     float64 `json:"DPS"`
}

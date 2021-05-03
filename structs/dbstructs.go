package structs

import "time"

type Player struct {
	PUUID     string    `json:"PUUID" pg:",pk"`
	PID       int64     `json:"PID"`
	Name      string    `json:"Name"`
	Server    string    `json:"Server"`
	Class     int       `json:"Class"`
	GUUID     string    `json:"GUUID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Guild struct {
	GUUID     string    `json:"GUUID" pg:",pk"`
	GID       int64     `json:"GID"`
	Name      string    `json:"Name"`
	Server    string    `json:"Server"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Report struct {
	RUUID     string    `json:"RUUID" pg:",pk"`
	GUUID     string    `json:"GUUID"`
	RID       string    `json:"RID"`
	Name      string    `json:"Name"`
	NumFights int       `json:"NumFights"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Fight struct {
	FUUID     string  `json:"FUUID" pg:",pk"`
	RUUID     string  `json:"RUUID"`
	Fnum      int     `json:"Fnum"`
	Diff      int     `json:"Diff"`
	Eid       int64   `json:"Eid"`
	ComValid  bool    `json:"ComValid"`
	ComParsed bool    `json:"ComParsed"`
	StartTime float64 `json:"StartTime"`
	EndTime   float64 `json:"EndTime"`
}

type ComData struct {
	COMUUID   string  `json:"COMUUID" pg:",pk"`
	FUUID     string  `json:"FUUID"`
	StartTime float64 `json:"StartTime"`
	EndTime   float64 `json:"EndTime"`
	ComId     int     `json:"ComId"`
}

type ComPlayerData struct {
	COMPUUID string  `json:"COMPUUID" pg:",pk"`
	COMUUID  string  `json:"COMUUID"`
	PUUID    string  `json:"PUUID"`
	DPS      float64 `json:"DPS"`
}

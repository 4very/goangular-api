package controllers

import (
	"time"

	"github.com/google/uuid"
	st "github.com/sommea/goangular-api/structs"
)

func insertPlayer(p st.Player) (bool, error) {

	exists, _ := dbConnect.Model(&p).Where("PID = ?", p.PID).Exists()
	if exists {
		return false, nil
	}

	p.PUUID = uuid.New().String()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	insertError := dbConnect.Insert(&p)
	if insertError != nil {
		return false, insertError
	}
	return true, nil

}

func insertGuild(g st.Guild) (bool, error) {

	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()

	insertError := dbConnect.Insert(&g)
	if insertError != nil {
		return false, insertError
	}
	return true, nil

}

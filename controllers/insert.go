package controllers

import (
	"time"

	"github.com/google/uuid"
	st "github.com/sommea/goangular-api/structs"
)

func insertPlayer(p st.Player) (bool, error) {

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

func insertReport(r st.Report) (bool, error) {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()

	insertError := dbConnect.Insert(&r)
	if insertError != nil {
		return false, insertError
	}
	return true, nil

}

func insertFight(f st.Fight) (bool, error) {

	insertError := dbConnect.Insert(&f)
	if insertError != nil {
		return false, insertError
	}
	return true, nil

}

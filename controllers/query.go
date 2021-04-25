package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	st "github.com/sommea/goangular-api/structs"
	wcl "github.com/sommea/goangular-api/wcl"
)

func GetAllPlayers(c *gin.Context) {
	var players []st.Player
	err := dbConnect.Model(&players).Select()
	if err != nil {
		log.Printf("Error while getting all players, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Player",
		"data":    players,
	})
}

func CreatePlayer(c *gin.Context) {
	var player st.Player
	c.BindJSON(&player)

	if player.PID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please include a player id",
		})
		return
	}

	code, ret := createPlayer(player.PID)
	c.JSON(code, ret)
}

func playerExists(pid int64) bool {
	var players []st.Player
	exists, _ := dbConnect.Model(&players).Where("PID = ?", pid).Exists()
	return exists

}

func createPlayer(pid int64) (int, gin.H) {

	exists := playerExists(pid)
	if exists {
		return http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "Player already exists in database",
		}
	}

	p := wcl.GetUserData(pid)
	ret, insertError := insertPlayer(p)

	if !ret {
		log.Printf("Error while inserting new Player into db, Reason: %v\n", insertError)
		return http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		}
	}

	return http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Player created Successfully",
	}
}

func GetPlayer(c *gin.Context) {
	var player st.Player
	c.BindJSON(&player)

	if player.PID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please include a player id",
		})
		return
	}

	var players []st.Player
	err := dbConnect.Model(&players).Where("PID = ?", player.PID).Select()

	if err != nil {
		log.Printf("Error while getting player %v, Reason: %v\n", player.PID, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get Player",
		"data":    players,
	})
}

func GetAllGuilds(c *gin.Context) {
	var guilds []st.Guild
	err := dbConnect.Model(&guilds).Select()
	if err != nil {
		log.Printf("Error while getting all guilds, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Player",
		"data":    guilds,
	})
}

func guildExists(gid int64) bool {
	var guilds []st.Guild
	exists, _ := dbConnect.Model(&guilds).Where("GID = ?", gid).Exists()
	return exists
}

func getGuild(gid int64) st.Guild {
	var guild st.Guild
	dbConnect.Model(&guild).Where("GID = ?", gid).Select()
	return guild
}

func GetGuild(c *gin.Context) {
	var guild st.Guild
	c.BindJSON(&guild)

	if guild.GID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please include a Guild id",
		})
		return
	}

	var guilds []st.Guild
	err := dbConnect.Model(&guilds).Where("PID = ?", guild.GID).Select()

	if err != nil {
		log.Printf("Error while getting player %v, Reason: %v\n", guild.GID, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Get Player",
		"data":    guilds,
	})
}

func createGuild(gid int64) (int, gin.H) {

	gdata, psdata := wcl.GetGuildData(gid)

	exists := guildExists(gid)
	if exists {
		return http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "Guild already exists in database",
		}
	}

	guuid := uuid.New().String()
	gdata.GUUID = guuid

	ret, insertError := insertGuild(gdata)

	for _, elt := range psdata {
		if playerExists(elt.PID) {
			continue
		}

		elt.GUUID = guuid
		pret, pinsertError := insertPlayer(elt)
		if !pret {
			log.Printf("Error while inserting new Player into db, Reason: %v\n", pinsertError)
			return http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Something went wrong",
			}
		}
	}

	if !ret {
		log.Printf("Error while inserting new Guild into db, Reason: %v\n", insertError)
		return http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		}

	}

	return http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Guild created Successfully",
	}

}

func CreateGuild(c *gin.Context) {
	var guild st.Guild
	c.BindJSON(&guild)

	if guild.GID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please include a guild id",
		})
		return
	}

	code, ret := createGuild(guild.GID)
	c.JSON(code, ret)

}

func createReport(rid string) (int, gin.H) {

	rdata, fsdata, gid := wcl.GetReportData(rid)

	if !guildExists(gid) {
		createGuild(gid)
	}
	guild := getGuild(gid)

	ruuid := uuid.New().String()

	for _, elt := range fsdata {
		elt.RUUID = ruuid
		elt.FUUID = uuid.New().String()
		pret, pinsertError := insertFight(elt)
		if !pret {
			log.Printf("Error while inserting new Player into db, Reason: %v\n", pinsertError)
			return http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Something went wrong",
			}
		}

	}

	rdata.GUUID = guild.GUUID
	rdata.RUUID = ruuid
	ret, insertError := insertReport(rdata)

	if !ret {
		log.Printf("Error while inserting new Guild into db, Reason: %v\n", insertError)
		return http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		}
	}

	return http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Report created Successfully",
	}

}

func CreateReport(c *gin.Context) {
	var report st.Report
	c.BindJSON(&report)

	if report.RID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please include a Report id",
		})
		return
	}

	code, json := createReport(report.RID)
	c.JSON(code, json)

}

func GetAllReports(c *gin.Context) {
	var reports []st.Report
	err := dbConnect.Model(&reports).Select()
	if err != nil {
		log.Printf("Error while getting all reports, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Reports",
		"data":    reports,
	})
}

func GetAllFights(c *gin.Context) {
	var fights []st.Fight
	err := dbConnect.Model(&fights).Select()
	if err != nil {
		log.Printf("Error while getting all Fights, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Fights",
		"data":    fights,
	})
}

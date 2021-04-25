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
		return
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
	var players []st.Player
	exists, _ := dbConnect.Model(&players).Where("PID = ?", player.PID).Exists()

	if exists {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "Player already exists in database",
		})
		return
	}

	p := wcl.GetUserData(player.PID)
	ret, insertError := insertPlayer(p)

	if !ret {
		log.Printf("Error while inserting new Player into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Player created Successfully",
	})
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

func CreateGuild(c *gin.Context) {
	var guild st.Guild
	c.BindJSON(&guild)

	if guild.GID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Please include a player id",
		})
		return
	}
	var players []st.Player
	exists, _ := dbConnect.Model(&players).Where("PID = ?", guild.GID).Exists()

	if exists {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "Player already exists in database",
		})
		return
	}

	gdata, psdata := wcl.GetGuildData(guild.GID)

	gid := uuid.New().String()
	gdata.GUUID = gid

	ret, insertError := insertGuild(gdata)

	for _, elt := range psdata {
		elt.GUUID = gid
		pret, pinsertError := insertPlayer(elt)
		if !pret {
			log.Printf("Error while inserting new Player into db, Reason: %v\n", pinsertError)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Something went wrong",
			})
			return
		}

	}

	if !ret {
		log.Printf("Error while inserting new Guild into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Guild created Successfully",
	})
}

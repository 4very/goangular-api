package wcl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
	st "github.com/sommea/goangular-api/structs"
)

type Token struct {
	Expires int    `json:"expires_in"`
	Token   string `json:"access_token"`
}

func getToken() string {
	params := url.Values{}
	params.Add("grant_type", `client_credentials`)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://www.warcraftlogs.com/oauth/token", body)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	godotenv.Load()
	req.SetBasicAuth(os.Getenv("BASIC_AUTH_USER"), os.Getenv("BASIC_AUTH_PASSWORD"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	defer resp.Body.Close()

	rbody, _ := ioutil.ReadAll(resp.Body)

	rstruct := Token{}
	json.Unmarshal(rbody, &rstruct)
	return rstruct.Token
}

func query(q string, stuct interface{}) bool {

	jsonData := map[string]string{
		"query": q,
	}
	jsonValue, _ := json.Marshal(jsonData)

	req, err := http.NewRequest("GET", "https://www.warcraftlogs.com/api/v2/client", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+getToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	defer resp.Body.Close()

	rbody, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(rbody, &stuct)
	return true
}

type CharDataQuery struct {
	Data struct {
		CharacterData struct {
			Character struct {
				Name    string `Json:"Name"`
				ClassID int    `Json:"classID"`
				Server  struct {
					Name string `Json:"Name"`
				} `Json:"Server"`
			} `Json:"Character"`
		} `Json:"characterData"`
	} `Json:"data"`
}

type CharRet struct {
	PID    int64
	Name   string
	Class  int
	Server string
}

func GetUserData(id int64) st.Player {
	rstruct := CharDataQuery{}

	q := `{
		characterData {
		  character(id: ` + fmt.Sprint(id) + `){
			name,
			classID,
			server{
			  name
			}
		  }
		}
	  }
	`

	query(q, &rstruct)

	var retval st.Player = st.Player{
		PID:    id,
		Name:   rstruct.Data.CharacterData.Character.Name,
		Class:  rstruct.Data.CharacterData.Character.ClassID,
		Server: rstruct.Data.CharacterData.Character.Server.Name,
	}
	return retval
}

type GuildDataQuery struct {
	Data struct {
		GuildData struct {
			Guild struct {
				Name   string `Json:"name"`
				Server struct {
					Name string `Json:"name"`
				} `Json:"Server"`
			} `Json:"guild"`
		} `Json:"guildData"`
	} `Json:"data"`
}

type GuildMemberQuery struct {
	Data struct {
		GuildData struct {
			Guild struct {
				Members struct {
					Data []struct {
						Id      int64  `Json:"id"`
						Name    string `Json:"Name"`
						ClassID int    `Json:"classID"`
						Server  struct {
							Name string `Json:"Name"`
						} `Json:"Server"`
					} `Json:"data"`
				} `Json:"members"`
			} `Json:"guild"`
		} `Json:"guildData"`
	} `Json:"data"`
}

type GuildLastPage struct {
	Data struct {
		GuildData struct {
			Guild struct {
				Members struct {
					LastPage int `Json:"total"`
				} `Json:"members"`
			} `Json:"guild"`
		} `Json:"guildData"`
	} `Json:"data"`
}

type GuildRet struct {
	Name   string
	Server string
	Ids    []int64
}

func GetGuildData(id int64) (st.Guild, []st.Player) {
	var ps []st.Player
	var page int = 1

	for {
		memberq := `{
			guildData{
			  guild(id: ` + fmt.Sprint(id) + `){
				members(limit: 100, page: ` + fmt.Sprint(page) + `){
				  data{
					id
					name
					classID
					server{
					  name
					}
				  }
				}
			  }
			}
		  }`

		rstruct := GuildMemberQuery{}
		query(memberq, &rstruct)

		if len(rstruct.Data.GuildData.Guild.Members.Data) == 0 {
			break
		}

		for _, elt := range rstruct.Data.GuildData.Guild.Members.Data {
			ps = append(ps, st.Player{
				PID:    elt.Id,
				Name:   elt.Name,
				Server: elt.Server.Name,
				Class:  elt.ClassID,
			})
		}

		page++
		fmt.Println(page)

	}

	guildq := `{
		guildData{
		  guild(id: ` + fmt.Sprint(id) + `){
			name
			server{
				name
			}
		  }
		}
	  }
	`
	rstruct := GuildDataQuery{}
	query(guildq, &rstruct)

	gret := st.Guild{
		GID:    id,
		Server: rstruct.Data.GuildData.Guild.Server.Name,
		Name:   rstruct.Data.GuildData.Guild.Name,
	}
	fmt.Println(len(ps))
	return gret, ps
}

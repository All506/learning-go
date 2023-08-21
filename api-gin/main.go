package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"container/list"
)


func main(){
	router := gin.Default()

	//Funcions related with players
	router.GET("/player", getPlayers)
	router.GET("/player/:firstname",getPlayerByName)

	//Functions related with teams
	router.GET("/team/:name",getTeamInfo)
	router.Run("localhost:8080")
}

func getPlayers(c *gin.Context){
	c.IndentedJSON(http.StatusOK, players)
}

func getPlayerByName(c *gin.Context){
	firstname := c.Param("firstname")

	for _, player := range players{
		if player.Firstname == firstname{
			c.IndentedJSON(http.StatusOK, player)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Player has not been found"})
}

func getTeamInfo(c *gin.Context){

	name := c.Param("name")
	exists := false;
	playersInTeam := list.New()
	fmt.Println(name)

	for _, team := range teams{
		if team.Name == name{
			exists = true
		}
	}

	//TODO: Is not filtering information regarding players and their teams
	if exists == true {
		for _, player := range players{
			if player.Team.Name == name {
				playersInTeam.PushFront(player)
				fmt.Println(player)
			}
		}
		c.IndentedJSON(http.StatusOK, playersInTeam)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Team does not exists"})
	}
	
}
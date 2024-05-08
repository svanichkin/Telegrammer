package rest

import (
	"main/conf"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func checkBid(bid string) string {

	if len(bid) == 0 {
		return ""
	}
	for _, bot := range conf.Config.Bots {
		if bot.Bid == bid {
			return bid
		}
	}
	return ""

}

func checkCommandName(bid string, commandName string) string {

	if len(commandName) == 0 {
		return ""
	}
	for _, bot := range conf.Config.Bots {
		if bot.Bid == bid {
			for _, command := range bot.Commands {
				if command.Name == commandName {
					return commandName
				}
			}
		}
	}
	return ""

}

func checkInt(s string) int64 {

	if len(s) == 0 {
		return 0
	}
	u, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return u

}

func checkStrings(s string) []string {

	params := strings.Split(s, ",")
	cleanParams := make([]string, 0, len(params))
	for _, param := range params {
		cleanParam := strings.TrimSpace(param)
		cleanParams = append(cleanParams, cleanParam)
	}
	return cleanParams

}

func checkType(c *gin.Context) string {

	if len(c.Query("type")) == 0 {
		return "json"
	}
	return c.Query("type")

}

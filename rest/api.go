package rest

import (
	"fmt"
	"main/conf"
	"main/data"
	"main/tele"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/telebot.v3"
)

func Init() error {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/send", sendHandler)
	r.POST("/send", sendHandler)

	r.GET("/check", checkHandler)
	r.POST("/check", checkHandler)

	r.GET("/data", dataHandler)
	r.POST("/data", dataHandler)

	r.GET("/command", commandHandler)
	r.POST("/command", commandHandler)

	r.GET("/alias", aliasHandler)
	r.POST("/alias", aliasHandler)

	r.GET("/restart", restartHandler)
	r.POST("/restart", restartHandler)

	r.GET("/access", accessHandler)
	r.POST("/access", accessHandler)

	r.Run(conf.Config.Server.Host + ":" + fmt.Sprint(conf.Config.Server.Port))

	return nil

}

func sendHandler(c *gin.Context) {

	// Check return type

	rt := checkType(c)

	// Validate bid

	bid := checkBid(c.Query("bid"))
	if len(bid) == 0 {
		returnError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Convertation uid to int64

	uid := checkInt(c.Query("uid"))
	aid := c.Query("aid")
	if uid == 0 {
		if len(aid) == 0 {
			returnError(c, fmt.Errorf("bad request: uid is not valid"), rt)
			return
		}
		var err error
		if uid, err = data.FindUidWithAlias(bid, aid); err != nil {
			returnError(c, fmt.Errorf("bad request: uid for alias %s not found: %v", aid, err), rt)
			return
		}
	}

	// Validate chat

	cid := checkInt(c.Query("cid"))

	// Main switch

	switch {
	case c.Query("text") != "":
		if err := tele.SendMessage(bid, uid, cid, c.Query("text")); err != nil {
			returnError(c, err, rt)
			return
		}

	case c.Query("html") != "":
		if err := tele.SendHtml(bid, uid, cid, c.Query("html")); err != nil {
			returnError(c, err, rt)
			return
		}

	case c.Query("markdown") != "":
		if err := tele.SendMarkdown(bid, uid, cid, c.Query("markdown")); err != nil {
			returnError(c, err, rt)
			return
		}

	case c.Query("markdownv2") != "":
		if err := tele.SendMarkdownV2(bid, uid, cid, c.Query("markdownv2")); err != nil {
			returnError(c, err, rt)
			return
		}

	case c.Query("photo") != "":

		// Check photo type [base64, url or local file]

		photo := c.Query("photo")
		if strings.HasPrefix(photo, "data:image/") {
			if err := tele.SendPhotoBase64(bid, uid, cid, photo); err != nil {
				returnError(c, err, rt)
				return
			}
		} else if strings.HasPrefix(photo, "http") {
			if err := tele.SendPhotoUrl(bid, uid, cid, photo); err != nil {
				returnError(c, err, rt)
				return
			}
		} else {
			if err := tele.SendPhotoPath(bid, uid, cid, photo); err != nil {
				returnError(c, err, rt)
				return
			}
		}

	case c.Query("file") != "":
		if err := tele.SendFile(bid, uid, cid, c.Query("file")); err != nil {
			returnError(c, err, rt)
			return
		}

	default:
		returnError(c, fmt.Errorf("bad request: missing command"), rt)
		return
	}

	c.Status(http.StatusOK)

}

func checkHandler(c *gin.Context) {

	// Check return type

	rt := checkType(c)

	// Validate bid

	bid := checkBid(c.Query("bid"))
	if len(bid) == 0 {
		returnError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Convertation uid to int64

	uid := checkInt(c.Query("uid"))
	aid := c.Query("aid")
	if uid == 0 {
		if len(aid) == 0 {
			returnError(c, fmt.Errorf("bad request: uid is not valid"), rt)
			return
		}
		var err error
		if uid, err = data.FindUidWithAlias(bid, aid); err != nil {
			returnError(c, fmt.Errorf("bad request: uid for alias %s not found: %v", aid, err), rt)
			return
		}
	}

	// Validate chat

	cid := checkInt(c.Query("cid"))
	if cid == 0 {
		returnError(c, fmt.Errorf("bad request: cid is not valid"), rt)
		return
	}

	// Check subscribe user

	isSubscribed, err := isSubscribed(bid, uid, cid)
	if err != nil {
		returnError(c, fmt.Errorf("bad check: %v", err), rt)
		return
	}

	if !isSubscribed {
		returnError(c, fmt.Errorf("bad check: user is not subscribed"), rt)
		return
	}

	c.Status(http.StatusOK)

}

func dataHandler(c *gin.Context) {

	// Check return type

	rt := checkType(c)

	// Validate bid

	bid := checkBid(c.Query("bid"))
	if len(bid) == 0 {
		returnError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Convertation uid to int64

	uid := checkInt(c.Query("uid"))
	aid := c.Query("aid")
	if uid == 0 {
		if len(aid) == 0 {
			returnError(c, fmt.Errorf("bad request: uid is not valid"), rt)
			return
		}
		var err error
		if uid, err = data.FindUidWithAlias(bid, aid); err != nil {
			returnError(c, fmt.Errorf("bad request: uid for alias %s not found: %v", aid, err), rt)
			return
		}
	}

	// Validate do

	do := c.Query("do")
	if len(do) == 0 {
		returnError(c, fmt.Errorf("bad request: do is not valid"), rt)
		return
	}

	// Validate key and value or keys and values

	if do == "set" {

		key := c.Query("key")
		value := c.Query("value")
		if len(key) > 0 && len(value) > 0 {
			if err := data.SetKey(bid, uid, key, value); err != nil {
				returnError(c, fmt.Errorf("bad request: %v", err), rt)
			}
			return
		}

		keys := checkStrings(c.Query("keys"))
		values := checkStrings(c.Query("values"))
		if len(keys) > 0 && len(values) > 0 && len(keys) == len(values) {
			for i, key := range keys {
				if err := data.SetKey(bid, uid, key, values[i]); err != nil {
					returnError(c, fmt.Errorf("bad request: %v", err), rt)
				}
			}
			return
		}

		if len(key) == 0 {
			returnError(c, fmt.Errorf("bad request: key is not valid"), rt)
			return
		}
		if len(value) == 0 {
			returnError(c, fmt.Errorf("bad request: value is not valid"), rt)
			return
		}
		if len(keys) == 0 {
			returnError(c, fmt.Errorf("bad request: keys is not valid"), rt)
			return
		}
		if len(values) == 0 {
			returnError(c, fmt.Errorf("bad request: values is not valid"), rt)
			return
		}

		c.Status(http.StatusOK)
		return

	}

	if do == "get" {

		key := c.Query("key")
		if len(key) > 0 {
			value, err := data.GetKey(bid, uid, key)
			if err != nil {
				returnError(c, fmt.Errorf("bad request: %v", err), rt)
				return
			}
			response(c, http.StatusOK, gin.H{key: value}, rt)
			return
		}

		keys := checkStrings(c.Query("keys"))
		if len(keys) > 0 {
			r := gin.H{}
			for _, key := range keys {
				value, err := data.GetKey(bid, uid, key)
				if err != nil {
					returnError(c, fmt.Errorf("bad request: %v", err), rt)
					return
				}
				r[key] = value
			}
			response(c, http.StatusOK, r, rt)
			return
		}

		if len(key) == 0 {
			returnError(c, fmt.Errorf("bad request: key is not valid"), rt)
			return
		}
		if len(keys) == 0 {
			returnError(c, fmt.Errorf("bad request: keys is not valid"), rt)
			return
		}

		c.Status(http.StatusOK)
		return

	}

	if do == "remove" {

		key := c.Query("key")
		if len(key) > 0 {
			err := data.RemoveKey(bid, uid, key)
			if err != nil {
				returnError(c, fmt.Errorf("bad request: %v", err), rt)
				return
			}
			c.Status(http.StatusOK)
			return
		}

		keys := checkStrings(c.Query("keys"))
		if len(keys) > 0 {
			for _, key := range keys {
				err := data.RemoveKey(bid, uid, key)
				if err != nil {
					returnError(c, fmt.Errorf("bad request: %v", err), rt)
					return
				}
			}
			c.Status(http.StatusOK)
			return
		}

		if len(key) == 0 {
			returnError(c, fmt.Errorf("bad request: key is not valid"), rt)
			return
		}
		if len(keys) == 0 {
			returnError(c, fmt.Errorf("bad request: keys is not valid"), rt)
			return
		}

		c.Status(http.StatusOK)
		return

	}

}

func commandHandler(c *gin.Context) {

	// Check return type

	rt := checkType(c)

	// Validate bid

	bid := checkBid(c.Query("bid"))
	if len(bid) == 0 {
		returnError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Convertation uid to int64, uid iz optionally

	uid := checkInt(c.Query("uid"))
	aid := c.Query("aid")
	if uid == 0 && len(aid) > 0 {
		uid, _ = data.FindUidWithAlias(bid, aid)
	}

	// Convertation cid to int64, uid iz optionally

	cid := checkInt(c.Query("cid"))

	// Validate command name

	commandName := checkCommandName(bid, c.Query("name"))
	if len(commandName) == 0 {
		returnError(c, fmt.Errorf("bad request: command name is not valid"), rt)
		return
	}

	// Do action

	err := tele.DoAction(bid, uid, cid, commandName, c.Query("data"), c.Query("firstname"), c.Query("lastname"), c.Query("username"), c.Query("language"))
	if err != nil {
		returnError(c, err, rt)
		return
	}

	c.Status(http.StatusOK)

}

func aliasHandler(c *gin.Context) {

	// Check return type

	rt := checkType(c)

	// Convertation uid to int64

	uid := checkInt(c.Query("uid"))

	// Validate bid

	bid := checkBid(c.Query("bid"))
	if len(bid) == 0 {
		returnError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Validate alias

	aid := c.Query("aid")
	if len(aid) == 0 {
		returnError(c, fmt.Errorf("bad request: alias aid is not valid"), rt)
		return
	}

	// Validate do

	do := c.Query("do")
	if len(do) == 0 {
		returnError(c, fmt.Errorf("bad request: do is not valid"), rt)
		return
	}

	// Validate key and value or keys and values

	if do == "add" {
		if uid == 0 {
			returnError(c, fmt.Errorf("bad request: uid is not valid"), rt)
			return
		}
		if err := data.AddAlias(bid, uid, aid); err != nil {
			returnError(c, fmt.Errorf("bad request: alias %s not added: %v", aid, err), rt)
			return
		}
		c.Status(http.StatusOK)
		return
	}

	if do == "remove" {
		if err := data.RemoveAlias(bid, aid); err != nil {
			returnError(c, fmt.Errorf("bad request: alias %s not added: %v", aid, err), rt)
			return
		}
		c.Status(http.StatusOK)
		return
	}

	if do == "find" {
		uid, err := data.FindUidWithAlias(bid, aid)
		if err != nil {
			returnError(c, fmt.Errorf("bad request: uid for alias %s not found", aid), rt)
			return
		}
		response(c, http.StatusOK, gin.H{"uid": uid}, rt)
		return
	}

}

func restartHandler(c *gin.Context) {

	// Check return type

	rt := checkType(c)

	// Validate bid

	bid := checkBid(c.Query("bid"))
	if len(bid) == 0 {
		returnError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Restart

	if err := conf.RestartBot(bid); err != nil {
		returnError(c, err, rt)
		return
	}
	if err := tele.RestartClient(bid); err != nil {
		returnError(c, err, rt)
		return
	}
	c.Status(http.StatusOK)

}

func accessHandler(c *gin.Context) {

	// Check return type

	rt := checkType(c)

	// Validate bid

	bid := checkBid(c.Query("bid"))
	if len(bid) == 0 {
		returnError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Convertation uid and aid and chat

	uid := checkInt(c.Query("uid"))
	aid := c.Query("aid")
	cid := checkInt(c.Query("cid"))

	if len(aid) == 0 && cid == 0 && uid == 0 {
		returnError(c, fmt.Errorf("bad request: need uid, aid or cid"), rt)
		return
	}

	if len(aid) > 0 {
		var err error
		if uid, err = data.FindUidWithAlias(bid, aid); err != nil {
			returnError(c, fmt.Errorf("bad request: uid for alias %s not found: %v", aid, err), rt)
			return
		}
	}

	// Validate group

	group := c.Query("group")
	if len(group) == 0 {
		returnError(c, fmt.Errorf("bad request: group is not valid"), rt)
		return
	}

	// Validate do

	do := c.Query("do")
	if len(do) == 0 {
		returnError(c, fmt.Errorf("bad request: do is not valid"), rt)
		return
	}

	// Do

	if do == "set" {
		if uid != 0 {
			if err := conf.SetUserForGroup(bid, uid, group); err != nil {
				returnError(c, fmt.Errorf("bad request: user not added to group %s: %v", group, err), rt)
				return
			}
		} else {
			if err := conf.SetChatForGroup(bid, cid, group); err != nil {
				returnError(c, fmt.Errorf("bad request: chat not added to group %s: %v", group, err), rt)
				return
			}
		}
		c.Status(http.StatusOK)
		return
	}

	if do == "get" {
		var groups []conf.AccessGroup
		var err error
		if uid != 0 {
			if groups, err = conf.GetUserForGroup(bid, uid); err != nil {
				returnError(c, fmt.Errorf("bad request: user not found: %v", err), rt)
				return
			}
		} else {
			if groups, err = conf.GetChatForGroup(bid, cid); err != nil {
				returnError(c, fmt.Errorf("bad request: chat not found: %v", err), rt)
				return
			}
		}
		var groupNames []string
		for _, g := range groups {
			groupNames = append(groupNames, g.Name)
		}
		response(c, http.StatusOK, gin.H{"group": groupNames}, rt)
		return
	}

	if do == "remove" {
		if uid != 0 {
			if err := conf.RemoveUserForGroup(bid, uid, group); err != nil {
				returnError(c, fmt.Errorf("bad request: user not removed from group %s: %v", group, err), rt)
				return
			}
		} else {
			if err := conf.RemoveChatForGroup(bid, cid, group); err != nil {
				returnError(c, fmt.Errorf("bad request: chat not removed from group %s: %v", group, err), rt)
				return
			}
		}
		c.Status(http.StatusOK)
		return
	}

}

func formatAsINI(obj map[string]any) string {

	parts := []string{}
	for k, v := range obj {
		parts = append(parts, fmt.Sprintf("%s = %v", k, v))
	}
	return strings.Join(parts, "\n")

}

func returnError(c *gin.Context, err error, rt string) {

	if teleErr, ok := err.(*telebot.Error); ok {
		response(c, teleErr.Code, gin.H{"error": teleErr.Description}, rt)
	} else {
		response(c, http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)}, rt)
	}

}

func response(c *gin.Context, code int, obj map[string]any, t string) {

	switch strings.ToLower(t) {
	case "json":
		c.JSON(code, obj)
	case "jsonp":
		c.JSONP(code, obj)
	case "securejson":
		c.SecureJSON(code, obj)
	case "indentedjson":
		c.IndentedJSON(code, obj)
	case "asciijson":
		c.AsciiJSON(code, obj)
	case "purejson":
		c.PureJSON(code, obj)
	case "xml":
		c.XML(code, obj)
	case "html":
		c.HTML(code, "error", obj)
	case "yaml":
		c.YAML(code, obj)
	case "toml":
		c.TOML(code, obj)
	case "protobuf":
		c.ProtoBuf(code, obj)
	case "ini":
		c.String(code, formatAsINI(obj))
	}

}

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

func isSubscribed(bid string, uid int64, chat int64) (bool, error) {

	for _, client := range tele.Clients {
		if client.Bot.Bid == bid {
			member, err := client.Telebot.ChatMemberOf(&telebot.Chat{ID: chat}, &telebot.User{ID: uid})
			if err != nil {
				return false, err
			}
			if member.Role != telebot.Left && member.Role != telebot.Kicked {
				return true, nil
			}
		}
	}
	return false, nil

}

func checkType(c *gin.Context) string {

	if len(c.Query("type")) == 0 {
		return "json"
	}
	return c.Query("type")

}

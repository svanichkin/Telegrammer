package rest

import (
	"fmt"
	"main/conf"
	"main/data"
	"main/tele"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func sendHandler(c *gin.Context) {

	// Check return type

	rt := checkType(c)

	// Validate bid

	bid := checkBid(c.Query("bid"))
	if len(bid) == 0 {
		responseError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Convertation uid to int64

	uid := checkInt(c.Query("uid"))
	aid := c.Query("aid")
	if uid == 0 {
		if len(aid) == 0 {
			responseError(c, fmt.Errorf("bad request: uid is not valid"), rt)
			return
		}
		var err error
		if uid, err = data.FindUidWithAlias(bid, aid); err != nil {
			responseError(c, fmt.Errorf("bad request: uid for alias %s not found: %v", aid, err), rt)
			return
		}
	}

	// Validate chat

	cid := checkInt(c.Query("cid"))

	// Main switch

	switch {
	case c.Query("text") != "":
		if err := tele.SendMessage(bid, uid, cid, c.Query("text")); err != nil {
			responseError(c, err, rt)
			return
		}

	case c.Query("html") != "":
		if err := tele.SendHtml(bid, uid, cid, c.Query("html")); err != nil {
			responseError(c, err, rt)
			return
		}

	case c.Query("markdown") != "":
		if err := tele.SendMarkdown(bid, uid, cid, c.Query("markdown")); err != nil {
			responseError(c, err, rt)
			return
		}

	case c.Query("markdownv2") != "":
		if err := tele.SendMarkdownV2(bid, uid, cid, c.Query("markdownv2")); err != nil {
			responseError(c, err, rt)
			return
		}

	case c.Query("photo") != "":

		// Check photo type [base64, url or local file]

		photo := c.Query("photo")
		if strings.HasPrefix(photo, "data:image/") {
			if err := tele.SendPhotoBase64(bid, uid, cid, photo); err != nil {
				responseError(c, err, rt)
				return
			}
		} else if strings.HasPrefix(photo, "http") {
			if err := tele.SendPhotoUrl(bid, uid, cid, photo); err != nil {
				responseError(c, err, rt)
				return
			}
		} else {
			if err := tele.SendPhotoPath(bid, uid, cid, photo); err != nil {
				responseError(c, err, rt)
				return
			}
		}

	case c.Query("file") != "":
		if err := tele.SendFile(bid, uid, cid, c.Query("file")); err != nil {
			responseError(c, err, rt)
			return
		}

	default:
		responseError(c, fmt.Errorf("bad request: missing command"), rt)
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
		responseError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Convertation uid to int64

	uid := checkInt(c.Query("uid"))
	aid := c.Query("aid")
	if uid == 0 {
		if len(aid) == 0 {
			responseError(c, fmt.Errorf("bad request: uid is not valid"), rt)
			return
		}
		var err error
		if uid, err = data.FindUidWithAlias(bid, aid); err != nil {
			responseError(c, fmt.Errorf("bad request: uid for alias %s not found: %v", aid, err), rt)
			return
		}
	}

	// Validate chat

	cid := checkInt(c.Query("cid"))
	if cid == 0 {
		responseError(c, fmt.Errorf("bad request: cid is not valid"), rt)
		return
	}

	// Check subscribe user

	isSubscribed, err := tele.CheckSubscribe(bid, uid, cid)
	if err != nil {
		responseError(c, fmt.Errorf("bad check: %v", err), rt)
		return
	}

	if !isSubscribed {
		responseError(c, fmt.Errorf("bad check: user is not subscribed"), rt)
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
		responseError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Convertation uid to int64

	uid := checkInt(c.Query("uid"))
	aid := c.Query("aid")
	if uid == 0 {
		if len(aid) == 0 {
			responseError(c, fmt.Errorf("bad request: uid is not valid"), rt)
			return
		}
		var err error
		if uid, err = data.FindUidWithAlias(bid, aid); err != nil {
			responseError(c, fmt.Errorf("bad request: uid for alias %s not found: %v", aid, err), rt)
			return
		}
	}

	// Validate do

	do := c.Query("do")
	if len(do) == 0 {
		responseError(c, fmt.Errorf("bad request: do is not valid"), rt)
		return
	}

	// Validate key and value or keys and values

	if do == "set" {

		key := c.Query("key")
		value := c.Query("value")
		if len(key) > 0 && len(value) > 0 {
			if err := data.SetKey(bid, uid, key, value); err != nil {
				responseError(c, fmt.Errorf("bad request: %v", err), rt)
			}
			return
		}

		keys := checkStrings(c.Query("keys"))
		values := checkStrings(c.Query("values"))
		if len(keys) > 0 && len(values) > 0 && len(keys) == len(values) {
			for i, key := range keys {
				if err := data.SetKey(bid, uid, key, values[i]); err != nil {
					responseError(c, fmt.Errorf("bad request: %v", err), rt)
				}
			}
			return
		}

		if len(key) == 0 {
			responseError(c, fmt.Errorf("bad request: key is not valid"), rt)
			return
		}
		if len(value) == 0 {
			responseError(c, fmt.Errorf("bad request: value is not valid"), rt)
			return
		}
		if len(keys) == 0 {
			responseError(c, fmt.Errorf("bad request: keys is not valid"), rt)
			return
		}
		if len(values) == 0 {
			responseError(c, fmt.Errorf("bad request: values is not valid"), rt)
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
				responseError(c, fmt.Errorf("bad request: %v", err), rt)
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
					responseError(c, fmt.Errorf("bad request: %v", err), rt)
					return
				}
				r[key] = value
			}
			response(c, http.StatusOK, r, rt)
			return
		}

		if len(key) == 0 {
			responseError(c, fmt.Errorf("bad request: key is not valid"), rt)
			return
		}
		if len(keys) == 0 {
			responseError(c, fmt.Errorf("bad request: keys is not valid"), rt)
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
				responseError(c, fmt.Errorf("bad request: %v", err), rt)
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
					responseError(c, fmt.Errorf("bad request: %v", err), rt)
					return
				}
			}
			c.Status(http.StatusOK)
			return
		}

		if len(key) == 0 {
			responseError(c, fmt.Errorf("bad request: key is not valid"), rt)
			return
		}
		if len(keys) == 0 {
			responseError(c, fmt.Errorf("bad request: keys is not valid"), rt)
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
		responseError(c, fmt.Errorf("bad request: bid is not valid"), rt)
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
		responseError(c, fmt.Errorf("bad request: command name is not valid"), rt)
		return
	}

	// Do action

	err := tele.DoAction(bid, uid, cid, commandName, c.Query("data"), c.Query("firstname"), c.Query("lastname"), c.Query("username"), c.Query("language"))
	if err != nil {
		responseError(c, err, rt)
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
		responseError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Validate alias

	aid := c.Query("aid")
	if len(aid) == 0 {
		responseError(c, fmt.Errorf("bad request: alias aid is not valid"), rt)
		return
	}

	// Validate do

	do := c.Query("do")
	if len(do) == 0 {
		responseError(c, fmt.Errorf("bad request: do is not valid"), rt)
		return
	}

	// Validate key and value or keys and values

	if do == "add" {
		if uid == 0 {
			responseError(c, fmt.Errorf("bad request: uid is not valid"), rt)
			return
		}
		if err := data.AddAlias(bid, uid, aid); err != nil {
			responseError(c, fmt.Errorf("bad request: alias %s not added: %v", aid, err), rt)
			return
		}
		c.Status(http.StatusOK)
		return
	}

	if do == "remove" {
		if err := data.RemoveAlias(bid, aid); err != nil {
			responseError(c, fmt.Errorf("bad request: alias %s not added: %v", aid, err), rt)
			return
		}
		c.Status(http.StatusOK)
		return
	}

	if do == "find" {
		uid, err := data.FindUidWithAlias(bid, aid)
		if err != nil {
			responseError(c, fmt.Errorf("bad request: uid for alias %s not found", aid), rt)
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
		responseError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Restart

	if err := conf.RestartBot(bid); err != nil {
		responseError(c, err, rt)
		return
	}
	if err := tele.RestartClient(bid); err != nil {
		responseError(c, err, rt)
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
		responseError(c, fmt.Errorf("bad request: bid is not valid"), rt)
		return
	}

	// Convertation uid and aid and chat

	uid := checkInt(c.Query("uid"))
	aid := c.Query("aid")
	cid := checkInt(c.Query("cid"))

	if len(aid) == 0 && cid == 0 && uid == 0 {
		responseError(c, fmt.Errorf("bad request: need uid, aid or cid"), rt)
		return
	}

	if len(aid) > 0 {
		var err error
		if uid, err = data.FindUidWithAlias(bid, aid); err != nil {
			responseError(c, fmt.Errorf("bad request: uid for alias %s not found: %v", aid, err), rt)
			return
		}
	}

	// Validate group

	group := c.Query("group")
	if len(group) == 0 {
		responseError(c, fmt.Errorf("bad request: group is not valid"), rt)
		return
	}

	// Validate do

	do := c.Query("do")
	if len(do) == 0 {
		responseError(c, fmt.Errorf("bad request: do is not valid"), rt)
		return
	}

	// Do

	if do == "set" {
		if uid != 0 {
			if err := conf.SetUserForGroup(bid, uid, group); err != nil {
				responseError(c, fmt.Errorf("bad request: user not added to group %s: %v", group, err), rt)
				return
			}
		} else {
			if err := conf.SetChatForGroup(bid, cid, group); err != nil {
				responseError(c, fmt.Errorf("bad request: chat not added to group %s: %v", group, err), rt)
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
				responseError(c, fmt.Errorf("bad request: user not found: %v", err), rt)
				return
			}
		} else {
			if groups, err = conf.GetChatForGroup(bid, cid); err != nil {
				responseError(c, fmt.Errorf("bad request: chat not found: %v", err), rt)
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
				responseError(c, fmt.Errorf("bad request: user not removed from group %s: %v", group, err), rt)
				return
			}
		} else {
			if err := conf.RemoveChatForGroup(bid, cid, group); err != nil {
				responseError(c, fmt.Errorf("bad request: chat not removed from group %s: %v", group, err), rt)
				return
			}
		}
		c.Status(http.StatusOK)
		return
	}

}

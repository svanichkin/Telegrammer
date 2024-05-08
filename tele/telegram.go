package tele

import (
	"fmt"
	"io"
	"main/conf"
	"main/data"
	"os"
	"os/exec"
	"strings"
	"time"

	"gopkg.in/telebot.v3"
)

var Clients map[string]*Client

type Client struct {
	Bot     *conf.Bot
	Telebot *telebot.Bot
}

var inputs map[int64]Input

type Input struct {
	stdin io.WriteCloser
	bid   string
}

func Init() error {

	inputs = make(map[int64]Input)
	Clients = make(map[string]*Client)

	for _, bot := range conf.Config.Bots {
		if err := initBot(bot); err != nil {
			return err
		}
	}
	return nil

}

func initBot(bot *conf.Bot) error {

	c, err := initClient(&Client{Bot: bot})
	if err != nil {
		return err
	}
	Clients[bot.Bid] = c
	bot.Senders = make(map[int64]*conf.Sender)

	return nil

}

func initClient(client *Client) (*Client, error) {

	// Create new client

	var err error
	if client.Telebot, err = telebot.NewBot(telebot.Settings{Token: client.Bot.Key, Poller: &telebot.LongPoller{Timeout: 10 * time.Second}}); err != nil {
		return client, err
	}

	// Set commands on user language

	for _, command := range client.Bot.Commands {
		client.Telebot.Handle("/"+command.Name, func(context telebot.Context) error {
			if command.Name == "start" {
				if err := setCommands(context, client); err != nil {
					return err
				}
			}
			sender := senderFromContext(context.Sender())
			client.Bot.Senders[sender.Uid] = sender
			output, err := doPythonScript(client, sender, context, command.Action, command, context.Data(), "action", data.AllAliases(client.Bot.Bid, sender.Uid), checkAccess(client, command, sender.Uid, context.Chat().ID))
			if err != nil {
				return fmt.Errorf("error walking through bots directory: %v , otuput:\n%s", err, output)
			}
			if len(output) > 0 {
				return context.Reply(string(output))
			}
			return nil
		})
	}

	// Set listener to simple text from telegram

	client.Telebot.Handle(telebot.OnText, func(context telebot.Context) error {
		uid := context.Sender().ID
		input := inputs[uid]
		if input.bid == client.Bot.Bid {
			defer input.stdin.Close()
			defer delete(inputs, uid)
			if _, err := io.WriteString(input.stdin, context.Text()); err != nil {
				return fmt.Errorf("python: error writing to stdin: %v", err)
			}
		}
		return nil
	})

	// Start client

	go func() {
		client.Telebot.Start()
	}()

	return client, nil

}

func RestartClient(bid string) error {

	// Get bot

	bot := conf.Config.Bots[bid]
	if bot == nil {
		return fmt.Errorf("bot with bid %s, not found", bid)
	}

	// Stop and remove old client

	client := Clients[bid]
	client.Telebot.Stop()
	Clients[bid] = nil

	// Init new client

	if err := initBot(bot); err != nil {
		return err
	}

	return nil

}

func setCommands(context telebot.Context, client *Client) error {

	l := context.Sender().LanguageCode
	cmds := []telebot.Command{}
	for _, command := range client.Bot.Commands {
		if command.Showed {
			sender := senderFromContext(context.Sender())
			client.Bot.Senders[sender.Uid] = sender
			output, err := doPythonScript(client, sender, context, command.Detail, command, context.Data(), "detail", data.AllAliases(client.Bot.Bid, sender.Uid), checkAccess(client, command, sender.Uid, context.Chat().ID))
			if err != nil {
				return err
			}
			description := string(output)
			if len(description) == 0 {
				description = ""
			}
			cmds = append(cmds, telebot.Command{Text: "/" + command.Name, Description: description})
		}
	}
	return client.Telebot.SetCommands(l, cmds)

}

func doPythonScript(client *Client, sender *conf.Sender, context telebot.Context, path string, command conf.Command, data string, dt string, aliases []string, access bool) (string, error) {

	var cid string
	if context != nil {
		cid = fmt.Sprint(context.Chat().ID)
	}

	cmd := exec.Command("python", path)
	cmd.Env = append(os.Environ(),
		"LANG="+sender.LanguageCode,
		"COMMAND="+command.Name,
		"UID="+fmt.Sprint(sender.Uid),
		"BID="+client.Bot.Bid,
		"CID="+cid,
		"FIRST_NAME="+sender.FirstName,
		"LAST_NAME="+sender.LastName,
		"USER_NAME="+sender.Username,
		"DO_TYPE="+dt,
		"DATA="+data,
		"ALIASES="+strings.Join(aliases, ","),
		"ACCESS_GROUPS="+strings.Join(command.Access, ","),
		"ACCESS="+boolToString(access),
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("python: error creating stdout pipe: %v", err)
	}
	defer stdout.Close()

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", fmt.Errorf("python: error creating stdin pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("python: error starting command: %v", err)
	}

	inputs[sender.Uid] = Input{stdin: stdin, bid: client.Bot.Bid}

	output, err := io.ReadAll(stdout)
	if err != nil {
		return "", fmt.Errorf("python: error reading from stdout: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("python: command finished with error: %v", err)
	}

	return string(output), nil

}

func DoAction(bid string, uid int64, cid int64, commandName string, d string, firstName string, lastName string, username string, language string) error {

	client := Clients[bid]
	var command conf.Command
	for _, c := range client.Bot.Commands {
		if c.Name == commandName {
			command = c
			break
		}
	}

	// Get sender from bot OR from cache OR from URL parameter

	sender := client.Bot.Senders[uid]
	if sender == nil {
		sender = senderFromContext(client.Telebot.Me)
	}
	if len(firstName) > 0 {
		sender.FirstName = firstName
	}
	if len(lastName) > 0 {
		sender.LastName = lastName
	}
	if len(username) > 0 {
		sender.Username = username
	}
	if len(language) > 0 {
		sender.LanguageCode = language
	}
	output, err := doPythonScript(client, sender, nil, command.Action, command, d, "action", data.AllAliases(client.Bot.Bid, uid), checkAccess(client, command, sender.Uid, 0))
	if err != nil {
		return fmt.Errorf("error walking through bots directory: %v , otuput:\n%s", err, output)
	}
	if len(output) > 0 {
		if err := SendMessage(bid, uid, cid, string(output)); err != nil {
			return err
		}
	}
	return nil

}

func CheckSubscribe(bid string, uid int64, cid int64) (bool, error) {

	member, err := Clients[bid].Telebot.ChatMemberOf(&telebot.Chat{ID: cid}, &telebot.User{ID: uid})
	if err != nil {
		return false, err
	}
	if member.Role != telebot.Left && member.Role != telebot.Kicked {
		return true, nil
	}
	return false, nil

}

func boolToString(b bool) string {

	if b {
		return "true"
	}
	return "false"

}

func checkAccess(client *Client, command conf.Command, uid int64, cid int64) bool {

	for _, name := range command.Access {
		if group, ok := client.Bot.Groups[name]; ok {
			if uid != 0 && cid != 0 {
				if len(group.Users) > 0 && len(group.Chats) > 0 {
					return containsValue(group.Users, uid) && containsValue(group.Chats, cid)
				} else if len(group.Users) > 0 {
					return containsValue(group.Users, uid)
				} else if containsValue(group.Chats, cid) {
					return containsValue(group.Chats, cid)
				}
			} else if uid != 0 {
				return containsValue(group.Users, uid)
			} else if cid != 0 {
				return containsValue(group.Chats, cid)
			}
		}
	}

	return true

}

func containsValue(slice []int64, value int64) bool {

	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false

}

func senderFromContext(user *telebot.User) *conf.Sender {

	return &conf.Sender{
		Uid:          user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		LanguageCode: user.LanguageCode,
	}

}

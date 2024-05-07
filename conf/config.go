package conf

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

type Server struct {
	Host string
	Port int
}

type config struct {
	Server Server
	Bots   map[string]*Bot
}

type Bot struct {
	Path     string
	Bid      string
	Key      string
	Commands []Command
	Senders  map[int64]*Sender
	Groups   map[string]*AccessGroup
}

type Sender struct {
	Uid          int64
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
}

type Command struct {
	Name   string
	Showed bool
	Detail string
	Action string
	Access []string
}

type AccessGroup struct {
	Name  string
	Users []int64
	Chats []int64
}

var Config *config

var botsPath string

func Init() error {

	// Load main config.ini

	cfg, err := ini.Load("./config.ini")
	if err != nil {
		return err
	}

	// Parse server configuration

	var c config
	if c.Server.Host = cfg.Section("").Key("host").String(); len(c.Server.Host) == 0 {
		return errors.New("config error, field 'host' not found or not have values")
	}
	if c.Server.Port, err = cfg.Section("").Key("port").Int(); err != nil {
		return errors.New("config error, field 'port' not found or not have values")
	}
	Config = &c

	// Parse bots path

	if botsPath = cfg.Section("").Key("bots").String(); len(botsPath) == 0 {
		return errors.New("config error, field 'bots' not found or not have values")
	}

	return initBots()

}

func initBots() error {

	// Check bots folder

	if _, err := os.Stat(botsPath); os.IsNotExist(err) {
		return nil
	}

	// Read folder by folder

	Config.Bots = make(map[string]*Bot)
	err := filepath.WalkDir(botsPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}
		if path == botsPath {
			return nil
		}
		bot := &Bot{Path: "./" + path}
		if err := initBot(bot); err != nil {
			return nil
			// return fmt.Errorf("failed to init bot at %s: %v", path, err)
		}
		Config.Bots[bot.Bid] = bot
		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking through bots directory: %v", err)
	}

	return nil

}

func initBot(bot *Bot) error {

	// Load bot config.ini

	c := filepath.Join(bot.Path, "config.ini")
	if _, err := os.Stat(c); os.IsNotExist(err) {
		return fmt.Errorf("config.ini not found in: %s", bot.Path)
	}
	cfg, err := ini.Load(c)
	if err != nil {
		return fmt.Errorf("failed to load config.ini from %s: %v", c, err)
	}

	// Parse commands

	if bot.Bid = cfg.Section("").Key("bid").String(); len(bot.Bid) == 0 {
		return fmt.Errorf("error reading bid from config: %v", err)
	}
	if bot.Key = cfg.Section("").Key("key").String(); len(bot.Key) == 0 {
		return fmt.Errorf("error reading key from config: %v", err)
	}
	cmds := strings.Split(strings.ReplaceAll(cfg.Section("").Key("commands").String(), " ", ""), ",")
	if len(cmds) == 0 {
		return fmt.Errorf("error reading commands from config: %v", err)
	}
	for _, cmd := range cmds {
		command := Command{Name: cmd}
		s := cfg.Section(cmd)
		if command.Showed, err = s.Key("showed").Bool(); err != nil {
			command.Showed = true
		}
		if command.Showed {
			if command.Detail = s.Key("detail").String(); len(command.Detail) == 0 {
				return fmt.Errorf("error reading description from config: %v", err)
			}
		}
		if command.Action = s.Key("action").String(); len(command.Action) == 0 {
			return fmt.Errorf("error reading action from config: %v", err)
		}
		command.Access = strings.Split(strings.ReplaceAll(s.Key("access").String(), " ", ""), ",")
		bot.Commands = append(bot.Commands, command)
	}

	// Parse access groups

	grps := strings.Split(strings.ReplaceAll(cfg.Section("").Key("groups").String(), " ", ""), ",")
	if len(grps) == 0 {
		return nil
	}
	bot.Groups = make(map[string]*AccessGroup)
	for _, grp := range grps {
		access := &AccessGroup{Name: grp}
		s := cfg.Section(grp)
		users := s.Key("users").String()
		if len(users) > 0 {
			access.Users = stringToIds(users)
		}
		chats := s.Key("chats").String()
		if len(chats) > 0 {
			access.Chats = stringToIds(chats)
		}
		bot.Groups[grp] = access
	}
	return nil

}

func RestartBot(bid string) error {

	// Remove old bot settings

	Config.Bots[bid] = nil

	// Init new bot settings

	if _, err := os.Stat(botsPath); os.IsNotExist(err) {
		return nil
	}
	err := filepath.WalkDir(botsPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			return nil
		}
		if path == botsPath {
			return nil
		}
		bot := &Bot{Path: "./" + path}
		if err := initBot(bot); err != nil {
			return nil
		}
		if bot.Bid == bid {
			Config.Bots[bot.Bid] = bot
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking through bots directory: %v", err)
	}
	return nil

}

func stringToIds(s string) []int64 {

	strs := strings.Split(s, ",")
	var ids []int64
	for _, str := range strs {
		str = strings.TrimSpace(str)
		id, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	return ids

}

func idsToString(ids []int64) string {

	strs := make([]string, len(ids))
	for i, id := range ids {
		strs[i] = strconv.FormatInt(id, 10)
	}
	return strings.Join(strs, ", ")

}

func SetUserForGroup(bid string, uid int64, group string) error {

	// Add uid to ram

	bot := Config.Bots[bid]
	if bot == nil {
		return fmt.Errorf("error bid %s not found", bid)
	}
	accessGroup := bot.Groups[group]
	if accessGroup == nil {
		return fmt.Errorf("error group %s not found", group)
	}

	// If uid already in users slice

	for _, u := range accessGroup.Users {
		if u == uid {
			return nil
		}
	}

	// If uid its new, add

	accessGroup.Users = append(accessGroup.Users, uid)

	// Add uid to config, load bot config.ini

	c := filepath.Join(bot.Path, "config.ini")
	if _, err := os.Stat(c); os.IsNotExist(err) {
		return fmt.Errorf("config.ini not found in: %s", bot.Path)
	}
	cfg, err := ini.Load(c)
	if err != nil {
		return fmt.Errorf("failed to load config.ini from %s: %v", c, err)
	}

	cfg.Section(group).Key("users").SetValue(idsToString(accessGroup.Users))
	err = cfg.SaveTo(c)
	if err != nil {
		return fmt.Errorf("failed to save config.ini: %v", err)
	}
	return nil

}

func SetChatForGroup(bid string, cid int64, group string) error {

	// Add uid to ram

	bot := Config.Bots[bid]
	if bot == nil {
		return fmt.Errorf("error bid %s not found", bid)
	}
	accessGroup := bot.Groups[group]
	if accessGroup == nil {
		return fmt.Errorf("error group %s not found", group)
	}

	// If cid already in chat slice

	for _, c := range accessGroup.Chats {
		if c == cid {
			return nil
		}
	}

	// If cid its new, add

	accessGroup.Chats = append(accessGroup.Chats, cid)

	// Add uid to config, load bot config.ini

	c := filepath.Join(bot.Path, "config.ini")
	if _, err := os.Stat(c); os.IsNotExist(err) {
		return fmt.Errorf("config.ini not found in: %s", bot.Path)
	}
	cfg, err := ini.Load(c)
	if err != nil {
		return fmt.Errorf("failed to load config.ini from %s: %v", c, err)
	}

	cfg.Section(group).Key("chats").SetValue(idsToString(accessGroup.Chats))
	err = cfg.SaveTo(c)
	if err != nil {
		return fmt.Errorf("failed to save config.ini: %v", err)
	}
	return nil

}

func GetUserForGroup(bid string, uid int64) ([]AccessGroup, error) {

	// Find uid from ram

	bot := Config.Bots[bid]
	if bot == nil {
		return nil, fmt.Errorf("error bid %s not found", bid)
	}
	var ag []AccessGroup
	for _, accessGroup := range bot.Groups {
		for _, u := range accessGroup.Users {
			if u == uid {
				ag = append(ag, *accessGroup)
			}
		}
	}
	return ag, nil

}

func GetChatForGroup(bid string, cid int64) ([]AccessGroup, error) {

	// Find uid from ram

	bot := Config.Bots[bid]
	if bot == nil {
		return nil, fmt.Errorf("error bid %s not found", bid)
	}
	var ag []AccessGroup
	for _, accessGroup := range bot.Groups {
		for _, c := range accessGroup.Chats {
			if c == cid {
				ag = append(ag, *accessGroup)
			}
		}
	}
	return ag, nil

}

func RemoveUserForGroup(bid string, uid int64, group string) error {

	// Remove uid from ram

	bot := Config.Bots[bid]
	if bot == nil {
		return fmt.Errorf("error bid %s not found", bid)
	}
	accessGroup := bot.Groups[group]
	if accessGroup == nil {
		return fmt.Errorf("error group %s not found", group)
	}

	// If uid in users slice

	index := -1
	for i, u := range accessGroup.Users {
		if u == uid {
			index = i
			break
		}
	}
	if index == -1 {
		return nil
	}

	// If uid its new, add

	accessGroup.Users = append(accessGroup.Users[:index], accessGroup.Users[index+1:]...)

	// Add uid to config, load bot config.ini

	c := filepath.Join(bot.Path, "config.ini")
	cfg, err := ini.Load(c)
	if err != nil {
		return fmt.Errorf("failed to load config.ini from %s: %v", c, err)
	}

	cfg.Section(group).Key("users").SetValue(idsToString(accessGroup.Users))
	err = cfg.SaveTo(c)
	if err != nil {
		return fmt.Errorf("failed to save config.ini: %v", err)
	}
	return nil

}

func RemoveChatForGroup(bid string, uid int64, group string) error {

	// Remove uid from ram

	bot := Config.Bots[bid]
	if bot == nil {
		return fmt.Errorf("error bid %s not found", bid)
	}
	accessGroup := bot.Groups[group]
	if accessGroup == nil {
		return fmt.Errorf("error group %s not found", group)
	}

	// If uid in users slice

	index := -1
	for i, c := range accessGroup.Chats {
		if c == uid {
			index = i
			break
		}
	}
	if index == -1 {
		return nil
	}

	// If uid its new, add

	accessGroup.Chats = append(accessGroup.Chats[:index], accessGroup.Chats[index+1:]...)

	// Add uid to config, load bot config.ini

	c := filepath.Join(bot.Path, "config.ini")
	cfg, err := ini.Load(c)
	if err != nil {
		return fmt.Errorf("failed to load config.ini from %s: %v", c, err)
	}

	cfg.Section(group).Key("chats").SetValue(idsToString(accessGroup.Chats))
	err = cfg.SaveTo(c)
	if err != nil {
		return fmt.Errorf("failed to save config.ini: %v", err)
	}
	return nil

}

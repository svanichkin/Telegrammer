package data

import (
	"bufio"
	"fmt"
	"main/conf"
	"os"
	"regexp"

	"gopkg.in/ini.v1"
)

func AddAlias(bid string, uid int64, alias string) error {

	bot := conf.Config.Bots[bid]
	if bot == nil {
		return fmt.Errorf("bad data: bid %s not found", bid)
	}

	filePath := bot.Path + "/aliases.ini"
	cfg, err := ini.Load(filePath)
	if err != nil {
		cfg = ini.Empty()
	}
	section := cfg.Section("")
	section.Key(alias).SetValue(fmt.Sprint(uid))
	err = cfg.SaveTo(filePath)
	if err != nil {
		return fmt.Errorf("bad data: error save alias file: %v", err)
	}
	return nil

}

func RemoveAlias(bid string, alias string) error {

	bot := conf.Config.Bots[bid]
	if bot == nil {
		return fmt.Errorf("bad data: bid %s not found", bid)
	}

	filePath := bot.Path + "/aliases.ini"
	cfg, err := ini.Load(filePath)
	if err != nil {
		cfg = ini.Empty()
	}
	section := cfg.Section("")
	section.DeleteKey(alias)
	err = cfg.SaveTo(filePath)
	if err != nil {
		return fmt.Errorf("bad data: error save alias file: %v", err)
	}
	return nil

}

func FindUidWithAlias(bid string, alias string) (int64, error) {

	bot := conf.Config.Bots[bid]
	if bot == nil {
		return 0, fmt.Errorf("bad data: bid %s not found", bid)
	}

	filePath := bot.Path + "/aliases.ini"
	cfg, err := ini.Load(filePath)
	if err != nil {
		cfg = ini.Empty()
	}
	section := cfg.Section("")
	return section.Key(alias).Int64()

}

func AllAliases(bid string, uid int64) []string {

	bot := conf.Config.Bots[bid]
	if bot == nil {
		return nil
	}

	var aliases []string
	filePath := bot.Path + "/aliases.ini"
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	regexPattern := fmt.Sprintf(`^%s\s*=\s*(.+)`, regexp.QuoteMeta(fmt.Sprint(uid)))
	re := regexp.MustCompile(regexPattern)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			aliases = append(aliases, matches[1])
		}
	}
	if err := scanner.Err(); err != nil {
		return nil
	}
	return aliases

}

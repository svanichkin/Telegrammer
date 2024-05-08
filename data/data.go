package data

import (
	"fmt"
	"main/conf"
	"os"
)

func SetKey(bid string, uid int64, key string, value string) error {

	for _, bot := range conf.Config.Bots {
		if bot.Bid == bid {

			// Create data folder if needed

			folderPath := bot.Path + "/data/" + fmt.Sprint(uid) + "/"
			err := createFolder(folderPath)
			if err != nil {
				return fmt.Errorf("bad data: problem with data folder: %v", err)
			}

			// Write value to file

			err = os.WriteFile(folderPath+key+".txt", []byte(value), 0644)
			if err != nil {
				return fmt.Errorf("bad data: problem with writing file: %v", err)
			}
			return nil
		}
	}
	return fmt.Errorf("bad data: bid %s not found", bid)

}

func GetKey(bid string, uid int64, key string) (string, error) {

	for _, bot := range conf.Config.Bots {
		if bot.Bid == bid {
			folderPath := bot.Path + "/data/" + fmt.Sprint(uid) + "/"
			content, err := os.ReadFile(folderPath + key + ".txt")
			if err != nil {
				return "", fmt.Errorf("bad data: problem with reading file: %v", err)
			}
			return string(content), nil
		}
	}
	return "", fmt.Errorf("bad data: bid %s not found", bid)

}

func RemoveKey(bid string, uid int64, key string) error {

	for _, bot := range conf.Config.Bots {
		if bot.Bid == bid {
			folderPath := bot.Path + "/data/" + fmt.Sprint(uid) + "/"
			os.Remove(folderPath + key + ".txt")
			return nil
		}
	}
	return fmt.Errorf("bad data: bid %s not found", bid)

}

// Helpers

func createFolder(folderPath string) error {

	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		return err
	}
	return err

}

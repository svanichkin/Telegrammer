package tele

import (
	"encoding/base64"
	"log"
	"strings"

	"gopkg.in/telebot.v3"
)

func SendHtml(bid string, uid int64, cid int64, html string) error {

	var recepient telebot.User
	if cid != 0 {
		recepient = telebot.User{ID: cid}
	} else {
		recepient = telebot.User{ID: uid}
	}
	_, err := Clients[bid].Telebot.Send(&recepient, html, &telebot.SendOptions{ParseMode: telebot.ModeHTML})
	return err

}

func SendMarkdown(bid string, uid int64, cid int64, md string) error {

	var recepient telebot.User
	if cid != 0 {
		recepient = telebot.User{ID: cid}
	} else {
		recepient = telebot.User{ID: uid}
	}
	_, err := Clients[bid].Telebot.Send(&recepient, md, &telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
	return err

}

func SendMarkdownV2(bid string, uid int64, cid int64, md string) error {

	var recepient telebot.User
	if cid != 0 {
		recepient = telebot.User{ID: cid}
	} else {
		recepient = telebot.User{ID: uid}
	}
	_, err := Clients[bid].Telebot.Send(&recepient, md, &telebot.SendOptions{ParseMode: telebot.ModeMarkdownV2})
	return err

}

func SendMessage(bid string, uid int64, cid int64, text string) error {

	var recepient telebot.User
	if cid != 0 {
		recepient = telebot.User{ID: cid}
	} else {
		recepient = telebot.User{ID: uid}
	}
	_, err := Clients[bid].Telebot.Send(&recepient, text, &telebot.SendOptions{})
	return err

}

func SendPhotoBase64(bid string, uid int64, cid int64, b64 string) error {

	var recepient telebot.User
	if cid != 0 {
		recepient = telebot.User{ID: cid}
	} else {
		recepient = telebot.User{ID: uid}
	}
	i := strings.Index(b64, ",")
	if i < 0 {
		log.Fatal("no comma")
	}
	b := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64[i+1:]))
	p := &telebot.Photo{File: telebot.FromReader(b)}
	_, err := Clients[bid].Telebot.SendAlbum(&recepient, telebot.Album{p})
	return err

}

func SendPhotoUrl(bid string, uid int64, cid int64, url string) error {

	var recepient telebot.User
	if cid != 0 {
		recepient = telebot.User{ID: cid}
	} else {
		recepient = telebot.User{ID: uid}
	}
	p := &telebot.Photo{File: telebot.FromURL(url)}
	_, err := Clients[bid].Telebot.SendAlbum(&recepient, telebot.Album{p})
	return err

}

func SendPhotoPath(bid string, uid int64, cid int64, path string) error {

	var recepient telebot.User
	if cid != 0 {
		recepient = telebot.User{ID: cid}
	} else {
		recepient = telebot.User{ID: uid}
	}
	p := &telebot.Photo{File: telebot.FromDisk(path)}
	_, err := Clients[bid].Telebot.SendAlbum(&recepient, telebot.Album{p})
	return err

}

func SendFile(bid string, uid int64, cid int64, path string) error {

	var recepient telebot.User
	if cid != 0 {
		recepient = telebot.User{ID: cid}
	} else {
		recepient = telebot.User{ID: uid}
	}
	_, err := Clients[bid].Telebot.Send(&recepient, &telebot.Document{FileName: path, File: telebot.FromDisk(path)})
	return err

}

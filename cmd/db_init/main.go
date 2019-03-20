package main

import (
	"fmt"
	"log"

	"github.com/heyuan110/go-mega-code/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)
	db.DropTableIfExists(model.User{}, model.Post{})
	db.CreateTable(model.User{}, model.Post{})
	users := []model.User{
		{
			Username:     "bruce",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "heyuan110@gmail.com",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("heyuan110@gmail.com")),
			Posts: []model.Post{
				{Body: "Beautiful day in Portland!"},
				{Body: "Sun shine is beautiful!"},
			},
		},
		{
			Username:     "bonfy",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "i@bonfy.im",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("i@bonfy.im")),
			Posts: []model.Post{
				{Body: "Beautiful day in Portland!"},
			},
		},
	}

	for _, u := range users {
		db.Debug().Create(&u)
	}

}

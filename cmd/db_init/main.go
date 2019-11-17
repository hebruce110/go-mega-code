package main

import (
	"fmt"
	"go-mega-code/model"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	log.Println("DB Init ...")
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)
	db.DropTableIfExists(model.User{}, model.Post{}, "follower")
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
			Username:     "tonfy",
			PasswordHash: model.GeneratePasswordHash("abc123"),
			Email:        "i@tonfy.im",
			Avatar:       fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", model.Md5("i@tonfy.im")),
			Posts: []model.Post{
				{Body: "Beautiful day in USA!"},
			},
		},
	}

	for _, u := range users {
		db.Debug().Create(&u)
	}

	model.AddUser("bonfy", "abc123", "i@bonfy.im")
	model.AddUser("rene", "abc123", "rene@test.com")

	u1, _ := model.GetUserByUsername("bonfy")
	u1.CreatePost("Beautiful day in Portland!")
	model.UpdateAboutMe(u1.Username, `I'm the author of Go-Mega Tutorial you are reading now!`)

	u2, _ := model.GetUserByUsername("rene")
	u2.CreatePost("The Avengers movie was so cool!")
	u2.CreatePost("Sun shine is beautiful")

	u1.Follow(u2.Username)

}

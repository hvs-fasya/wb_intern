package main

import (
	"log"

	"github.com/hvs-fasya/wb_intern/pkg/memstore"
	"github.com/hvs-fasya/wb_intern/pkg/models"
	nfier "github.com/hvs-fasya/wb_intern/pkg/notifier"
	"github.com/hvs-fasya/wb_intern/pkg/service"
)

var (
	newUsers = [2]models.CreateUserInput{
		{Name: "new_user"},
		{Name: ""},
	}
	notifyTmpl = "user %s created\n"
)

func main() {
	n := nfier.NewNotifier(nfier.Cfg{Tmpl: notifyTmpl})
	s := memstore.NewMemStore()
	srv := service.NewUserService(n, s)

	for _, u := range newUsers {
		if err := srv.CreateUser(u); err != nil {
			log.Printf("create user '" + u.Name + "' error: " + err.Error())
		}
	}
}

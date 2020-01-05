package service

import (
	"testing"

	"github.com/hvs-fasya/wb_intern/pkg/models"
	nfier "github.com/hvs-fasya/wb_intern/pkg/notifier"
)

func Test_userService_CreateUser(t *testing.T) {
	var n = nfier.NewNotifier(nfier.Cfg{Tmpl: "user %s created\n"})
	var s = NewUserService(n)
	type args struct {
		inp models.CreateUserInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"create_good_user", args{models.CreateUserInput{Name: "test_user"}}, false},
		{"create_empty_name_user", args{models.CreateUserInput{Name: ""}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.CreateUser(tt.args.inp); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

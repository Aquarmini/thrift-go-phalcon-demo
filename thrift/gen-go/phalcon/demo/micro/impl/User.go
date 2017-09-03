package impl

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"phalcon/demo/micro/config"
	"phalcon/demo/micro/service"
)

//import "runtime"
func init() {
	path := config.LOGDIR + "/user.go.log"
	file, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(file)
}

type User struct {
}

func (this *User) GetByUserId(user_id int64) (r *service.UserDTO, err error) {
	name := fmt.Sprintf("name[%d]", user_id)
	user := &service.UserDTO{
		UserID: user_id,
		Name:   name,
	}

	r = user
	return
}

func (this *User) Save(user_id int64, name string) (r bool, err error) {
	go log.WithFields(log.Fields{
		"user_id": user_id,
		"name":    name,
	}).Info("User Save")
	r = true
	return
}

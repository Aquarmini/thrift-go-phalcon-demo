package impl

import (
	"app/container"
	"github.com/sirupsen/logrus"
	"service"
)

type User struct {
}

func (this *User) GetByUserId(id int64) (r *service.UserDTO, err error) {
	di := container.GetInstance();

	user := &service.UserDTO{}
	user.UserID = id;
	user.Name = "limx";

	logger := di.Get("logger").(*logrus.Logger)
	logger.WithField("id", id).Infoln("User.GetByUserId");

	r = user
	return
}

func (this *User) Save(id int64, name string) (r bool, err error) {
	//di := container.GetInstance();
	//
	//user := &service.UserDTO{}
	//user.UserID = id;
	//user.Name = "limx";
	//
	//logger := di.Get("logger").(*logrus.Logger)
	//logger.WithField("id", id).Infoln("User.GetByUserId");

	r = true
	return
}

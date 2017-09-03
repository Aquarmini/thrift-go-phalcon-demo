package impl

import "fmt"
import "phalcon/demo/micro/service"
//import "runtime"

type User struct {
}

func (this *User) GetByUserId(user_id int64) (r *service.UserDTO, err error) {
	name := fmt.Sprintf("name[%d]", user_id)
	user := &service.UserDTO{
		UserID:user_id,
		Name:name,
	}

	r = user
	return
}

func (this *User) Save(user_id int64, name string) (r bool, err error) {

	r = true
	return
}

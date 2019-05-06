package module

import (
	"fmt"
	"goooo/source"
)

type UserController struct {
	source.App
}

type user struct {
	UserId string
	Name string
}

func (this *UserController) Login() {
	a := this.GetParams("a")
	fmt.Println(a)
	list := []user{}
	list = append(list, user{"1", "onlyfu"})
	list = append(list, user{"2", "yuiitsu"})
	this.Write(list)
}

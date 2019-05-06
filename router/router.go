package router

import (
	"goooo/module"
	"goooo/source"
)

func init() {
	source.Router("GET", "/", &module.GoodsController{})
	source.Router("GET", "/user/auth/Login", &module.UserController{})
}

package module

import (
	"goooo/source"
)

type GoodsController struct {
	source.App
}

func (this *GoodsController) GET() {
	list := []user{}
	list = append(list, user{"1", "onlyfu"})
	list = append(list, user{"2", "yuiitsu"})
	this.Write(list)
}

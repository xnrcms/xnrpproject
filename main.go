package main

import (
	"time"

	"github.com/xnrcms/xnrpproject/models"
	_ "github.com/xnrcms/xnrpproject/routers"

	"github.com/astaxie/beego"
	"github.com/xnrcms/xnrpproject/utils"
	cache "github.com/patrickmn/go-cache"
)

func main() {
	models.Init()
	utils.Che = cache.New(60*time.Minute, 120*time.Minute)
	beego.Run()
}

package routers

import (
	"github.com/astaxie/beego"
	"github.com/xnrcms/xnrpproject/controllers"
)

func init() {
	// 默认登录
	beego.Router("/", &controllers.ApiDocController{}, "*:Index")
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")
	beego.Router("/register", &controllers.RegisterController{}, "*:RegisterIn")

	//新增路径
	beego.Router("/project/my", &controllers.ProjectController{}, "*:My")

	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
	beego.AutoRouter(&controllers.ApiController{})
	beego.AutoRouter(&controllers.ApiSourceController{})
	beego.AutoRouter(&controllers.ApiPublicController{})
	beego.AutoRouter(&controllers.TemplateController{})
	beego.AutoRouter(&controllers.ApiDocController{})
	// beego.AutoRouter(&controllers.ApiMonitorController{})
	beego.AutoRouter(&controllers.EnvController{})
	beego.AutoRouter(&controllers.CodeController{})

	beego.AutoRouter(&controllers.GroupController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})

}

/**********************************************
** @Des: login
** @Author: haodaquan
** @Date:   2017-09-07 16:30:10
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:55:21
***********************************************/
package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/patrickmn/go-cache"
	"github.com/xnrcms/xnrpproject/libs"
	"github.com/xnrcms/xnrpproject/models"
	"github.com/xnrcms/xnrpproject/utils"
)

type LoginController struct {
	BaseController
}

//登录 TODO:XSRF过滤
func (self *LoginController) LoginIn() {
	if self.userId > 0 {
		self.redirect(beego.URLFor("HomeController.Index"))
	}
	beego.ReadFromRequest(&self.Controller)
	if self.isPost() {

		username := strings.TrimSpace(self.GetString("username"))
		password := strings.TrimSpace(self.GetString("password"))

		if username != "" && password != "" {
			user, err := models.AdminGetByName(username)

			/*errorMsg := ""*/
			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {
				self.ajaxMsg2("帐号或密码错误1", MSG_ERR,"")
			} else if user.Status == 0 {
				self.ajaxMsg2("该帐号已禁用", MSG_ERR,"")
			} else {
				//登录信息更新
				user.LastIp 		= self.getClientIp()
				user.LastLogin 		= time.Now().Unix()
				user.Update()

				//数据缓存
				utils.Che.Set("uid"+strconv.Itoa(user.Id), user, cache.DefaultExpiration)
				authkey := libs.Md5([]byte(self.getClientIp() + "|" + user.Password + user.Salt))
				self.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)

				self.ajaxMsg2("登录成功", MSG_OK, beego.URLFor("ProjectController.My"))
			}

			self.ajaxMsg2("帐号或密码异常", MSG_ERR,"")
		}
	}

	self.Layout 	= self.tplname + "/public/global_html.html"
	self.TplName 	= self.tplname + "/login/login.html"
}

//登出
func (self *LoginController) LoginOut() {
	self.Ctx.SetCookie("auth", "")
	self.redirect(beego.URLFor("LoginController.LoginIn"))
}

func (self *LoginController) NoAuth() {
	self.Ctx.WriteString("没有权限")
}

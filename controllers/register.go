/**********************************************
** @Des: login
** @Author: haodaquan
** @Date:   2017-09-07 16:30:10
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:55:21
***********************************************/
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/xnrcms/xnrpproject/libs"
	"github.com/xnrcms/xnrpproject/models"
	"strings"
	"time"
)

type RegisterController struct {
	BaseController
}

//注册 TODO:XSRF过滤
func (self *RegisterController) RegisterIn() {
	beego.ReadFromRequest(&self.Controller)
	if self.isPost() {

		username 	:= strings.TrimSpace(self.GetString("email"))
		nickname 	:= strings.TrimSpace(self.GetString("nickname"))
		password 	:= strings.TrimSpace(self.GetString("password"))
		repassword 	:= strings.TrimSpace(self.GetString("repassword"))

		if !libs.VerifyEmailFormat(username) {
			self.ajaxMsg2("帐号必须是邮箱格式", MSG_ERR,"")
		}

		if len(password) <= 0 || !libs.VerifyPasswordFormat(password) {
			self.ajaxMsg2("密码必须是6-20位字母数字组合", MSG_ERR,"")
		}

		if len(nickname) <= 0 {
			self.ajaxMsg2("用户昵称不能为空", MSG_ERR,"")
		}

		if password != repassword {
			self.ajaxMsg2("两次密码输入不一致", MSG_ERR,"")
		}

		//定义用户模型
		User 				:= new(models.Admin)
		
		//检测用户名是否存在
		if models.CheckFieldExist("login_name",username) {
			self.ajaxMsg2("你注册的账号已经存在", MSG_ERR,"")
		}

		pwd, salt 			:= libs.Password(4, password)

		User.LoginName 			= username
		User.RealName 			= nickname
		User.Password 			= pwd
		User.Email 				= username
		User.Salt 				= salt
		User.RegIp 				= self.getClientIp()
		User.UpdateId 			= 0
		User.CreateId 			= 0
		User.CreateTime 		= time.Now().Unix()
		User.UpdateTime 		= time.Now().Unix()
		User.Status 			= 1

		if _, err := models.AdminAdd(User); err != nil {
			self.ajaxMsg2(err.Error(), MSG_ERR,"")
		}

		self.ajaxMsg2("注册成功", MSG_OK,beego.URLFor("LoginController.LoginIn"))
	}

	self.Layout 	= self.tplname + "/public/global_html.html"
	self.TplName 	= self.tplname + "/register/register.html"
}
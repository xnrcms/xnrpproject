package controllers

type ProjectController struct {
	BaseController
}
func (self *ProjectController) My()  {

	self.Layout 	= self.tplname + "/public/global_html.html"
	self.TplName 	= self.tplname + "/project/my.html"
}
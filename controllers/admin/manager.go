package admin

import (
	"beego_xiaomi/models"
	"strings"
)

// ManagerController ...
type ManagerController struct {
	BaseController
}

// Get ...
func (c *ManagerController) Get() {
	c.TplName = "admin/manager/index.html"
}

// Add ...
func (c *ManagerController) Add() {
	// 获取所有的角色列表
	role := []models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"] = role

	c.TplName = "admin/manager/add.html"
}

func (c *ManagerController) DoAdd() {
	// 获取传递的数据
	roleId, err := c.GetInt("role_id")
	if err != nil {
		c.Error("非法请求", "/manager/add")
		return
	}

	username := strings.Trim(c.GetString("usernmae"), "")
	password := strings.Trim(c.GetString("password"), "")
	mobile := strings.Trim(c.GetString("mobile"), "")
	email := strings.Trim(c.GetString("email"), "")

	// 判断用户名和密码的长度是否合法
	if len(username) < 2 || len(password) < 6 {
		c.Error("用户名或密码长度不合法", "/manager/add")
		return
	}

	// 执行增加操作
	manager := models.Manager{}
	manager.Username = username
	manager.Password = models.Md5(password)
	manager.Mobile = mobile
	manager.Email = email
	manager.Status = 1
	manager.AddTime = int(models.GetUnix())
	manager.RoleID = roleId
	addErr := models.DB.Create(&manager).Error
	if addErr != nil {
		c.Error("增加管理员失败", "/manager/add")
		return
	} else {
		c.Success("增加管理员成功", "/manager")
	}
}

// Edit ...
func (c *ManagerController) Edit() {
	c.TplName = "admin/manager/edit.html"
}

func (c *ManagerController) DoEdit() {
	c.Ctx.WriteString("执行编辑")
}

func (c *ManagerController) Delete() {
	c.Ctx.WriteString("执行删除")
}

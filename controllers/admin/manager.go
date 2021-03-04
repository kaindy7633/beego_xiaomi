package admin

import (
	"beego_xiaomi/models"
	"strconv"
	"strings"
)

// ManagerController ...
type ManagerController struct {
	BaseController
}

// Get ...
func (c *ManagerController) Get() {
	// 获取管理员列表数据，并关联查询其角色类型
	manager := []models.Manager{}
	models.DB.Preload("Role").Find(&manager)
	c.Data["managerList"] = manager

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

	username := strings.Trim(c.GetString("username"), "")
	password := strings.Trim(c.GetString("password"), "")
	mobile := strings.Trim(c.GetString("mobile"), "")
	email := strings.Trim(c.GetString("email"), "")

	// 判断用户名和密码的长度是否合法
	if len(username) < 2 || len(password) < 6 {
		c.Error("用户名或密码长度不合法", "/manager/add")
		return
	}

	// 判断当前用户是否已存在
	managerList := []models.Manager{}
	models.DB.Where("username=?", username).Find(&managerList)
	if len(managerList) > 0 {
		c.Error("该用户名已存在", "/manager/add")
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
	manager.RoleId = roleId
	manager.IsSuper = 0
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
	// 获取需要编辑的条目的id
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("非法请求", "/manager")
		return
	}

	manager := models.Manager{Id: id}
	models.DB.Find(&manager)
	c.Data["managerList"] = manager

	// 获取所有的角色列表
	role := []models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"] = role

	c.TplName = "admin/manager/edit.html"
}

func (c *ManagerController) DoEdit() {
	// 获取传递的 manager id
	mId, mErr := c.GetInt("id")
	if mErr != nil {
		c.Error("非法请求", "/manager")
		return
	}

	// 获取当前 manager 的角色 id
	roleId, rErr := c.GetInt("role_id")
	if rErr != nil {
		c.Error("非法请求", "/manager")
		return
	}

	// 获取其他传递的数据
	password := strings.Trim(c.GetString("password"), "")
	mobile := strings.Trim(c.GetString("mobile"), "")
	email := strings.Trim(c.GetString("email"), "")

	// 获取数据并保存
	manager := models.Manager{Id: mId}
	models.DB.Find(&manager)
	manager.Mobile = mobile
	manager.Email = email
	manager.RoleId = roleId

	// 判断新密码长度是否合法
	if password != "" {
		if len(password) < 6 {
			c.Error("新密码长度不合法", "/manager/edit?id="+strconv.Itoa(mId))
			return
		}
		manager.Password = models.Md5(password)
	}

	// 执行修改
	err := models.DB.Save(&manager).Error
	if err != nil {
		c.Error("修改数据失败", "/manager/edit?id="+strconv.Itoa(mId))
	} else {
		c.Success("修改数据成功", "/manager")
	}
}

func (c *ManagerController) Delete() {
	// 获取要删除的条目的id
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("参数错误", "/manager")
		return
	}

	manager := models.Manager{Id: id}
	delErr := models.DB.Delete(&manager).Error
	if delErr != nil {
		c.Error("删除管理员失败", "/manager")
	} else {
		c.Success("删除管理员成功", "/manager")
	}

}

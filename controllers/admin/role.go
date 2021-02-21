package admin

import (
	"beego_xiaomi/models"
	"strconv"
	"strings"
)

type RoleController struct {
	BaseController
}

func (c *RoleController) Get() {
	role := []models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"] = role

	c.TplName = "admin/role/index.html"
}

func (c *RoleController) Add() {
	c.TplName = "admin/role/add.html"
}

func (c *RoleController) DoAdd() {
	title := strings.Trim(c.GetString("title"), "")
	description := strings.Trim(c.GetString("description"), "")

	if title == "" {
		c.Error("标题不能为空", "/role/add")
		return
	}

	role := models.Role{}
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = int(models.GetUnix())
	err := models.DB.Create(&role).Error
	if err != nil {
		c.Error("增加角色失败", "/role/add")
	} else {
		c.Success("增加角色成功", "/role")
	}

	// c.Ctx.WriteString("执行增加")
}

func (c *RoleController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}

	role := models.Role{Id: id}
	models.DB.Find(&role)

	c.Data["role"] = role
	c.TplName = "admin/role/edit.html"
}

func (c *RoleController) DoEdit() {
	// 获取需要修改的数据的id
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入的ID参数错误", "/role")
		return
	}

	// 获取需要修改的数据
	title := strings.Trim(c.GetString("title"), "")
	description := strings.Trim(c.GetString("description"), "")

	// 执行修改
	role := models.Role{Id: id}
	models.DB.Find(&role)
	role.Title = title
	role.Description = description

	saveErr := models.DB.Save(&role).Error
	if saveErr != nil {
		c.Error("修改角色失败", "/role/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改角色成功", "/role")
	}
}

func (c *RoleController) Delete() {
	// 获取传入的id
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入的ID参数错误", "/role")
		return
	}

	role := models.Role{Id: id}
	delErr := models.DB.Delete(&role).Error
	if delErr != nil {
		c.Error("删除角色失败", "/role")
	} else {
		c.Success("删除角色成功", "/role")
	}
}
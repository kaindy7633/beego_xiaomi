package admin

import (
	"beego_xiaomi/models"
	"strconv"
)

type AccessController struct {
	BaseController
}

func (c *AccessController) Get() {
	access := []models.Access{}
	models.DB.Preload("AccessItem").Where("module_id=0").Find(&access)
	c.Data["accessList"] = access

	c.TplName = "admin/access/index.html"
}

func (c *AccessController) Add() {
	// 获取顶级模块
	access := []models.Access{}
	models.DB.Where("module_id=0").Find(&access)
	c.Data["accessList"] = access

	c.TplName = "admin/access/add.html"
}

func (c *AccessController) DoAdd() {
	// 获取提交的数据
	moduleName := c.GetString("module_name")
	iType, typeErr := c.GetInt("type")
	actionName := c.GetString("action_name")
	url := c.GetString("url")
	moduleId, midErr := c.GetInt("module_id")
	sort, sErr := c.GetInt("sort")
	description := c.GetString("description")
	status, statusErr := c.GetInt("status")
	if typeErr != nil || midErr != nil || sErr != nil || statusErr != nil {
		c.Error("参数错误", "/access/add")
		return
	}

	access := models.Access{
		ModuleName:  moduleName,
		Type:        iType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
		AddTime:     int(models.GetUnix()),
	}
	err := models.DB.Create(&access).Error
	if err != nil {
		c.Error("增加数据失败", "/access/add")
	} else {
		c.Success("增加数据成功", "/access")
	}
}

func (c *AccessController) Edit() {
	// 获取要修改的数据id
	id, idErr := c.GetInt("id")
	if idErr != nil {
		c.Error("传入参数错误", "/access")
		return
	}

	access := models.Access{Id: id}
	models.DB.Find(&access)
	c.Data["access"] = access

	// 获取顶级模块
	accessList := []models.Access{}
	models.DB.Where("module_id=0").Find(&accessList)
	c.Data["accessList"] = accessList

	c.TplName = "admin/access/edit.html"
}

func (c *AccessController) DoEdit() {
	// 获取提交的数据
	id, idErr := c.GetInt("id")
	moduleName := c.GetString("module_name")
	iType, typeErr := c.GetInt("type")
	actionName := c.GetString("action_name")
	url := c.GetString("url")
	moduleId, midErr := c.GetInt("module_id")
	sort, sErr := c.GetInt("sort")
	description := c.GetString("description")
	status, statusErr := c.GetInt("status")
	if idErr != nil || typeErr != nil || midErr != nil || sErr != nil || statusErr != nil {
		c.Error("参数错误", "/access/edit")
		return
	}

	access := models.Access{Id: id}
	models.DB.Find(&access)
	access.ModuleName = moduleName
	access.Type = iType
	access.ActionName = actionName
	access.Url = url
	access.ModuleId = moduleId
	access.Sort = sort
	access.Description = description
	access.Status = status

	saveErr := models.DB.Save(&access).Error
	if saveErr != nil {
		c.Error("修改数据失败", "/access/edit?id="+strconv.Itoa(id))
		return
	} else {
		c.Success("修改数据成功", "/access")
	}
}

func (c *AccessController) Delete() {
	// 获取要删除的数据的id
	id, idErr := c.GetInt("id")
	if idErr != nil {
		c.Error("传入参数错误", "/access")
		return
	}

	// 获取当前数据
	access := models.Access{Id: id}
	models.DB.Find(&access)
	if access.ModuleId == 0 {
		// 如果是顶级模块，需要判断当前模块下是否还有子权限存在
		result := []models.Access{}
		models.DB.Where("module_id=?", access.Id).Find(&result)
		if len(result) > 0 {
			c.Error("顶级模块下的子模块不为空", "/access")
			return
		}
	}

	delErr := models.DB.Delete(&access).Error
	if delErr != nil {
		c.Error("删除数据失败", "/access")
		return
	} else {
		c.Success("删除数据成功", "/access")
	}
}

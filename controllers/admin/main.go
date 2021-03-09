package admin

import (
	"beego_xiaomi/models"

	beego "github.com/beego/beego/v2/server/web"
)

// MainController ...
type MainController struct {
	beego.Controller
}

// Get ...
func (c *MainController) Get() {
	// 获取当前登录用户的信息
	if userinfo, ok := c.GetSession("userinfo").(models.Manager); ok {
		// 获取用户名
		c.Data["username"] = userinfo.Username

		// 获取角色id
		roleId := userinfo.RoleId
		// 获取全部权限
		access := []models.Access{}
		models.DB.Preload("AccessItem").Where("module_id=0").Find(&access)
		// 获取当前角色拥有的权限并放在Map对象中
		roleAccess := []models.RoleAccess{}
		models.DB.Where("role_id=?", roleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}
		// 循环遍历所有的权限数据，判断当前权限id是否在角色权限Map对象中，如果是则设置Checked为true
		for i := 0; i < len(access); i++ {
			if _, ok := roleAccessMap[access[i].Id]; ok {
				access[i].Checked = true
			}
			for j := 0; j < len(access[i].AccessItem); j++ {
				if _, ok := roleAccessMap[access[i].AccessItem[j].Id]; ok {
					access[i].AccessItem[j].Checked = true
				}
			}
		}
		// 渲染权限数据
		c.Data["accessList"] = access
		// 超级管理员
		c.Data["isSuper"] = userinfo.IsSuper

	}

	c.TplName = "admin/main/index.html"
}

// Welcome ...
func (c *MainController) Welcome() {
	c.TplName = "admin/main/welcome.html"
}

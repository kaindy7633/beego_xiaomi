package middleware

import (
	"beego_xiaomi/models"
	"fmt"
	"net/url"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func AdminAuth(ctx *context.Context) {
	pathname := ctx.Request.URL.String()
	userinfo, ok := ctx.Input.Session("userinfo").(models.Manager)
	if !(ok && userinfo.Username != "") {
		if pathname != "/"+models.AdminPath+"/login" && pathname != "/"+models.AdminPath+"/login/doLogin" {
			ctx.Redirect(302, "/"+models.AdminPath+"/login")
		}
	} else {
		// 获取当前访问的url对应的权限id
		adminPath, _ := beego.AppConfig.String("admin_path")
		pathname = strings.Replace(pathname, "/"+adminPath, "", 1)
		urlPath, _ := url.Parse(pathname)

		fmt.Println(urlPath.Path)
		fmt.Println(excludeAuthPath(string(urlPath.Path)))

		// 判断是不是超级管理员
		if userinfo.IsSuper != 1 && !excludeAuthPath(string(urlPath.Path)) { // 如果不是超级管理员
			// 获取当前角色权限列表
			roleId := userinfo.RoleId
			roleAccess := []models.RoleAccess{}
			models.DB.Where("role_id=?", roleId).Find(&roleAccess)
			roleAccessMap := make(map[int]int)
			for _, v := range roleAccess {
				roleAccessMap[v.AccessId] = v.AccessId
			}

			access := models.Access{}
			models.DB.Where("url=?", urlPath.Path).Find(&access)

			// 判断当前访问的url对应的权限id是否在权限列表id中
			if _, ok := roleAccessMap[access.Id]; !ok {
				ctx.WriteString("没有权限")
				return
			}
		}
	}
}

// 定义方法用于判断当前访问路径是否在排除路径之内
func excludeAuthPath(urlPath string) bool {
	exclude_auth_paths, _ := beego.AppConfig.String("excludeAuthPath")
	excludeAuthPathSlice := strings.Split(exclude_auth_paths, ",")

	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}

	return false
}

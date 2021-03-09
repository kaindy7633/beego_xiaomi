package admin

import (
	"beego_xiaomi/models"
	"strings"

	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/server/web/captcha"
)

// LoginController ...
type LoginController struct {
	BaseController
}

// var admin_path, _ = beego.AppConfig.String("admin_path")
var cpt *captcha.Captcha

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)

	// 配置 cpt
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 40
}

// Get is ...
func (c *LoginController) Get() {
	//var user []models.User
	//models.DB.Find(&user)

	c.TplName = "admin/login/login.html"
	//c.Data["json"] = user
	//c.ServeJSON()
}

// DoLogin 验证码
func (c *LoginController) DoLogin() {
	// 1、验证码是否正确
	var flag = cpt.VerifyReq(c.Ctx.Request)
	if flag {
		// 2、获取表单值
		username := strings.Trim(c.GetString("username"), "")
		password := models.Md5(strings.Trim(c.GetString("password"), ""))

		// 3、匹配
		manager := []models.Manager{}
		models.DB.Where("username=? and password=?", username, password).Find(&manager)
		if len(manager) > 0 {
			// 登录成功
			// 设置 session 保存用户信息
			err := c.SetSession("userinfo", manager[0])
			if err != nil {
				c.Error("设置Session失败", "/")
				return
			}

			c.Success("登录成功", "/")
		} else {
			c.Error("用户名密码错误", "/login")
		}
	} else {
		// c.Ctx.WriteString("验证码错误")
		c.Error("验证码错误", "/login")
	}
}

// 退出登录
func (c *LoginController) LoginOut() {
	if err := c.DelSession("userinfo"); err != nil {
		c.Error("退出登录失败", "/login")
	}
	c.Success("退出登录成功", "/login")
}

package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/onethefour/REST-GO-demo/app/models"
)

type AccountController struct {
}

func (this *AccountController) Router(r *gin.Engine) {
	group := r.Group("/account")
	{
		group.GET("/login", this.login)
		group.GET("/logout", this.logout)
		group.GET("/get", this.get)
	}
}
func (this *AccountController) get(ctx *gin.Context) {
	var params models.GetAccountParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		//已注销
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	session := sessions.Default(ctx)
	account := session.Get("account")
	if account == nil {
		account = ""
	}
	ctx.JSON(200, gin.H{"code": 0, "message": "", "account": account})
	return
}
func (this *AccountController) logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("account")
	session.Save()

	ctx.JSON(200, gin.H{"code": 0, "message": "logout success"})
	return
}
func (this *AccountController) login(ctx *gin.Context) {
	var params models.LoginParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		ctx.JSON(200, gin.H{"code": 1, "message": err.Error()})
		return
	}
	pwd, has := models.Accounts[params.Account]
	if !has {
		ctx.JSON(200, gin.H{"code": 1, "message": "account not exist"})
		return
	}
	if pwd != params.Pwd {
		ctx.JSON(200, gin.H{"code": 1, "message": "password incorrect"})
		return
	}
	session := sessions.Default(ctx)
	session.Set("account", params.Account)
	session.Save()

	ctx.JSON(200, gin.H{"code": 0, "message": "success"})
	return
}

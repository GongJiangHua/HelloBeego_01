package controllers

import (
	"BeegoProject0603/db_mysql"
	"BeegoProject0603/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type QueryController struct {
	beego.Controller
}
func (c *QueryController) Post() {
	databases,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err!=nil {
		c.Ctx.WriteString("数据接收错误，请重试！！")
		return
	}
	var queryuser models.User
	err =json.Unmarshal(databases,&queryuser)
	//fmt.Println(queryuser)
	if err!=nil {
		c.Ctx.WriteString("对不起，数据解析错误！！")
		QueryResult := models.QueryResult{
			Code:    2,
			Message: "数据解析错误，请重新尝试",
			Data:    nil,
		}
		c.Data[`json`]=&QueryResult
		c.ServeJSON()
		return
	}

	//以name查询数据信息，并返回json格式数据
	row,err:=db_mysql.QueryUser(queryuser)
	if err != nil {
		QueryResult := models.QueryResult{
			Code:    0,
			Message: "用户查询失败",
			Data:    nil,
		}
		c.Data[`json`]=&QueryResult
		c.ServeJSON()
		return
	}
	if row==0 {
		//若查询出0条数据，说明不存在该用户
		QueryResult := models.QueryResult{
			Code:    2,
			Message: "该用户不存在！！",
			Data:    nil,
		}
		c.Data[`json`]=&QueryResult
		c.ServeJSON()
		return
	}else{
		//若row大于0则说明存在该用户，返回json数据
		QueryResult := models.QueryResult{
			Code:    1,
			Message: "恭喜你，该用户存在！！",
			Data:    queryuser,
		}
		c.Data[`json`]=&QueryResult
		c.ServeJSON()
	}
}

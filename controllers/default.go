package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MainController struct {
	beego.Controller
}

type PhonesController struct {
	beego.Controller
}

type ExcelController struct {
	beego.Controller
}

type Item struct {
	ID         string `db:"id"`
	Name       string `db:"pname"`
	Os         string `db:"os"`
	Ptype      string `db:"ptype"`
	Brand      string `db:"brand"`
	Size       string `db:"size"`
	Resolution string `db:"resolution"`
	Pcpu       string `db:"pcpu"`
	Link       string `db:"link"`
}

func (c *MainController) Get() {
	c.TplName = "app.html"
}

func (c *PhonesController) Get() {
	fmt.Println("this is in PhonesController")
	db, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/hotphones")
	if err != nil {
		fmt.Println("error1")
		fmt.Println(err)
		return
	}
	var item []Item
	error := db.Select(&item, "select id,pname,os,ptype,brand,size,resolution,pcpu,link from phones")
	if error != nil {
		fmt.Println("error2")
		fmt.Println(error)
		return
	}
	c.Data["json"] = item
	c.ServeJSON()
}

func (c *ExcelController) Get() {
	c.Ctx.Output.Download("./static/热门机型.xls")
}

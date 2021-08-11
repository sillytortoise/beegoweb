package controllers

import (
	"database/sql"
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

type FormController struct {
	beego.Controller
}

type Item struct {
	ID         string         `db:"id"`
	Name       string         `db:"pname"`
	Os         string         `db:"os"`
	Ptype      string         `db:"ptype"`
	Brand      string         `db:"brand"`
	Size       string         `db:"size"`
	Resolution string         `db:"resolution"`
	Pcpu       string         `db:"pcpu"`
	Link       string         `db:"link"`
	Screen     sql.NullString `db:"screen"`
	Img        string         `db:"img"`
}

func (c *MainController) Get() {
	c.TplName = "app.html"
}

func (c *PhonesController) Get() {
	db, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/hotphones")
	if err != nil {
		fmt.Println("error1")
		fmt.Println(err)
		return
	}
	var item []Item
	error := db.Select(&item, "select id,pname,os,ptype,brand,size,resolution,pcpu,link,screen,img from newphones")
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

func (c *FormController) Post() {
	db, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/hotphones")
	if err != nil {
		fmt.Println(err)
		return
	}

	screen := c.Input()
	for key, value := range screen {
		_, err := db.Exec("update newphones set screen=? where pname=?", value[0], key)
		if err != nil {
			fmt.Println("error3")
			fmt.Println(err)
			return
		}

	}
	c.Data["json"] = "success"
	c.ServeJSON()
}

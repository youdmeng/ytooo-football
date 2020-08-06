package repository

import (
	"fmt"
	"github.com/xormplus/xorm"
)
import _ "github.com/go-sql-driver/mysql"

var ENGINE *xorm.Engine

func GetEngine() *xorm.Engine {

	var err error
	ENGINE, err = xorm.NewEngine("mysql", "root:Qiuqiu0712!@tcp(192.168.10.58:3306)/football?charset=utf8")

	if err != nil {
		fmt.Println(err)
		return nil
	}
	ENGINE.ShowSQL(true)

	return ENGINE
}

package league

import (
	"../../model"
	"../../repository"
	"fmt"
	"github.com/xormplus/xorm"
	"log"
)

var engine *xorm.Engine

func init() {
	engine = repository.GetEngine()
	//连接测试
	if err := engine.Ping(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("数据库链接成功")
}

// 条件查询列表
func Select(ywhere *[]model.Ywhere) *[]model.League {
	var league []model.League
	session := engine.NewSession()
	for i := 0; i < len(*ywhere); i++ {
		session = session.Where((*ywhere)[i].Key+" "+(*ywhere)[i].Trans+" ? ", (*ywhere)[i].Value)
	}
	err := session.Find(&league)
	if err != nil {
		log.Print("查询失败\n", err)
		return nil
	}
	return &league
}

// 条件查询数量
func SelectByCount(ywhere *[]model.Ywhere) int64 {
	var league []model.League
	session := engine.NewSession()
	for i := 0; i < len(*ywhere); i++ {
		session = session.Where((*ywhere)[i].Key+" "+(*ywhere)[i].Trans+" ? ", (*ywhere)[i].Value)
	}
	count, err := session.FindAndCount(&league)
	if err != nil {
		log.Print("查询失败\n", err)
		return 0
	}
	return count
}

func Insert(league *model.League) int64 {

	if nil == league {
		return 0
	}
	session := engine.NewSession()

	row, err := session.Insert(*league)
	if nil == err {
		return 0
	}
	return row
}

//更新数据
func Update(league *model.League, ywhere *[]model.Ywhere) bool {
	session := engine.NewSession()
	for i := 0; i < len(*ywhere); i++ {
		session = session.Where((*ywhere)[i].Key+" "+(*ywhere)[i].Trans+" ? ", (*ywhere)[i].Value)
	}
	_, err := session.Update(*league)
	if nil == err {
		return true
	} else {
		return false
	}
}

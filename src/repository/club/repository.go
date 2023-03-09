package club

import (
	"fmt"
	"football/src/model"
	"football/src/repository"
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
func Select(ywhere *[]model.Ywhere) *[]model.Club {
	var club []model.Club
	session := engine.NewSession()
	for i := 0; i < len(*ywhere); i++ {
		session = session.Where((*ywhere)[i].Key+" "+(*ywhere)[i].Trans+" ? ", (*ywhere)[i].Value)
	}
	err := session.Find(&club)
	if err != nil {
		log.Print("查询失败\n", err)
		return nil
	}
	return &club
}

// 条件查询数量
func SelectByCount(ywhere *[]model.Ywhere) int64 {
	var club []model.Club
	session := engine.NewSession()
	for i := 0; i < len(*ywhere); i++ {
		session = session.Where((*ywhere)[i].Key+" "+(*ywhere)[i].Trans+" ? ", (*ywhere)[i].Value)
	}
	count, err := session.FindAndCount(&club)
	if err != nil {
		log.Print("查询失败\n", err)
		return 0
	}
	return count
}

func Update(club *model.Club, ywhere *[]model.Ywhere) bool {
	session := engine.NewSession()
	for i := 0; i < len(*ywhere); i++ {
		session = session.Where((*ywhere)[i].Key+" "+(*ywhere)[i].Trans+" ? ", (*ywhere)[i].Value)
	}
	_, err := session.Update(*club)
	if nil == err {
		return true
	} else {
		return false
	}
}

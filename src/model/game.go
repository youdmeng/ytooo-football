package model

import (
	"time"
)

type Game struct {
	Id            int       `xorm:"not null pk autoincr INT"`
	Key           string    `xorm:"comment('比赛标识') VARCHAR(100)"`
	Mid           int       `xorm:"comment('主队主键') INT"`
	MName         string    `xorm:"comment('主队名称') VARCHAR(255)"`
	Gid           int       `xorm:"comment('客队主键') INT"`
	GName         string    `xorm:"comment('客队名称') VARCHAR(255)"`
	MSort         string    `xorm:"comment('主队赛事排名') VARCHAR(100)"`
	GSort         string    `xorm:"comment('客队赛事排名') VARCHAR(100)"`
	LeagueName    string    `xorm:"comment('联赛名称') VARCHAR(255)"`
	Score         string    `xorm:"comment('实时比分') VARCHAR(100)"`
	Status        int       `xorm:"comment('比赛状态 0 未开始 1 进行中 2 已结束 3 已延期 4已取消') INT"`
	GameStartTime time.Time `xorm:"comment('比赛开始时间') DATETIME"`
}

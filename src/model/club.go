package model

type Club struct {
	Id     int    `xorm:"not null pk autoincr INT"`
	Key    string `xorm:"comment('俱乐部标识') VARCHAR(100)"`
	Name   string `xorm:"not null comment('俱乐部名称') VARCHAR(500)"`
	League string `xorm:"not null comment('所属联赛') VARCHAR(50)"`
	Soft   int    `xorm:"comment('联赛排名') INT"`
}

package model

type League struct {
	Id         int    `xorm:"not null pk comment('主键') INT"`
	Key        string `xorm:"not null comment('联赛标识') VARCHAR(100)"`
	Name       string `xorm:"not null comment('联赛名称') VARCHAR(500)"`
	Continents string `xorm:"not null comment('所属大洲 亚洲-Asia 欧洲-Europe 北美洲-North America 南美洲-South America 非洲-Africa 大洋洲-Oceania') VARCHAR(20)"`
}

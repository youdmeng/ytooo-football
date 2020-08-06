package main

import (
	"./model"
	"./repository/club"
	"fmt"
)

func main() {

	w1 := model.Ywhere{Key: "name", Trans: "like", Value: "a%"}
	w2 := model.Ywhere{Key: "id", Trans: "=", Value: "1"}
	ws := []model.Ywhere{w1, w2}
	fmt.Print(club.SelectByCount(&ws))
}

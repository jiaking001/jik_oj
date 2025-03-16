package model

import (
	"app/rpc/question/model/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}

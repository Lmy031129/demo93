package inits

import "user-srv/basic/config"

func init() {
	config.InitViper()
	InitMysql()
	InitRedis()
}

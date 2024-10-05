package initialize

import (
	"fmt"

	"github.com/go-ecommerce-backend-api/global"
)

func Run() {
	LoadingConfig()

	fmt.Println("Loading configuration mysql ", global.Config.Mysql.Username)
	InitLogger()
	InitMySQL()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}

package initialize

import (
	"fmt"

	"github.com/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	LoadingConfig()

	fmt.Println("Loading configuration mysql ", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Info("Config log ok", zap.String("ok", "success"))
	InitMySQL()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}

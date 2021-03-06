package server

import (
	"VueGin/global"
	router "VueGin/router"
	"log"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	//ver1. 使用gin設定路由、啟動API
	r := gin.Default()
	//此處直接匯入global Config Settings
	mode := global.Global_Config.System.Mode
	gin.SetMode(mode)
	//這兩種都可以
	// gin.SetMode(viper.GetString("system.mode"))
	// port := global.Global_Viper.GetString("system.port")
	router.Routers(r)

	port := global.Global_Config.System.Port
	err := r.Run(port)

	if err != nil {
		log.Fatalf("r.Run(port) failed:%v", err)
	}

	//TODO: ver2. 使用gin設定路由、使用原生Http監聽，並用middleware設定CasbinHandler攔截訊息
	// r := gin.Default()
	// r.Use(middleware.Cors())
	// gin.SetMode(viper.GetString("mode"))
	// //要用原生http的話，要再用一個CasbinHandler做middleware
	// Router := initSettings.Routers(r)
	// address := fmt.Sprintf(":%s", viper.GetString("port"))

	// s := initHTTPServer(address, Router)
	// err := s.ListenAndServe()
	// if err != nil {
	// 	log.Fatalf(" s.ListenAndServe() failed, err: %v", err)
	// }
}

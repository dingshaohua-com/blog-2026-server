package main

import (
	"blog-2026-server/routers"
	"blog-2026-server/utils"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("加载环境变量文件出错")
	}

	// 初始化连接池,程序结束时关闭连接池
	utils.InitDB()
	defer utils.CloseDB()

	// 注册路由
	r := routers.InitRouter()

	// 注册openApi
	utils.RegisterOpenapi(r)

	// 在启动前打印
	fmt.Println("🚀 系统启动成功：http://localhost:8080")

	// 启动服务：必须加上这一行！它会让程序“卡”在这里，持续监听请求，默认监听 8080 端口
	r.Run()
}

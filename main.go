package main

import (
	"etcdWeb/controller"
	"etcdWeb/etcd"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//连接数据库
	err := etcd.Init()
	if err != nil {
		fmt.Println("etcd init err", err)
		return
	}
	//创建路由引擎
	r := gin.Default()
	//加载静态模板
	r.Static("/xxx", "static")
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//等待请求
	r.GET("/", controller.Home)
	r.GET("/getEtcd", controller.CheckData)
	r.GET("/Test/etcd", controller.GetData)
	r.GET("/postEtcd", controller.PostEtcd)
	r.POST("/postEtcd", controller.PushEtcd)
	//启动引擎
	r.Run()
}

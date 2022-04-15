package controller

import (
	"etcdWeb/etcd"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func CheckData(ctx *gin.Context) {
	//查询数据库
	data := etcd.GetData()
	ctx.HTML(http.StatusOK, "getEtcd.html", data)
}

//获取输入的key、value
func GetData(ctx *gin.Context) {
	key := ctx.Query("one")
	value := ctx.Query("two")
	//放入chan中
	etcd.SendToChan(key, value)
}

func PostEtcd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "postEtcd.html", nil)
}

func PushEtcd(ctx *gin.Context) {
	//获取输入数据
	var data etcd.Etcd
	ctx.BindXML(&data)
	//将数据存到etcd中
	etcd.SendToChan(data.Key, data.Value)
}

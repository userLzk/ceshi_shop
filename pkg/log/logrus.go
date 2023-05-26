package log

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logPath = "./log.log"

//定义log实现接口
var logrusController LogrusController

type LogrusController struct {
}

func NewLogurs() {
	log := logrus.New()
	//设置log类型
	log.Formatter = &logrus.JSONFormatter{}
	//设置log等级
	log.SetLevel(logrus.InfoLevel)
	//生成log文件
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("log set err!!", err.Error())
	}
	log.Out = f
	gin.DefaultWriter = log.Out
}

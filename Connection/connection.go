package Connection

import (
	"database/sql"
	"fmt"
	"log"

	"ceshi_shop/config"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

//链接服务 mysql redis 等。。

var (
	ConfigData *config.ConfigData //配置文件
	//mysql client
	MysqlClient *sql.DB
	//redis client
	RedisPool *redis.Pool
)

func loadConfigData() {
	ConfigData = config.GetConfigDesc()
}

//连接函数
func ConnectionStart() {
	//加载配置
	loadConfigData()
	//连接mysql
	LoadMysqlClient()
	//连接redi
	LoadRedisClient()
}

func LoadMysqlClient() *sql.DB {
	conn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		ConfigData.MysqlConfig.User,
		ConfigData.MysqlConfig.PassWord,
		ConfigData.MysqlConfig.NetWork,
		ConfigData.MysqlConfig.Host,
		ConfigData.MysqlConfig.Port,
		ConfigData.MysqlConfig.Databases,
	)
	//获取配置
	Db, err := sql.Open("mysql", conn)
	//设置最大连接数
	Db.SetMaxIdleConns(100)
	//设置最大空闲连接数
	Db.SetMaxIdleConns(10)
	if err != nil {
		log.Printf("mysql open err:%v\n", err)
	}
	//连接测试
	//
	if pingErr := Db.Ping(); pingErr != nil {
		log.Printf("mysql open Ping err :%v\n", pingErr)
	}
	return Db
}

func LoadRedisClient() *redis.Pool {
	return &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}

}

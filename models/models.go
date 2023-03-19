package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"jaingke2023.com/BlogService/pkg/settings"
	"log"
)

//Model gorm的嵌入model结构体
type Model struct {
	Id         int   `gorm:"primaryKey" json:"id"`
	CreatedOn  int64 `gorm:"autoUpdateTime:nano" json:"created_on"`
	ModifiedOn int64 `gorm:"autoUpdateTime:nano" json:"modified_on"`
}

var (
	db *gorm.DB
	dbType,
	username,
	password, host, dbname, tablePrefix string
)

//init 初始化时现有Mysql数据库的方式进行连接
func init() {
	//1. 读取配置文件 conf/my.ini 下的数据库连接配置
	dbSection := settings.Cfg.Section("database")
	var err error
	//if err != nil {
	//	log.Fatal("Fail to get section 'database':", err)
	//	return
	//}
	dbType = dbSection.Key("TYPE").MustString("mysql")
	username = dbSection.Key("USER").MustString("root")
	password = dbSection.Key("PASSWORD").MustString("123456")
	host = dbSection.Key("HOST").MustString("192.168.157.131:3306")
	dbname = dbSection.Key("NAME").MustString("blog")
	tablePrefix = dbSection.Key("TABLE_PREFIX").MustString("blog_")
	
	//2. 使用gorm连接数据库，采取现有mysql数据库的方式连接对应的数据库
	//"user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local"
	
	options := "charset=utf8&parseTime=True&loc=Local"
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		username,
		password,
		host,
		dbname,
		options)
	db, err = gorm.Open(dbType, dbUrl)
	if err != nil {
		log.Fatalln("Failure to connect database :", err)
	}
	// 3.打开数据库连接后，我们配置一些默认可开启的gorm相关选项
	//数据库表名前缀默认处理
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	//默认使用单表 ，底层使用lock机制，并发安全
	db.SingularTable(true)
	//开启数据库使用的详细日志
	db.LogMode(true)
	//设置空闲连接池中的最大连接数。
	db.DB().SetMaxIdleConns(10)
	//设置数据库的最大打开连接数。
	db.DB().SetMaxOpenConns(100)
}

//CloseDB 用于关闭数据库
func CloseDB() {
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		
		}
	}(db)
}

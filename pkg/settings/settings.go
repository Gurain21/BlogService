package settings

import (
	"github.com/go-ini/ini"
	"log"
	
	"time"
)

var (
	Cfg *ini.File
	
	RunMode string
	
	PageSize  int
	JwtSecret int64
	
	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
)

func init() {
	var err error
	// 加载ini配置文件，读取指定路径的配置文件
	Cfg, err = ini.LooseLoad("conf/my.ini")
	if err != nil {
		log.Fatal("Fail to get source config file:" + err.Error())
		return
	}
	//获取 section 为 [""] 的所有配置
	loadDefault()
	
	//获取 section 为 [app] 的所有配置
	loadApp()
	
	//获取 section  [server] 的所有配置
	loadServer()
	
	//获取 section [database] 的所有配置
	//loadDB()
}

func loadDefault() {
	RunModeKey := Cfg.Section("").Key("RUN_MODE")
	//if err != nil {
	//	log.Fatal("Fail to get section '' :" + err.Error())
	//	return
	//}
	// 将key转换为对应的string类型，转换失败则取默认值
	RunMode = RunModeKey.MustString("debug")
}

func loadApp() {
	appSection := Cfg.Section("app")
	//if err != nil {
	//	log.Fatal("Fail to get section 'app':" + err.Error())
	//}
	psKey := appSection.Key("PAGE_SIZE")
	PageSize = psKey.MustInt(10)
	
	jwtKey := appSection.Key("JWT_SECRET")
	milli := time.Now().UnixMilli()
	JwtSecret = jwtKey.MustInt64(milli)
}

func loadServer() {
	serverSection := Cfg.Section("server")
	//if err != nil {
	//	log.Fatal("Fail to get section 'server':" + err.Error())
	//}
	hpKey := serverSection.Key("HTTP_PORT")
	HttpPort = hpKey.MustInt(9021)
	
	//使用 类型强制转换 将 int64转为 time.Duration 类型 再进行时间片段的计算
	//ReadTimeOut = time.Duration(serverSection.Key("READ_TIMEOUT").MustInt64(60)) * time.Second
	//WriteTimeOut = time.Duration(serverSection.Key("WRITE_TIMEOUT").MustInt64(60)) * time.Second
	//
	ReadTimeOut = serverSection.Key("READ_TIMEOUT").MustDuration(60 * time.Second)
	WriteTimeOut = serverSection.Key("WRITE_TIMEOUT").MustDuration(60 * time.Second)
	
}

//func loadDB() {
//	dbSection, err := Cfg.GetSection("database")
//	if err != nil {
//		log.Fatal("Fail to get section 'database':", err)
//		return
//	}
//	Type = dbSection.Key("TYPE").MustString("mysql")
//	User = dbSection.Key("USER").MustString("root")
//	Password = dbSection.Key("PASSWORD").MustString("123456")
//	Host = dbSection.Key("HOST").MustString("localhost:3306")
//	Name = dbSection.Key("NAME").MustString("blog")
//	TablePrefix = dbSection.Key("TABLE_PREFIX").MustString("blog_")
//}

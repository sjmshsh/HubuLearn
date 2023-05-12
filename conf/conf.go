package conf

import "gopkg.in/ini.v1"

// 例如我们要配置数据库，在项目里面添加数据库的相关配置信息

// MySQL的相关配置信息
var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

// redis
var (
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

// 我们的应用本身
var (
	AppMode  string
	HttpPort string
)

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("Db").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func Init() {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMysqlData(file)
	LoadRedis(file)
}

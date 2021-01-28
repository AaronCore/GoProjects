package util

type Config struct {
	SqlServerConf SqlServerConfig `ini:"server"`
	MySqlConf     MySqlConfig     `ini:"mysql"`
}

type SqlServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

type MySqlConfig struct {
	UserName string  `ini:"username"`
	PassWord string  `ini:"password"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	TimeOut  float32 `ini:"timeout"`
}

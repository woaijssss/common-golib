package conf

type MysqlConfig struct {
	Type        string `yaml:"Type"`
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Host        string `yaml:"Host"`
	Name        string `yaml:"Name"`
	PmsName     string `yaml:"PmsName"`
	TablePrefix string `yaml:"TablePrefix"`
	DBName      string `yaml:"DBName"`
}

var MysqlConf = &MysqlConfig{}

func GetMysqlConf(env string) MysqlConfig {
	//conf.Type = "mysql"
	//if env == "prod" {
	//	conf.MysqlHost = MysqlConfigProdHost
	//	conf.MysqlUser = MysqlConfigProdUser
	//	conf.MysqlPassword = MysqlConfigProdPassword
	//	return
	//}

	//conf.Host = MysqlConfigDevHost
	//conf.User = MysqlConfigDevUser
	//conf.Password = MysqlConfigDevPassword

	return *MysqlConf
}

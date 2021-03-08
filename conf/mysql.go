package conf

//const MysqlConfigDevHost = "127.0.0.1:3306"
//const MysqlConfigDevUser = "admin"
//const MysqlConfigDevPassword = "123456"

//const FWDBPrefix = "fw_"
//const FWDBParkingPrefix = "fw_"
//const OrderDBPrefix = "t_"
//const CashierDBPrefix = "t_"

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

//var DBPrefixName map[string]string

//func initDBPrefixName() map[string]string {
//	if DBPrefixName != nil {
//		return DBPrefixName
//	}
//
//	DBPrefixName = make(map[string]string)
//	DBPrefixName["db_property"] = FWDBPrefix
//	DBPrefixName["db_parking"] = FWDBParkingPrefix
//	DBPrefixName["db_order"] = OrderDBPrefix
//	DBPrefixName["db_cashier"] = CashierDBPrefix
//
//	return DBPrefixName
//}

/*
 * 根据数据库名称返回对应前缀
 */
//func GetDBPrefix(dbName string) string {
//	if prefix, ok := initDBPrefixName()[dbName]; ok {
//		return prefix
//	}
//
//	return ""
//}

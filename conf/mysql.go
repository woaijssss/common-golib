package conf

const MysqlConfigDevHost = "192.168.235.186:3306"
const MysqlConfigDevUser = "root"
const MysqlConfigDevPassword = "GmTech@2019"

const MysqlConfigProdHost = "10.0.0.8:3306"
const MysqlConfigProdUser = "db_property"
const MysqlConfigProdPassword = "zvtcnTqn5NYM"

const FWDBPrefix = "fw_"
const FWDBParkingPrefix = "fw_"
const OrderDBPrefix = "t_"
const CashierDBPrefix = "t_"

type MysqlConfig struct {
	MysqlType        string `yaml:"Type"`
	MysqlUser        string `yaml:"User"`
	MysqlPassword    string `yaml:"Password"`
	MysqlHost        string `yaml:"Host"`
	MysqlName        string `yaml:"Name"`
	MysqlPmsName     string `yaml:"PmsName"`
	MysqlTablePrefix string `yaml:"TablePrefix"`
}

func GetMysqlConf(env string) (conf MysqlConfig) {
	conf.MysqlType = "mysql"
	if env == "prod" {
		conf.MysqlHost = MysqlConfigProdHost
		conf.MysqlUser = MysqlConfigProdUser
		conf.MysqlPassword = MysqlConfigProdPassword
		return
	}

	conf.MysqlHost = MysqlConfigDevHost
	conf.MysqlUser = MysqlConfigDevUser
	conf.MysqlPassword = MysqlConfigDevPassword

	return
}

var DBPrefixName map[string]string

func initDBPrefixName() map[string]string {
	if DBPrefixName != nil {
		return DBPrefixName
	}

	DBPrefixName = make(map[string]string)
	DBPrefixName["db_property"] = FWDBPrefix
	DBPrefixName["db_parking"] = FWDBParkingPrefix
	DBPrefixName["db_order"] = OrderDBPrefix
	DBPrefixName["db_cashier"] = CashierDBPrefix

	return DBPrefixName
}

/*
 * 根据数据库名称返回对应前缀
 */
func GetDBPrefix(dbName string) string {
	if prefix, ok := initDBPrefixName()[dbName]; ok {
		return prefix
	}

	return ""
}

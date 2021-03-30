package conf

type MongoConfig struct {
	MongoConfPassword string
	MongoConfUser     string
	MongoConfUrl      string
}

const MongoConfUserOnline = "admin"
const MongoConfPasswordOnline = "123456"
const MongoConfUrlOnline = "127.0.0.1:28121"

const MongoConfUserDev = "admin"
const MongoConfPasswordDev = "123456"
const MongoConfUrlDev = "192.168.1.104:28121"

func GetMongoConf(env string) (conf MongoConfig) {
	if env == "offline" {
		conf.MongoConfUser = MongoConfUserDev
		conf.MongoConfPassword = MongoConfPasswordDev
		conf.MongoConfUrl = MongoConfUrlDev

		return
	}

	conf.MongoConfUser = MongoConfUserOnline
	conf.MongoConfPassword = MongoConfPasswordOnline
	conf.MongoConfUrl = MongoConfUrlOnline
	return
}

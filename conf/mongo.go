package conf

type MongoConfig struct {
	MongoConfPassword string
	MongoConfUser     string
	MongoConfUrl      string
}

const MongoConfUserOnline = "root"
const MongoConfPasswordOnline = "8NnzbHXGQD1F0pPR"
const MongoConfUrlOnline = "127.0.0.1:28121"

const MongoConfUserDev = "root"
const MongoConfPasswordDev = "gmtech1024"
const MongoConfUrlDev = "192.168.229.80:28121"

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

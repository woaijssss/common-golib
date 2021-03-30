package conf

const RedisHostDev = "192.168.1.104:6379"
const PasswordDev = "123456"

const RedisHostOnline = "127.0.0.1:6379"
const RedisPasswordOnline = "123456"
const RedisDB = 0
const RedisPoolMaxIdle = 100
const RedisPoolMaxActive = 10000
const RedisPoolIdleTimeout = 1

const ExpireTime = 86400

type RedisConfig struct {
	RedisHost            string
	RedisPassword        string
	RedisPoolMaxIdle     int
	RedisPoolMaxActive   int
	RedisPoolIdleTimeout int
}

func GetRedisConf(env string) (conf RedisConfig) {
	if env == "offline" {
		conf.RedisHost = RedisHostDev
		conf.RedisPassword = PasswordDev
		return
	}
	conf.RedisHost = RedisHostOnline
	conf.RedisPassword = RedisPasswordOnline
	return
}

func SetDefaultRedisConf(config *RedisConfig, env string) {
	if config.RedisHost == "" {
		if env == "offline" {
			config.RedisHost = RedisHostDev
		} else {
			config.RedisHost = RedisHostOnline
		}
	}

	if config.RedisPassword == "" {
		if env == "offline" {
			config.RedisPassword = PasswordDev
		} else {
			config.RedisPassword = RedisPasswordOnline
		}
	}

	if config.RedisPoolIdleTimeout == 0 {
		config.RedisPoolIdleTimeout = RedisPoolIdleTimeout
	}

	if config.RedisPoolMaxActive == 0 {
		config.RedisPoolMaxActive = RedisPoolMaxActive
	}

	if config.RedisPoolMaxIdle == 0 {
		config.RedisPoolMaxIdle = RedisPoolMaxIdle
	}
}

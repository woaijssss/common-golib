package conf

const RedisHostDev = "192.168.237.189:6379"
const PasswordDev = "KfUy0sCVuQTsXQv8"

const RedisHostOnline = "127.0.0.1:9763"
const RedisPasswordOnline = "KfUy0sCVuQTsXQv8"
const RedisDB = 0
const RedisPoolMaxIdle = 100
const RedisPoolMaxActive = 10000
const RedisPoolIdleTimeout = 1

const AllDevicePrefix = "ALL_DEVICE_LIST_PREFIX_"
const DeviceByIDPrefix = "DEVICE_INFO_BY_ID_PREFIX_"
const CommunityUserPrefix = "COMMUNITY_USER_PREFIX_"
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

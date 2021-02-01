package client

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/woaijssss/common-golib/app/logger"
	conf2 "github.com/woaijssss/common-golib/conf"
	"time"
)

var (
	// 定义常量
	RedisClient *redis.Pool
	RedisHost   string
	Env         string
	RedisDB     int
)

func Setup(env string) {
	initRedis(env)
}

func RedisSetup(env string) {
	initRedis(env)
}

func InitRedisByConf(conf *conf2.RedisConfig, env string) {
	conf2.SetDefaultRedisConf(conf, env)

	RedisHost = conf.RedisHost
	RedisDB = conf2.RedisDB
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     conf2.RedisPoolMaxIdle,
		MaxActive:   conf2.RedisPoolMaxActive,
		IdleTimeout: conf2.RedisPoolIdleTimeout * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisHost)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", conf.RedisPassword); err != nil {
				c.Close()
				return nil, err
			}
			// 选择db
			c.Do("SELECT", RedisDB)
			return c, nil
		},
	}
}

func initRedis(env string) {
	Env = env
	conf := conf2.GetRedisConf(env)
	// 从配置文件获取redis的ip以及db
	RedisHost = conf.RedisHost
	RedisDB = conf2.RedisDB
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     conf2.RedisPoolMaxIdle,
		MaxActive:   conf2.RedisPoolMaxActive,
		IdleTimeout: conf2.RedisPoolIdleTimeout * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisHost)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", conf.RedisPassword); err != nil {
				c.Close()
				return nil, err
			}
			// 选择db
			c.Do("SELECT", RedisDB)
			return c, nil
		},
	}
}

func getRedisConn(context *gin.Context) redis.Conn {
	client := RedisClient.Get()
	if client.Err() != nil {
		logger.Fatalf(context, "Get redis string err %+v", client.Err())
		initRedis(Env)
		return RedisClient.Get()
	}
	return client
}

func Get(context *gin.Context, key string) (string, error) {
	client := getRedisConn(context)

	defer client.Close()
	if client.Err() != nil {
		logger.Fatalf(context, "Get redis string err %+v", client.Err())
		return "", client.Err()
	}

	value, err := redis.String(client.Do("get", key))
	if err == redis.ErrNil {
		logger.Debugf(context, "Get redis string empty [%+v]", err)
		return "", nil
	}

	if err != nil {
		logger.Fatalf(context, "Get redis string err %+v", err)
		return "", err
	}

	return value, nil
}

func Set(context *gin.Context, key, val string, expire int64) error {
	client := getRedisConn(context)
	defer client.Close()

	if client.Err() != nil {
		logger.Fatalf(context, "Get redis string err %+v", client.Err())
		return client.Err()
	}

	_, err := client.Do("set", key, val, "ex", expire)
	if err != nil {
		logger.Fatalf(context, "Set redis string err %+v", err)
		return err
	}
	return err
}

func Incr(context *gin.Context, key string) error {
	client := getRedisConn(context)
	defer client.Close()

	if client.Err() != nil {
		logger.Fatalf(context, "incr redis string err %+v", client.Err())
		return client.Err()
	}
	_, err := client.Do("incr", key)
	if err != nil {
		logger.Fatalf(context, "incr redis string err %+v", err)
		return err
	}
	return err
}

func SetV2(context *gin.Context, key, val string) error {
	client := getRedisConn(context)
	defer client.Close()

	if client.Err() != nil {
		logger.Fatalf(context, "Get redis string err %+v", client.Err())
		return client.Err()
	}

	_, err := client.Do("set", key, val)
	if err != nil {
		logger.Fatalf(context, "Set redis string err %+v", err)
		return err
	}
	return err
}

func Expire(context *gin.Context, key string, expire int64) error {
	client := getRedisConn(context)
	defer client.Close()

	if client.Err() != nil {
		logger.Fatalf(context, "expire redis string err %+v", client.Err())
		return client.Err()
	}
	_, err := client.Do("expire", key, expire)
	if err != nil {
		logger.Fatalf(context, "expire redis string err %+v", err)
		return err
	}
	return err
}

func Delete(context *gin.Context, key string) error {
	client := RedisClient.Get()
	defer client.Close()

	if client.Err() != nil {
		logger.Fatalf(context, "Get redis string err %+v", client.Err())
		return client.Err()
	}

	_, err := client.Do("DEL", key)
	if err != nil {
		logger.Fatalf(context, "redis delelte failed: %+v", err)
	}
	return err
}

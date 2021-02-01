package client

import (
	"testing"

	"github.com/woaijssss/common-golib/conf"

	"github.com/gomodule/redigo/redis"
)

func TestRedis(t *testing.T) {
	config := &conf.RedisConfig{
		RedisHost:     "192.168.238.178:6379",
		RedisPassword: "GmTech@2019",
	}

	env := "offline"

	InitRedisByConf(config, env)

	client := RedisClient.Get()
	defer client.Close()

	key := "xhxhxhxhxhx"

	n, err := redis.Int(client.Do("setnx", key, "hxhxhxhxhxh"))

	if err != nil {
		t.Error(err)
	}

	if n != 1 {
		t.Error(n)
	}

	n, err = redis.Int(client.Do("setnx", key, "hxhxhxhxhxh"))

	if err != nil {
		t.Error(err)
	}

	if n != 0 {
		t.Error(n)
	}

}

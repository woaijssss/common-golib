package client

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/woaijssss/common-golib/app/context"
	"github.com/woaijssss/common-golib/app/logger"
	"github.com/woaijssss/common-golib/conf"
	"strings"
	"time"
)

var (
	env string
	db  *gorm.DB
)

func SetEnv(pEnv string) {
	env = pEnv
}

func GetDB() *gorm.DB {
	if db == nil {
		MysqlSetup()
	}
	return db
}

func MysqlSetup() {
	mysqlConfig := conf.GetMysqlConf(env)
	var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai&timeout=10s",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.DBName)
	var err error
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		logger.Fatalf(context.GetGinContextWithRequestId(), "models.Setup err: %v", err.Error())
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return conf.GetDBPrefix(db.Dialect().CurrentDatabase()) + defaultTableName
	//}

	myLogger := &MyLogger{}

	db.LogMode(true)
	db.SetLogger(myLogger)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Set("gorm:association", false)                //禁止自动创建/更新包含关系
	db.Set("gorm:association_save_reference", false) //禁止自动创建关联关系

	//if db == nil {
	//	db = make(map[string]*gorm.DB)
	//}
	//
	//db[dbName] = newDB
}

func CloseDB(dbName string) {
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
}

type MyLogger struct{}

func (l *MyLogger) Print(values ...interface{}) {
	var (
		level  = values[0]
		source = values[1]
	)

	ctx := context.GetGinContextWithRequestId()

	if level == "sql" {
		sql := strings.Replace(values[3].(string), "?", "%v", -1)
		for i := range values[4].([]interface{}) {
			values[4].([]interface{})[i] = transferSpecialType(values[4].([]interface{})[i])
		}
		logger.Infof(ctx, " MySQL: %s, Source: %s", fmt.Sprintf(sql, values[4].([]interface{})...), source)
	} else if level == "log" {
		logger.Infof(ctx, "MySQL %v", fmt.Sprintf("%v", values...))
	} else {
		logger.Infof(ctx, "MySQL %v", fmt.Sprintf("%v", values...))
	}
}

func transferSpecialType(obj interface{}) interface{} {
	switch t := obj.(type) {
	case time.Time:
		return fmt.Sprintf("'%s'", t.Format("2006-01-02 15:04:05"))
	case string:
		return fmt.Sprintf("'%s'", t)
	default:
		return obj
	}
}

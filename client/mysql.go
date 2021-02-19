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
	db  map[string]*gorm.DB
)

func SetEnv(pEnv string) {
	env = pEnv
}

func GetDB(dbName string) *gorm.DB {
	if _, ok := db[dbName]; !ok {
		MysqlSetup(dbName)
	}
	return db[dbName]
}

func GetYTBDB() *gorm.DB {
	return GetDB("ytb_service")
}

func MysqlSetup(dbName string) {
	mysqlConfig := conf.GetMysqlConf(env)
	var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai&timeout=10s",
		mysqlConfig.User,
		mysqlConfig.Password,
		mysqlConfig.Host,
		dbName)
	newDB, err := gorm.Open("mysql", dsn)

	if err != nil {
		logger.Fatalf(context.GetGinContextWithRequestId(), "models.Setup err: %v", err.Error())
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return conf.GetDBPrefix(db.Dialect().CurrentDatabase()) + defaultTableName
	//}

	myLogger := &MyLogger{}

	newDB.LogMode(true)
	newDB.SetLogger(myLogger)
	newDB.SingularTable(true)
	newDB.DB().SetMaxIdleConns(10)
	newDB.DB().SetMaxOpenConns(100)
	newDB.Set("gorm:association", false)                //禁止自动创建/更新包含关系
	newDB.Set("gorm:association_save_reference", false) //禁止自动创建关联关系

	if db == nil {
		db = make(map[string]*gorm.DB)
	}

	db[dbName] = newDB
}

// CloseDB closes database connection (unnecessary)
func CloseAllDB() {
	defer func() {
		for dbName, _ := range db {
			db[dbName].Close()
		}
	}()
}

func CloseDB(dbName string) {
	defer func() {
		if _, ok := db[dbName]; ok {
			db[dbName].Close()
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

package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
)

var DB *gorm.DB

func InitDB() {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")

	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=", charset, "&parseTime=True&loc=Local"}, "")

	err := Datanse(dsn)

	if err != nil {
		panic(err)
	}

}

func Datanse(dsn string) error {
	var ormLogger logger.Interface
	if gin.Mode() == gin.DebugMode {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true, // 重命名索引时删除和创建它，MySQL5.7前 不支持重命名索引
		DontSupportRenameColumn:   true, // 重命名列时重命名它，MySQL8前 不支持重命名列
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return err
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(20)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(60)
	DB = db

	migration()

	return nil
}

package config

import (
	"database/sql"
	"fmt"
	"kmi-be/model"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initLogger(config model.Config) {
	writer, err := rotatelogs.New(
		config.Logger.Path+".%Y%m%d",
		rotatelogs.WithLinkName(config.Logger.Path),
		rotatelogs.WithMaxAge(time.Duration(config.Logger.MaxAge*24)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(config.Logger.RotationTime*24)*time.Hour),
	)
	if err != nil {
		panic(err)
	}

	customFormatter := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	}
	logrus.SetFormatter(customFormatter)
	logrus.SetReportCaller(true)
	logrus.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
		},
		customFormatter,
	))
}

func InitConfig() (model.Config, *gorm.DB) {
	viper.SetConfigFile("config-dev.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config model.Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	initLogger(config)
	os.Setenv(config.Env.Code, config.Env.Timezone)

	db := initDB(config)

	return config, db
}

func initDB(config model.Config) *gorm.DB {
	conninfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=public",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Password,
		config.DB.DbName)

	pgx, err := sql.Open("pgx", conninfo)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: pgx,
	}), &gorm.Config{})

	if err != nil {
		logrus.Error(err)
		panic(err)
	} else {
		logrus.Info("Connected to database PostgreSQL")
	}

	return db

}

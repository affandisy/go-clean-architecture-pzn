package app

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viper *viper.Viper, log *logrus.Logger) (*gorm.DB, error) {
	username := viper.Get("database.username").(string)
	password := viper.Get("database.password").(string)
	host := viper.Get("database.host").(string)
	// port := int(viper.Get("database.port").(float64))
	port := viper.GetInt("database.port")
	database := viper.Get("database.name").(string)
	// idleConnection := viper.Get("database.pool.idle").(float64)
	// maxConnection := viper.Get("database.pool.max").(float64)
	// maxLifeTimeConnection := viper.Get("database.pool.lifetime").(float64)
	idleConnection := viper.GetInt("database.pool.idle")
	maxConnection := viper.GetInt("database.pool.max")
	maxLifeTimeConnection := viper.GetInt("database.pool.lifetime")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(&LogrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})
	if err != nil {
		return nil, err
	}

	connection, err := db.DB()
	if err != nil {
		return nil, err
	}

	connection.SetMaxIdleConns(int(idleConnection))
	connection.SetMaxOpenConns(int(maxConnection))
	connection.SetConnMaxIdleTime(time.Second * time.Duration(maxLifeTimeConnection))

	return db, nil
}

type LogrusWriter struct {
	Logger *logrus.Logger
}

func (l *LogrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}

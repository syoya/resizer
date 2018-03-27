package storage

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/syoya/resizer/logger"
	"github.com/syoya/resizer/options"
	"go.uber.org/zap"
)

type Storage struct {
	*gorm.DB
	l *zap.Logger
}

func New(o *options.Options) (*Storage, error) {
	dblogger := o.Logger.Named(logger.TagKeyDatabaseInitializing)
	var db *gorm.DB
	for {
		var err error
		db, err = gorm.Open("mysql", o.DataSourceName)
		if err == nil {
			break
		}
		dblogger.Warn("wait for connection", zap.Error(err))
		time.Sleep(time.Second)
	}
	db.LogMode(false)
	if o.Enviroment == "development" {
		// db.LogMode(true)
		// db.SetLogger(&Logger{})
		db.DropTable(&Image{})
	}
	db.CreateTable(&Image{})
	db.AutoMigrate(&Image{})

	for {
		err := db.DB().Ping()
		if err == nil {
			break
		}
		dblogger.Warn("wait for communication", zap.Error(err))
		time.Sleep(time.Second)
	}

	dblogger.Info("connected to database", zap.String(logger.FieldKeyDBDataSourceName, o.DataSourceName))
	return &Storage{db, o.Logger.Named(logger.TagKeyStorage)}, nil
}

// Close close db
func (s *Storage) Close() error {
	return s.DB.DB().Close()
}

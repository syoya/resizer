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
	var db *gorm.DB
	for {
		var err error
		db, err = gorm.Open("mysql", o.DataSourceName)
		if err == nil {
			break
		}
		o.Logger.Named(logger.TagKeyDatabaseInitializing).Warn("wait for connection", zap.Error(err))
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
		o.Logger.Named(logger.TagKeyDatabaseInitializing).Warn("wait for communication", zap.Error(err))
		time.Sleep(time.Second)
	}

	return &Storage{db, o.Logger.Named(logger.TagKeyFetcherStorage)}, nil
}

// Close close db
func (s *Storage) Close() error {
	return s.DB.DB().Close()
}

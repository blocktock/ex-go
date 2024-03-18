package app

import (
	"ex-go/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func InitDB() (*gorm.DB, func(), error) {

	db, err := initGorm()

	if err != nil {
		return nil, nil, err
	}

	cleanFunc := func() {
		//todo ... 错误处理
		d, _ := db.DB()
		//todo ... 错误处理
		_ = d.Close()
	}

	return db, cleanFunc, nil
}

func initGorm() (*gorm.DB, error) {

	sql := config.C.MySQL
	orm := config.C.Gorm

	dialectic := mysql.Open(sql.DSN())

	opts := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   orm.TablePrefix,
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialectic, opts)
	if err != nil {
		return nil, err
	}

	if orm.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(orm.MaxIdleConns)
	sqlDB.SetMaxOpenConns(orm.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(orm.MaxLifetime) * time.Second)

	return db, nil
}

// Package repository
// @author： Boice
// @createTime：2022/11/28 12:05
package repository

import (
	"context"
	"credit-platform/resource"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	ctxTxKey = "ctx:tx"
)

type (
	Client interface {
		DB(ctx context.Context) *gorm.DB
		Begin(ctx context.Context) context.Context
		Commit(ctx context.Context)
		Rollback(ctx context.Context)
	}
	client struct {
		db *gorm.DB
	}
)

func (c *client) Rollback(ctx context.Context) {
	value := ctx.Value(ctxTxKey)
	if value != nil {
		value.(*gorm.DB).Rollback()
	}
}

func (c *client) Commit(ctx context.Context) {
	value := ctx.Value(ctxTxKey)
	if value != nil {
		value.(*gorm.DB).Commit()
	}
}

func (c *client) Begin(ctx context.Context) context.Context {
	db := c.db.WithContext(ctx).Begin()
	ctx = context.WithValue(ctx, ctxTxKey, db)
	return ctx
}

func (c *client) DB(ctx context.Context) *gorm.DB {
	if ctx.Value(ctxTxKey) != nil {
		return ctx.Value(ctxTxKey).(*gorm.DB)
	}
	return c.db.WithContext(ctx)
}

func NewClient(conf resource.Config) Client {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DB.Username,
		conf.DB.Password,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.Database)
	var err error
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields:            true,
		Logger:                 dbLogger,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// Set the db connection configuration
	sqlDB.SetMaxOpenConns(conf.DB.MaxOpenConn) // set the max openning connection number
	sqlDB.SetMaxIdleConns(conf.DB.MaxIdleConn) // set the max idle connection number
	sqlDB.SetConnMaxLifetime(time.Minute * 10)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	return &client{
		db: db,
	}
}

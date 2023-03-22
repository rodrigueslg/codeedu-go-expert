package server

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CoinRepository struct {
	DB *gorm.DB
}

type CoinLog struct {
	ID  int `gorm:primary_key`
	Bid string
	gorm.Model
}

func NewSQLiteRepository() *CoinRepository {
	db, err := gorm.Open(sqlite.Open("./server/coinlog.db"), &gorm.Config{})
	if err != nil {
		fmt.Printf("can't connect to sqlite: %s\n", err)
		return nil
	}
	err = db.AutoMigrate(&CoinLog{})
	if err != nil {
		fmt.Printf("can't run migration: %s\n", err)
		return nil
	}
	return &CoinRepository{
		DB: db,
	}
}

func (r *CoinRepository) Log(ctx context.Context, cp *CoinPrice) {
	r.DB.WithContext(ctx).Debug().Create(&CoinLog{
		Bid: cp.UsdBrl.Bid,
	})
}

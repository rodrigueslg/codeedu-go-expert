package server

import (
	"context"

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
		panic(err)
	}
	err = db.AutoMigrate(&CoinLog{})
	if err != nil {
		panic(err)
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

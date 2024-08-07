package repositories

import (
	"candles-go/entities"
)

type CandleRepository interface {
	Migrate() error
	All() ([]entities.Candle, error)
	GetByName(name string) (*entities.Candle, error)
}

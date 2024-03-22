package repository

import (
	"exchangeRatesAPI/internal/models"
	"exchangeRatesAPI/internal/repository/postgresql"
	"github.com/jmoiron/sqlx"
	"time"
)

var database *sqlx.DB

type RateRepository struct {
	db *sqlx.DB
}

func InitRepository(cfg *models.Database) {
	// Init currently using DB type connection PostgreSQL for example
	database = postgresql.InitDBConn(cfg)
	return
}

func NewRepository() RateRepository {
	return RateRepository{db: database}
}

func (rr RateRepository) UpdateRate(fCurrency, sCurrency string, price float32, id string) (err error) {
	_, err = rr.db.Exec(`UPDATE actual SET price=$2, update_time=$3 WHERE currency = $1`, fCurrency+"/"+sCurrency, price, time.Now())

	_, err = rr.db.Exec(`insert into history (id,currency,price,update_time) VALUES ($1,$2,$3,$4)`, id, fCurrency+"/"+sCurrency, price, time.Now())

	return err
}

func (rr RateRepository) GetRateByUpdateIdentifier(id string) (price float32, update time.Time, err error) {
	var dest struct {
		Price      float32   `db:"price"`
		UpdateTime time.Time `db:"update_time"`
	}

	err = rr.db.Get(&dest, `SELECT price,update_time from history WHERE id = $1`, id)

	return dest.Price, dest.UpdateTime, err
}

func (rr RateRepository) GetLatestRate(fCurrency, sCurrency string) (price float32, update time.Time, err error) {
	var dest struct {
		Price      float32   `db:"price"`
		UpdateTime time.Time `db:"update_time"`
	}

	err = rr.db.Get(&dest, `SELECT price,update_time from actual WHERE currency = $1`, fCurrency+"/"+sCurrency)

	return dest.Price, dest.UpdateTime, err
}

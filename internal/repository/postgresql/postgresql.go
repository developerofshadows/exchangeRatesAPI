package postgresql

import (
	"exchangeRatesAPI/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDBConn(config *models.Database) *sqlx.DB {
	// Creating connection string and attempt to connect DB
	connectionString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password,
		config.Host, config.Port,
		config.Database)
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

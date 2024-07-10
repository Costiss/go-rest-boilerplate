package infra

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	conn, _ := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	conn.SetMaxIdleConns(0)
	conn.SetMaxOpenConns(10)

	db, _ := gorm.Open(postgres.New(postgres.Config{Conn: conn}))

	return db
}
